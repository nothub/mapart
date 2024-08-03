#!/usr/bin/env sh

set -eu
cd "$(dirname "$(realpath "$0")")/.."

set -x

go test -v -vet='all' ./...
