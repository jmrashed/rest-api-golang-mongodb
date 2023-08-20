#!/bin/bash

# This is a simple build script for a Go application.
# It compiles the source code and generates an executable binary.

# Set the name of the output binary
OUTPUT_BINARY="myapp"

# Compile the Go source code
go build -o "$OUTPUT_BINARY" cmd/main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Build successful! Executable binary: $OUTPUT_BINARY"
else
    echo "Build failed."
fi
