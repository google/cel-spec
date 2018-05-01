package(
    default_visibility = ["//visibility:public"],
)

licenses(["notice"])  # Apache 2.0

proto_library(
    name = "test_value_proto",
    srcs = [
        "test_value.proto",
    ],
    # cc_api_version = 2,
    # java_api_version = 2,
    deps = [
        "//third_party/cel/spec/proto/v1:value_protos",
    ],
)

proto_library(
    name = "test_message1_proto",
    srcs = [
        "test_message1.proto",
    ],
    # cc_api_version = 2,
    # java_api_version = 2,
)

proto_library(
    name = "test_message2_proto",
    srcs = [
        "test_message2.proto",
    ],
    # cc_api_version = 2,
    # java_api_version = 2,
    deps = [
        ":test_message1_proto",
    ],
)

proto_library(
    name = "test_message3_proto",
    srcs = [
        "test_message3.proto",
    ],
    # cc_api_version = 2,
    # java_api_version = 2,
)

proto_library(
    name = "test_message4_proto",
    srcs = [
        "test_message4.proto",
    ],
    # cc_api_version = 2,
    # java_api_version = 2,
    deps = [
        ":test_message3_proto",
    ],
)
