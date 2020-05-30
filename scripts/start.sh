#!/bin/sh

set -e

[ -z "$DEBUG" ] || set -x

cd cmd/c4

echo "===> Generate wire...\n"

wire

echo "===> Generate build...\n"

go build -o bin/application main.go wire_gen.go

echo "===> Start local...\n"

./bin/application