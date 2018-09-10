package(
    default_compatible_with = ["//buildenv/target:appengine"],
    default_visibility = ["//visibility:public"],
)

proto_library(
    name = "value_proto",
    srcs = [
        "value.proto",
    ],
    cc_api_version = 2,
    js_api_version = 2,
    deps = [
        "//google/protobuf:any",
        "//google/protobuf:duration",
        "//google/protobuf:empty",
        "//google/protobuf:struct",
        "//google/protobuf:timestamp",
    ],
)

go_proto_library(
    name = "value_proto_go_proto",
    deps = [":value_proto"],
)

java_proto_library(
    name = "value_java_proto",
    deps = [":value_proto"],
)

java_lite_proto_library(
    name = "value_java_proto_lite",
    deps = [":value_proto"],
)

jspb_proto_library(
    name = "value_jspb_proto",
    deps = [
        ":value_proto",
    ],
)

proto_library(
    name = "eval_proto",
    srcs = [
        "eval.proto",
    ],
    cc_api_version = 2,
    js_api_version = 2,
    deps = [
        ":value_proto",
        "//google/rpc:status",
    ],
)

go_proto_library(
    name = "eval_proto_go_proto",
    deps = [":eval_proto"],
)

java_proto_library(
    name = "eval_java_proto",
    deps = [":eval_proto"],
)

java_lite_proto_library(
    name = "eval_java_proto_lite",
    deps = [":eval_proto"],
)

jspb_proto_library(
    name = "eval_jspb_proto",
    deps = [
        ":eval_proto",
    ],
)

proto_library(
    name = "syntax_proto",
    srcs = [
        "syntax.proto",
    ],
    cc_api_version = 2,
    js_api_version = 2,
    deps = [
        "//google/protobuf:any",
        "//google/protobuf:duration",
        "//google/protobuf:empty",
        "//google/protobuf:struct",
        "//google/protobuf:timestamp",
    ],
)

go_proto_library(
    name = "syntax_proto_go_proto",
    deps = [":syntax_proto"],
)

cc_proto_library(
    name = "syntax_cc_proto",
    deps = [":syntax_proto"],
)

py_proto_library(
    name = "syntax_py_pb2",
    api_version = 2,
    deps = [":syntax_proto"],
)

java_proto_library(
    name = "syntax_java_proto",
    deps = [":syntax_proto"],
)

java_lite_proto_library(
    name = "syntax_java_proto_lite",
    deps = [":syntax_proto"],
)

jspb_proto_library(
    name = "syntax_jspb_proto",
    deps = [
        ":syntax_proto",
    ],
)

proto_library(
    name = "checked_proto",
    srcs = [
        "checked.proto",
    ],
    cc_api_version = 2,
    js_api_version = 2,
    deps = [
        ":syntax_proto",
        "//google/protobuf:empty",
        "//google/protobuf:struct",
    ],
)

go_proto_library(
    name = "checked_proto_go_proto",
    deps = [":checked_proto"],
)

cc_proto_library(
    name = "checked_cc_proto",
    deps = [":checked_proto"],
)

py_proto_library(
    name = "checked_py_pb2",
    api_version = 2,
    deps = [":checked_proto"],
)

java_proto_library(
    name = "checked_java_proto",
    deps = [":checked_proto"],
)

java_lite_proto_library(
    name = "checked_java_proto_lite",
    deps = [":checked_proto"],
)

jspb_proto_library(
    name = "checked_jspb_proto",
    deps = [
        ":checked_proto",
    ],
)

proto_library(
    name = "explain_proto",
    srcs = [
        "explain.proto",
    ],
    cc_api_version = 2,
    js_api_version = 2,
    deps = [
        ":value_proto",
    ],
)

go_proto_library(
    name = "explain_proto_go_proto",
    deps = [":explain_proto"],
)

java_proto_library(
    name = "explain_java_proto",
    deps = [":explain_proto"],
)

java_lite_proto_library(
    name = "explain_java_proto_lite",
    deps = [":explain_proto"],
)

jspb_proto_library(
    name = "explain_jspb_proto",
    deps = [
        ":explain_proto",
    ],
)
