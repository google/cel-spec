http_archive(
    name = "com_google_protobuf",
    strip_prefix = "protobuf-3.5.0",
    urls = ["https://github.com/google/protobuf/archive/v3.5.0.zip"],
)
http_archive(
    name = "com_google_protobuf_javalite",
    strip_prefix = "protobuf-javalite",
    urls = ["https://github.com/google/protobuf/archive/javalite.zip"],
)
new_http_archive(
  name = "com_google_googleapis",
  url = "https://github.com/googleapis/googleapis/archive/common-protos-1_3_1.zip",
  strip_prefix = "googleapis-common-protos-1_3_1/",
  build_file_content = "proto_library(name = 'rpc_status', srcs = ['google/rpc/status.proto'], deps = ['@com_google_protobuf//:any_proto', '@com_google_protobuf//:empty_proto'], visibility = ['//visibility:public'])"
)
