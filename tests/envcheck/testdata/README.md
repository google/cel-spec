# Envcheck Conformance Test Data Files

The test data files for the `envcheck` conformance test suite
are described by the `Env` message defined in
[`envcheck.proto`](../../../proto/test/v1/envcheck.proto).

An implementation should generally choose a single data file to test against,
corresponding to the environment that it intends to support.

Over time, we expect the `simple` test suite to incorporate more detailed
tests which will subsume the functionality of this suite.  In the meantime,
this suite serves as an initial coverage test.

The available test files are:

- [`go-0.1.0.textproto`](go-0.1.0.textproto) The go-0.1.0 environment, including
  some deprecated entries for backward compatibility.
