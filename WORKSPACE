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
    url = "https://github.com/googleapis/googleapis/archive/master.zip",
    strip_prefix = "googleapis-master/",
    build_file_content = """
load('@io_bazel_rules_go//proto:def.bzl', 'go_proto_library')

proto_library(
    name = 'rpc_status',
    srcs = ['google/rpc/status.proto'],
    deps = [
        '@com_google_protobuf//:any_proto',
        '@com_google_protobuf//:empty_proto',
    ],
    visibility = ['//visibility:public'],
)

proto_library(
    name = 'rpc_code',
    srcs = ['google/rpc/code.proto'],
    visibility = ['//visibility:public'],
)

proto_library(
    name = 'expr_v1beta1',
    srcs = [
        'google/api/expr/v1beta1/eval.proto',
        'google/api/expr/v1beta1/value.proto',
        ],
    deps = [
        '@com_google_protobuf//:any_proto',
        '@com_google_googleapis//:rpc_status',
        '@com_google_protobuf//:struct_proto',
    ],
    visibility = ['//visibility:public'],
)

cc_proto_library(
    name = 'cc_rpc_status',
    deps = [':rpc_status'],
    visibility = ['//visibility:public'],
)

cc_proto_library(
    name = 'cc_rpc_code',
    deps = [':rpc_code'],
    visibility = ['//visibility:public'],
)

cc_proto_library(
    name = 'cc_expr_v1beta1',
    deps = [':expr_v1beta1'],
    visibility = ['//visibility:public'],
)

go_proto_library(
    name = 'rpc_status_go_proto',
    # TODO: Switch to the correct import path when bazel rules fixed.
    #importpath = 'google.golang.org/genproto/googleapis/rpc/status',
    importpath = 'github.com/googleapis/googleapis/google/rpc',
    proto = ':rpc_status',
    visibility = ['//visibility:public'],
)

go_proto_library(
    name = 'expr_v1beta1_go_proto',
    importpath = 'google.golang.org/genproto/googleapis/api/expr/v1beta1',
    proto = 'expr_v1beta1',
    visibility = ['//visibility:public'],
    deps = ['@com_google_googleapis//:rpc_status_go_proto'],
)
"""
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

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.12.0/rules_go-0.12.0.tar.gz",
    sha256 = "c1f52b8789218bb1542ed362c4f7de7052abcf254d865d96fb7ba6d44bc15ee3",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
