package(
    # default_compatible_with = ["//buildenv/target:appengine"],
    default_visibility = ["//visibility:public"],
)

licenses(["notice"])  # Apache 2.0

exports_files(["LICENSE"])

proto_library(
    name = "value_protos",
    srcs = [
        "value.proto",
    ],
    # cc_api_version = 2,
    # go_api_version = 2,
    # java_api_version = 2,
    # js_api_version = 2,
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:struct_proto",
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "value_protos_go_proto",
    deps = [":value_protos"],
)

alias(
    name = "value_go_proto",
    actual = ":value_protos_go_proto",
)

java_proto_library(
    name = "value_java_proto",
    deps = [":value_protos"],
)

java_lite_proto_library(
    name = "value_java_proto_lite",
    deps = [":value_protos"],
)

jspb_proto_library(
    name = "value_jspb_proto",
    deps = [":value_protos"],
)

proto_library(
    name = "syntax_protos",
    srcs = [
        "syntax.proto",
    ],
    # cc_api_version = 2,
    # go_api_version = 2,
    # java_api_version = 2,
    # js_api_version = 2,
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:struct_proto",
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_proto_library(
    name = "syntax_protos_go_proto",
    deps = [":syntax_protos"],
)

alias(
    name = "syntax_go_proto",
    actual = ":syntax_protos_go_proto",
)

py_proto_library(
    name = "syntax_py_pb2",
    # api_version = 2,
    deps = [":syntax_protos"],
)

java_proto_library(
    name = "syntax_java_proto",
    deps = [":syntax_protos"],
)

java_lite_proto_library(
    name = "syntax_java_proto_lite",
    deps = [":syntax_protos"],
)

jspb_proto_library(
    name = "syntax_jspb_proto",
    deps = [":syntax_protos"],
)
