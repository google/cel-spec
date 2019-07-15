#!/bin/bash

# Variant of conftest.sh to always succeed, but still report errors in test output.
# Necessary to work around a limitation in GCB for the conformance test dashboard
# (could change when GCB feature to allow passing builds when steps fail is released).

(exec "$@")
exit 0
