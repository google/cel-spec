#!/bin/bash
echo "Running test binary as " "$@"
echo "Current directory " "$(pwd)"
echo "Path is" "$PATH"
echo "Server binary:" "$(ls -l server/main/cel_server)"
echo "Contents are" "$(find -L -type f | xargs namei)"
echo "Basic textproto:"
echo =====
# cat external/com_google_cel_spec/tests/simple/testdata/basic.textproto || echo no such file
echo =====
exec "$@"
