load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

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

cc_proto_library(
    name = "value_cc_proto",
    deps = [":value_protos"],
)

go_proto_library(
    name = "value_go_proto",
    importpath = "github.com/google/cel-spec/proto/v1/value",
    proto = ":value_protos",
)

java_proto_library(
    name = "value_java_proto",
    deps = [":value_protos"],
)

java_lite_proto_library(
    name = "value_java_proto_lite",
    deps = [":value_protos"],
)

proto_library(
    name = "cel_service_proto",
    srcs = [
        "cel-service.proto",
    ],
    deps = [
        ":eval_protos",
        ":syntax_protos",
        ":value_protos",
        "//proto/checked/v1:checked_protos",
        "@com_google_googleapis//:rpc_status",
    ],
)

go_proto_library(
    name = "cel_service_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/google/cel-spec/proto/v1/cel_service",
    proto = ":cel_service_proto",
    deps = [
        ":eval_go_proto",
        ":syntax_go_proto",
        ":value_go_proto",
        "//proto/checked/v1:checked_go_proto",
        "@com_google_googleapis//:rpc_status_go_proto",
    ],
)

proto_library(
    name = "eval_protos",
    srcs = [
        "eval.proto",
    ],
    # cc_api_version = 2,
    # java_api_version = 2,
    # js_api_version = 2,
    deps = [
        ":value_protos",
        "@com_google_googleapis//:rpc_status",
    ],
)

go_proto_library(
    name = "eval_go_proto",
    importpath = "github.com/google/cel-spec/proto/v1/eval",
    proto = ":eval_protos",
    deps = [
        ":value_go_proto",
        "@com_google_googleapis//:rpc_status_go_proto",
    ]
)

java_proto_library(
    name = "eval_java_proto",
    deps = [":eval_protos"],
)

java_lite_proto_library(
    name = "eval_java_proto_lite",
    deps = [":eval_protos"],
)

proto_library(
    name = "syntax_protos",
    srcs = [
        "syntax.proto",
    ],
    # cc_api_version = 2,
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
    name = "syntax_go_proto",
    importpath = "github.com/google/cel-spec/proto/v1/syntax",
    proto = ":syntax_protos",
)

java_proto_library(
    name = "syntax_java_proto",
    deps = [":syntax_protos"],
)

java_lite_proto_library(
    name = "syntax_java_proto_lite",
    deps = [":syntax_protos"],
)
