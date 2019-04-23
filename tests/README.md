# CEL Conformance Tests

The CEL conformance test suites provide a complementary specification
of the CEL language in the form of executable software.  All CEL
implementations which pass the conformance tests should give valid CEL
expressions the same meaning.

A language-independent gRPC protocol allows the tests to run against CEL
components in any language which implement a server for the protocol.

As much as possible, the tests are driven by data files.  Implementations
which do not provide a ConformanceService server are welcome to write
their own drivers which read the same `testdata` files.

## Test Suites

We currently have the following test suites:
- *[simple](simple)* The _simple_ tests check the end-to-end
  parse-(check)-eval pipeline for fully-specified expressions,
  inputs, and output without parse or check failures.
- *[envcheck](envcheck)* The _envcheck_ suite confirms that the checker and
  runtime support a set of functions and overloads (though we don't check
  the behavior of these functions in this suite).
- Test suites more suitable for other kinds of validation may be introduced
  later.

## Integrating with the Tests

Implementations should write a gRPC server for the
[`ConformanceService`
API](https://github.com/googleapis/googleapis/blob/master/google/api/expr/v1alpha1/conformance_service.proto)
See [the cel-go server](https://github.com/google/cel-go/tree/master/server)
as an example.

The server should be available as an executable file.  When invoked without
without arguments, the server should listen on an arbitrary available TCP
port on the local loopback address (`127.0.0.1` or `::1`), then write its
address on `stdout` in the format

> `Listening on `_address_`:`_port_

See the [celrpc
library](https://github.com/google/cel-spec/tree/master/tools/celrpc) for an
example of implementing this protocol.

An implementation with only a subset of the execution phases (particularly
runtime-only implementations without a parser or checker) is free to
support only the corresponding API methods.  The test suite driver can be
instructed to use a different implementation for the unsupported methods.

Each conformance test suite will contain specific integration instructions.

Implementations which use the [bazel](https://bazel.build) build system
should copy the relevant entries out of the `WORKSPACE` file at the root
of this repository.

Each CEL implementation should also have its own unit tests and benchmarks
for testing subcomponents, its API, and other implementation-dependent
features.
