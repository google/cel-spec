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
  build_file_content = "proto_library(name = 'rpc_status', srcs = ['google/rpc/status.proto'], deps = ['@com_google_protobuf//:any_proto', '@com_google_protobuf//:empty_proto'], visibility = ['//visibility:public']), proto_library(name = 'rpc_code', srcs = ['google/rpc/code.proto'], visibility = ['//visibility:public']), cc_proto_library(name = 'cc_rpc_status', deps = [':rpc_status'], visibility = ['//visibility:public']), cc_proto_library(name = 'cc_rpc_code', deps = [':rpc_code'], visibility = ['//visibility:public'])"
)

http_archive(
     name = "com_google_googletest",
     urls = ["https://github.com/google/googletest/archive/master.zip"],
     strip_prefix = "googletest-master",
)

# gflags
http_archive(
    name = "com_github_gflags_gflags",
    sha256 = "6e16c8bc91b1310a44f3965e616383dbda48f83e8c1eaa2370a215057b00cabe",
    strip_prefix = "gflags-77592648e3f3be87d6c7123eb81cbad75f9aef5a",
    urls = [
        "https://mirror.bazel.build/github.com/gflags/gflags/archive/77592648e3f3be87d6c7123eb81cbad75f9aef5a.tar.gz",
        "https://github.com/gflags/gflags/archive/77592648e3f3be87d6c7123eb81cbad75f9aef5a.tar.gz",
    ],
)

# glog
http_archive(
    name = "com_google_glog",
    sha256 = "1ee310e5d0a19b9d584a855000434bb724aa744745d5b8ab1855c85bff8a8e21",
    strip_prefix = "glog-028d37889a1e80e8a07da1b8945ac706259e5fd8",
    urls = [
        "https://mirror.bazel.build/github.com/google/glog/archive/028d37889a1e80e8a07da1b8945ac706259e5fd8.tar.gz",
        "https://github.com/google/glog/archive/028d37889a1e80e8a07da1b8945ac706259e5fd8.tar.gz",
    ],
)

http_archive(
    name = "com_google_absl",
    strip_prefix = "abseil-cpp-master",
    urls = ["https://github.com/abseil/abseil-cpp/archive/master.zip"],
)
