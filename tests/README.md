# CEL Conformance Tests

The CEL conformance test suites provide a complementary specification
of the CEL language in the form of executable software. All CEL
implementations which pass the conformance tests should give valid CEL
expressions the same meaning.

As much as possible, the tests are driven by data files. Implementations
are expected to read the `tests/simple/testdata` files as inputs to the
parse, check, and evaluation behaviors of a given implementation. Reference
implementations may be found in the following locations:

*
*
*

Each CEL implementation should also have its own unit tests and benchmarks
for testing subcomponents, its API, and other implementation-dependent
features.
