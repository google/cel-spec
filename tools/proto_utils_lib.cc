#include "tools/proto_utils_lib.h"

#include <fcntl.h>
#include <stdio.h>
#include <sstream>

#include "google/rpc/code.pb.h"
#include "google/protobuf/compiler/parser.h"
#include "google/protobuf/io/tokenizer.h"
#include "google/protobuf/io/zero_copy_stream_impl.h"
#include "google/protobuf/dynamic_message.h"
#include "google/protobuf/text_format.h"
#include "absl/strings/str_cat.h"
#include "absl/strings/str_join.h"
#include "absl/strings/str_split.h"

namespace google {
namespace api {
namespace expr {

namespace {

google::rpc::Status OkStatus() { return google::rpc::Status(); }

google::rpc::Status InvalidArgumentError(absl::string_view message) {
  google::rpc::Status error;
  error.set_code(google::rpc::Code::INVALID_ARGUMENT);
  error.set_message(std::string(message));
  return error;
}

inline bool IsOk(const google::rpc::Status& status) {
  return status.code() == google::rpc::Code::OK;
}

class StringErrorCollector : public google::protobuf::io::ErrorCollector {
 public:
  explicit StringErrorCollector(std::string* error_text) : error_text_(error_text) {}

  void AddError(int line, int column, const std::string& message) override {
    absl::StrAppend(error_text_, "%d(%d): %s\n", line, column, message.c_str());
  }

  void AddWarning(int line, int column, const std::string& message) override {
    absl::StrAppend(error_text_, "%d(%d): %s\n", line, column, message.c_str());
  }

 private:
  std::string* error_text_;
  StringErrorCollector(const StringErrorCollector&) = delete;
  StringErrorCollector& operator=(const StringErrorCollector&) = delete;
};

class ProtoDescriptorFinder : public google::protobuf::TextFormat::Finder {
 public:
  explicit ProtoDescriptorFinder(const google::protobuf::DescriptorPool* pool)
      : pool_(pool),
        generated_pool_(google::protobuf::DescriptorPool::generated_pool()) {}

  const google::protobuf::FieldDescriptor* FindExtension(
      google::protobuf::Message* message, const std::string& name) const override {
    auto* descriptor = generated_pool_->FindExtensionByName(name);
    if (descriptor != nullptr) {
      return descriptor;
    }
    return pool_->FindExtensionByName(name);
  }

  // In open source, TextFormat::Finder does not provide a virtual FindAnyType
  // method. We can get rid of this preprocessor statement once open source
  // version of protos enable overriding FindAnyType.
#ifndef CEL_SPEC_OPEN_SOURCE
  const google::protobuf::Descriptor* FindAnyType(const google::protobuf::Message& message,
                                        const std::string& prefix,
                                        const std::string& name) const override {
    auto* descriptor =
        google::protobuf::TextFormat::Finder::FindAnyType(message, prefix, name);
    if (descriptor != nullptr) {
      return descriptor;
    }

    descriptor = generated_pool_->FindMessageTypeByName(name);
    if (descriptor != nullptr) {
      return descriptor;
    }

    return pool_->FindMessageTypeByName(name);
  }
#endif

