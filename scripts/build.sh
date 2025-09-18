#!/usr/bin/env sh

set -eu
cd "$(dirname "$(realpath "$0")")/.."

build() (

    file="dist/mapart_${1}_${2}"
    if test "$1" = "windows"; then
        file="${file}.exe"
    fi

    # build static binary
    GOOS="$1" GOARCH="$2" go build \
        -ldflags "-s -w" \
        -o "${file}" \
        .

    # compress with upx
    # ( except for mac because https://github.com/upx/upx/issues/612 )
    if test "$1" != "darwin"; then
        upx --best --lzma \
            --no-color \
            --no-progress \
            --no-time \
            "${file}"
    fi

)

set -x

rm    -rf "dist"
mkdir -p  "dist"

go generate .

build linux amd64
build linux arm64
build darwin amd64
build darwin arm64
build windows amd64
