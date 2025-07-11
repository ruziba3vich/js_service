#!/bin/bash
# chmod +x proto_gen.sh

set -e

PROTO_DIR="./"

OUT_DIR="./genprotos"

mkdir -p $OUT_DIR

protoc -I=$PROTO_DIR \
  --go_out=$OUT_DIR \
  --go-grpc_out=$OUT_DIR \
  $PROTO_DIR/compiler_protos/compiler.proto

echo "protos generated successfully"
