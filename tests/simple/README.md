# Simple Conformance Tests

The *simple* test suite checks end-to-end functionality of the entire
parse-(check)-eval pipeline.  Test inputs and output are fully specified.
The parse and check phases are expected to complete without error and their
outputs are not individually validated - we only check the final results of
evaluation.

The simple suite is executed by the _driver_ binary in `simple_test.go`.
This binary is invoked with one or more of the following flags:
- `--server` Path to the default ConformanceService server binary, to be used
  when no phase-specific server is specified.
- `--parse_server` Path to the ConformanceService server binary to use for
  parsing.
- `--check_server` Path to the ConformanceService server binary to use for
  checking.
- `--eval_server` Path to the ConformanceService server binary to use for
  evaluation.

The remaining arguments are paths to data files containing the test
inputs and expected outputs.

The driver binary should be run from this directory, looking in
the `testdata/` subdirectory for its data files.

Implementations which use the [bazel](https://bazel.build) build system
can invoke the simple test suite as their own unit tests with the following
entry in a `BUILD.bazel` file, substituting their own conformance server
binary target for `//server/main:cel_server`:

```
sh_test(
    name = "conformance_simple",
    srcs = ["@com_google_cel_spec//tests:conftest.sh"],
    args = [
        "$(location @com_google_cel_spec//tests/simple:simple_test)",
        "--server=$(location //server/main:cel_server)",
        "$(location @com_google_cel_spec//tests/simple:testdata/plumbing.textproto)",
        "$(location @com_google_cel_spec//tests/simple:testdata/basic.textproto)",
        "$(location @com_google_cel_spec//tests/simple:testdata/integer_math.textproto)",
    ],
    data = [
        "@com_google_cel_spec//tests/simple:simple_test",
        "//server/main:cel_server",
        "@com_google_cel_spec//tests/simple:testdata/plumbing.textproto",
        "@com_google_cel_spec//tests/simple:testdata/basic.textproto",
        "@com_google_cel_spec//tests/simple:testdata/integer_math.textproto",
    ],
)
```

Implementations should add more test data files as they increase their
conformance.  See [testdata](testdata) for a description of the individual
test files.
