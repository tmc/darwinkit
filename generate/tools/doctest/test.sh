#!/bin/bash
set -e

# Build the test tool
go build -o doctest

# Create test output directory
mkdir -p testdata/out

# Run test on sample documentation
for doc in testdata/*.json; do
    if [ -f "$doc" ]; then
        echo "Testing $doc..."
        ./doctest -doc "$doc" -out testdata/out
    fi
done

# Clean up
rm -f doctest 