package(
    default_hdrs_check = "strict",
    default_visibility = ["//visibility:public"],
)

licenses(["notice"])  # Apache 2.0

exports_files(["LICENSE"])

cc_library(
    name = "proto_utils_lib",
    srcs = [
        "proto_utils_lib.cc",
    ],
    hdrs = [
        "proto_utils_lib.h",
    ],
    deps = [
        "//base",
        "//google/protobuf:any",
        "//google/protobuf:duration",
        "//google/protobuf:empty",
        "//google/protobuf:struct",
        "//google/protobuf:timestamp",
        "//google/rpc:code",
        "//google/rpc:status",
        "//net/proto2/compiler/public:importer",
        "//net/proto2/io/public:io",
        "//net/proto2/public:proto2",
        "//third_party/absl/strings",
        "//third_party/cel/spec/proto/v1:value_protos",
    ],
)

cc_test(
    name = "proto_utils_lib_test",
    srcs = ["proto_utils_lib_test.cc"],
    data = [
        "//third_party/cel/spec/tools/tests:test_message2.proto",
        "//third_party/cel/spec/tools/tests:test_message4.proto",
    ],
    deps = [
        ":proto_utils_lib",
        "//base",
        "//google/api/expr/v1beta1:value_proto",
        "//google/protobuf:any",
        "//google/rpc:status",
        "//testing/base/public:gunit_main",
        "//third_party/absl/strings",
        "//third_party/cel/spec/proto/v1:value_protos",
        "//third_party/cel/spec/tools/tests:test_message1_proto",
        "//third_party/cel/spec/tools/tests:test_value_proto",
    ],
)

cc_binary(
    name = "prototext2bin",
    srcs = ["prototext2bin.cc"],
    deps = [
        ":proto_utils_lib",
        "//base",
        "//google/rpc:status",
        "//net/proto2/io/public:io",
        "//third_party/absl/strings",
    ],
)
