#!/bin/bash

# Travis tests, consisting of tests of the cel-spec repository,
# plus conformance tests of CEL implementations.
#
# We'll configure travis with a test matrix based upon multiple env settings.
#
# env:
#  matrix:
#    - CEL_NORMAL_TESTS=1
#    - TEST_TARGET=@cel_go//conformance:basic TRAVIS_ALLOW_FAILURE=false
#    - TEST_TARGET=@cel_go//conformance:advanced TRAVIS_ALLOW_FAILURE=true
#
# Each row with either set CEL_NORMAL_TESTS to nonempty, which runs the
# normal tests, otherwise sets the following in one row:
#
# TEST_TARGET=<bazel target>
#   a bazel target that runs a test, specifying the test server(s)
#   and test files
#
# TRAVIS_ALLOW_FAILURE=true/false
#   whether it's okay to not pass the test suite yet
#
# This runs the conformance tests, but only from cron, so a normal
# test won't have to run everything.

# If not running conformance tests, just build.
if [ -n $CEL_NORMAL_TESTS ]
then
  exec bazel build ...
fi

# Only run conformance tests if invoked from cron, and if there's a target.
if [ ${TRAVIS_EVENT_TYPE} = cron -a -n ${TEST_TARGET} ]
then
  exec bazel test ${TEST_TARGET}
fi