 private:
  const google::protobuf::DescriptorPool* pool_;
  const google::protobuf::DescriptorPool* generated_pool_;
};

google::rpc::Status ParseProtoDefinition(
    absl::string_view file_path, absl::string_view file_name,
    google::protobuf::FileDescriptorProto* file_desc) {
  FILE* fp = fopen(std::string(file_path).c_str(), "r");
  if (fp == nullptr) {
    return InvalidArgumentError(absl::StrCat("Cannot open ", file_path));
  }

  std::string err;
  StringErrorCollector collector(&err);
  google::protobuf::io::FileInputStream input(fileno(fp));
  google::protobuf::io::Tokenizer tokenizer(&input, &collector);
  google::protobuf::compiler::Parser parser;
  parser.RecordErrorsTo(&collector);

  if (!parser.Parse(&tokenizer, file_desc)) {
    return InvalidArgumentError(
        absl::StrCat("Cannot parse ", file_path, ": ", err));
  }
  file_desc->set_name(std::string(file_name));
  return OkStatus();
}

google::rpc::Status LoadWithDependencies(
    const google::protobuf::FileDescriptor& descriptor, google::protobuf::DescriptorPool* pool) {
  for (int i = 0; i < descriptor.dependency_count(); i++) {
    auto* dep = descriptor.dependency(i);
    if (dep == nullptr) {
      continue;
    }
    google::rpc::Status recursive_satus = LoadWithDependencies(*dep, pool);
    if (!IsOk(recursive_satus)) {
      return recursive_satus;
    }
  }
  google::protobuf::FileDescriptorProto dep_file;
  descriptor.CopyTo(&dep_file);
  // This will not be a nullptr if we loaded all dependencies to
  // pool successfully.
  if (pool->BuildFile(dep_file) == nullptr) {
    return InvalidArgumentError(
        absl::StrCat("Cannot load ", descriptor.name()));
  }

  return OkStatus();
}

}  // namespace

google::rpc::Status ProtoUtils::LoadProtoDefinition(
    absl::string_view file_path, absl::string_view file_name) {
  google::protobuf::FileDescriptorProto file_proto;
  google::rpc::Status status =
      ParseProtoDefinition(file_path, file_name, &file_proto);
  if (!IsOk(status)) {
    return status;
  }

  for (const std::string& dep : file_proto.dependency()) {
    if (pool_.FindFileByName(dep) != nullptr) {
      continue;
    }
    auto* dep_descriptor =
        google::protobuf::DescriptorPool::generated_pool()->FindFileByName(dep);
    if (dep_descriptor == nullptr) {
      return InvalidArgumentError(absl::StrCat(
          "Required proto", dep, " for ", file_path, " is not loaded/linked"));
    }
    LoadWithDependencies(*dep_descriptor, &pool_);
  }

  if (pool_.BuildFile(file_proto) == nullptr) {
    return InvalidArgumentError(absl::StrCat("Cannot load ", file_path));
  }
  return OkStatus();
}

google::rpc::Status ProtoUtils::TextToBinary(
    google::protobuf::io::ZeroCopyInputStream* input, absl::string_view message_name,
    google::protobuf::io::ZeroCopyOutputStream* output) {
  auto* descriptor =
      google::protobuf::DescriptorPool::generated_pool()->FindMessageTypeByName(
          std::string(message_name));
  if (descriptor == nullptr) {
    descriptor = pool_.FindMessageTypeByName(std::string(message_name));
  }
  if (descriptor == nullptr) {
    return InvalidArgumentError(
        absl::StrCat("Cannot find message type: ", message_name));
  }

  google::protobuf::DynamicMessageFactory factory;
  auto* prototype = factory.GetPrototype(descriptor);
  std::unique_ptr<google::protobuf::Message> message(prototype->New());

  ProtoDescriptorFinder finder(&pool_);
  google::protobuf::TextFormat::Parser parser;
  parser.AllowPartialMessage(true);
  parser.SetFinder(&finder);

  if (!parser.Parse(input, message.get())) {
    return InvalidArgumentError(
        absl::StrCat("Could not parse input for: ", message_name));
  }

  if (!message->SerializePartialToZeroCopyStream(output)) {
    return InvalidArgumentError(
        absl::StrCat("Could not serialize: ", message_name));
  }

  return OkStatus();
}

std::string ProtoUtils::TextToBinary(absl::string_view src,
                                absl::string_view message_name,
                                google::rpc::Status* status) {
  std::stringstream input{std::string(src)};
  std::ostringstream output;
  {
    google::protobuf::io::IstreamInputStream input_stream{&input};
    google::protobuf::io::OstreamOutputStream output_stream{&output, 0};
    *status = TextToBinary(&input_stream, message_name, &output_stream);
  }
  return output.str();
}

}  // namespace expr
}  // namespace api
}  // namespace google
