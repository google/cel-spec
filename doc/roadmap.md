# CEL Roadmap

CEL is under active development with improvements underway at all levels of the
stack. We are also working to make it easier to adopt and use CEL, including
testing tools and cases to ensure the interoperability and correctness of all
implementations.

This page is a summary of the CEL release history and plans for future releases.
For details on prior releases, see individual release notes and announcements.
For future releases, **do not consider dates and feature lists firm**.

Feedback and feature requests should be made on the buglist for either the
language overall on [cel-spec GitHub issues][], or the Go implementation on
[cel-go GitHub issues][]. Or ask a question or propose a change on the
[cel-go-discuss][] public forum.

[cel-spec GitHub issues]: https://github.com/google/cel-spec/issues
[cel-go GitHub issues]: https://github.com/google/cel-go/issues
[cel-go-discuss]: https://groups.google.com/forum/#!forum/cel-go-discuss

## Releases

### CEL 0.1

Feb 2019.
\[[code](https://github.com/google/cel-go/releases/tag/v0.1.0)\]
\[[announcement](https://groups.google.com/d/topic/cel-go-discuss/tk70TLMqcOo/discussion)\]

Features

- Formal language specification, see the [cel-spec repo](https://github.com/google/cel-spec).
- Conformance test suite. The driver and test cases are part of the language
    specification, with drivers for each implementation. Described in more
    detail [here](https://github.com/google/cel-spec/tests).
- Implementation in Golang only ([source](https://github.com/google/cel-go))
  - Parse and typecheck
  - Attribute binding
  - Evaluation
  - Performance: sub-microsecond performance for basic comparison operations

### CEL 0.2

Spring 2019.

### CEL 0.3

Summer 2019

### CEL 1.0

End-2019