# Simple Conformance Test Data Files

The test data files for the `simple` conformance test suite
are described by the `SimpleTestFile` message defined in
[`simple.proto`](../../../proto/test/v1/simple.proto).
See the documentation in that file for the meaning of the various fields.

A run of the simple tests chooses one or more data files and attempts
to run all the tests in those files.  Test files are organized so that
implementations can implement a prescribed subset of CEL functionality.
For instance, implementations which don't support macros can avoid testing
against the `macros.textproto` file.

The available test files are:

- [`plumbing.textproto`](plumbing.textproto) Checks the basics of the CelService
  protocol to ensure that the server is implemented correctly.

- [`basic.textproto`](basic.textproto) Checks the most basic operations that
  all CEL implementations should support:
  - literals of various types;
  - variables of various types.

- [`comparisons.textproto`](comparisons.textproto) Checks the standard functions
  that return a boolean value.

- [`conversions.textproto`](conversions.textproto) Checks conversions between
  types and type tests.

- [`dynamic.textproto`](dynamic.textproto) Checks the automatic conversions
  associated with the well-known protobuf messages.

- [`enums.textproto`](enums.textproto) Checks handling of protobuf enums.

- [`fields.textproto`](fields.textproto) Checks field selection in messages
  and maps.

- [`fp_math.textproto`](fp_math.textproto) Checks floating-point arithmetic.

- [`integer_math.textproto`](integer_math.textproto) Checks integer arithmetic.

- [`lists.textproto`](lists.textproto) Checks list operations.

- [`logic.textproto`](logic.textproto) Checks special logical operators.

- [`macros.textproto`](macros.textproto) Checks use of CEL macros.

- [`namespace.textproto`](namespace.textproto) Checks use of namespaces and
  qualified identifiers.

- [`parse.textproto`](parse.textproto) End-to-end tests of parsing. More
  detailed parsing tests will the subject of a later conformance suite.

- [`proto2.textproto`](proto2.textproto) Checks use of protocol buffers version
  2.

- [`proto3.textproto`](proto3.textproto) Checks use of protocol buffers version
  3.

- [`string.textproto`](string.textproto) Checks functions on strings.

- [`timestamps.textproto`](timestamps.textproto) Checks `timestamp` and `duration`
  values and operations.

- [`unknowns.textproto`](unknowns.textproto) Checks evaluation where some
  inputs are marked as unknown.
