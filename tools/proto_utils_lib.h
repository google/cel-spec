#ifndef THIRD_PARTY_CEL_SPEC_TOOLS_PROTO_UTILS_LIB_H_
#define THIRD_PARTY_CEL_SPEC_TOOLS_PROTO_UTILS_LIB_H_

#include <string>

#include "google/rpc/status.pb.h"
#include "google/protobuf/io/zero_copy_stream_impl.h"
#include "absl/strings/string_view.h"

namespace google {
namespace api {
namespace expr {

class ProtoUtils {
 public:
  ProtoUtils() = default;
  ProtoUtils(const ProtoUtils&) = default;
  ProtoUtils& operator=(const ProtoUtils&) = default;

  // Reads from the text input, parses it and returns a binary format. Proto
  // for message_name should either have been linked or explicity loaded via
  // LoadProtoDefinition.
  std::string TextToBinary(absl::string_view src, absl::string_view message_name,
                      google::rpc::Status* status);

  // Similar to above, but works with streams.
  google::rpc::Status TextToBinary(google::protobuf::io::ZeroCopyInputStream* input,
                                   absl::string_view message_name,
                                   google::protobuf::io::ZeroCopyOutputStream* output);

  // Loads the proto file form the given file and adds the descriptors.
  // All dependencies(imports) for the file should have been added
  // or linked to the binary, otherwise this will fail.
  google::rpc::Status LoadProtoDefinition(absl::string_view file_path,
                                          absl::string_view file_name);

 private:
  google::protobuf::DescriptorPool pool_;
};

}  // namespace expr
}  // namespace api
}  // namespace google

#endif  // THIRD_PARTY_CEL_SPEC_TOOLS_PROTO_UTILS_LIB_H_
