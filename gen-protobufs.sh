#!/usr/bin/env bash
set -eu

# Generates the protocol buffers in the 'service' directory. Requires 
# installation of the protocol buffer compiler and runtime.

# -----------------
# Protobuf Compiler
# -----------------
# https://github.com/google/protobuf

# -------------------
# Protobuf Go Runtime
# -------------------
# https://github.com/golang/protobuf

protoc --go_out=plugins=grpc:. "service/*.proto"
