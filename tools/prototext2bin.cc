#include <unistd.h>

#include "gflags/gflags.h"
#include "google/rpc/status.pb.h"
#include "google/protobuf/io/zero_copy_stream_impl.h"
#include "absl/strings/str_split.h"
#include "tools/proto_utils_lib.h"

DEFINE_string(protos, "",
              "Comma separated list of proto files:filenames pairs to "
              "dynamically load.");

DEFINE_string(message_name, "", "Fully qualified message to parse into.");

namespace {

google::rpc::Status TryLoad(google::api::expr::ProtoUtils* proto_utils) {
  google::rpc::Status status;
  std::string proto_files = FLAGS_protos;
  if (proto_files.empty()) {
    return google::rpc::Status{};
  }

  std::vector<std::string> protos = absl::StrSplit(proto_files, ',');
  for (const auto& proto_file : protos) {
    std::vector<std::string> file_path_name = absl::StrSplit(proto_file, ':');
    if (file_path_name.size() != 2) {
      status.set_message(
          "When loading a proto, both file_path and "
          "file_name should be provided separated by a colon");
      return status;
    }
    status =
        proto_utils->LoadProtoDefinition(file_path_name[0], file_path_name[1]);
    if (status.code() != 0) {
      return status;
    }
  }
  return google::rpc::Status{};
}

google::rpc::Status Parse(google::api::expr::ProtoUtils* proto_utils) {
  google::protobuf::io::FileInputStream input_stream(STDIN_FILENO);
  google::protobuf::io::FileOutputStream output_stream(STDOUT_FILENO);
  google::rpc::Status status = proto_utils->TextToBinary(
      &input_stream, FLAGS_message_name, &output_stream);
  output_stream.Flush();
  return status;
}

}  // namespace

int main(int argc, char* argv[]) {
  google::ParseCommandLineFlags(&argc, &argv, true);

  google::api::expr::ProtoUtils proto_utils;
  google::rpc::Status status = TryLoad(&proto_utils);
  if (status.code() != 0) {
    std::cerr << status.message() << std::endl << std::flush;
    return 0;
  }

  status = Parse(&proto_utils);
  if (status.code() != 0) {
    std::cerr << status.message() << std::endl << std::flush;
    return 0;
  }
  return 0;
}
