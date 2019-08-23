# Envcheck Conformance Tests

The *envcheck* test suite checks that a runtime supports a set of functions
and constants.  No effort is made to check the actual behavior of the
function or constant - it's just a check whether it's missing.  This is
useful as a first pass to see which detailed test suites should be used
for a runtime.

The envcheck suite is executed by the _driver_ binary in `envcheck_test.go`.
This binary is invoked with one or more of the following flags:
- `--server` Path to the default ConformanceService server binary, which must
  support the Eval phase.

The remaining arguments are paths to data files containing the declarations.

The driver binary should be run from this directory, looking in
the `testdata/` subdirectory for its data files.

Implementations which use the [bazel](https://bazel.build) build system
can invoke the simple test suite as their own unit tests with the following
entry in a `BUILD.bazel` file, substituting their own conformance server
binary target for `//server/main:cel_server`:

```
sh_test(
    name = "conformance_envcheck",
    srcs = ["@com_google_cel_spec//tests:conftest.sh"],
    args = [
        "$(location @com_google_cel_spec//tests/envcheck:envcheck_test)",
        "--server=$(location //server/main:cel_server)",
        "$(location @com_google_cel_spec//tests/envcheck:testdata/go-0.1.0.textproto)",
    ],
    data = [
        "@com_google_cel_spec//tests/envcheck:envcheck_test",
        "//server/main:cel_server",
        "@com_google_cel_spec//tests/envcheck:testdata/go-0.1.0.textproto",
    ],
)
```

Implementations should typically check against a single appropriate target to
see what functions or constants might have slipped between the cracks and
implement them.  Other test suites will have more detailed tests that actually
verify behavior and contents.  These more detailed suites should be used
in the long run.

See [testdata](testdata) for a description of the individual test files.
