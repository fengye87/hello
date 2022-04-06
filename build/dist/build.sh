#!/bin/bash
set -euxo pipefail

arch=$(uname -m)
version=$(git describe --long --tags --dirty | awk '{print substr($1,2)}')

mkdir -p out/dist/hello-$version.linux-$arch
rm -rf out/dist/hello-$version.linux-$arch/*

for name in hello helloctl; do
    go build -o out/dist/hello-$version.linux-$arch/$name cmd/$name/main.go
done

tar -zcf out/dist/hello-$version.linux-$arch.tar.gz -C out/dist hello-$version.linux-$arch
