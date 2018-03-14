package(
    # default_compatible_with = ["//buildenv/target:appengine"],
    default_visibility = ["//visibility:public"],
)

licenses(["notice"])  # Apache 2.0

exports_files(["LICENSE"])

proto_library(
    name = "checked_protos",
    srcs = [
        "checked.proto",
    ],
    # cc_api_version = 2,
    # go_api_version = 2,
    # java_api_version = 2,
    # js_api_version = 2,
    deps = [
        "//proto/v1:syntax_protos",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:struct_proto",
    ],
)

java_proto_library(
    name = "checked_java_proto",
    deps = [":checked_protos"],
)

java_lite_proto_library(
    name = "checked_java_proto_lite",
    deps = [":checked_protos"],
)
