#!/bin/sh
set -euxo pipefail

version=$(git describe --long --tags --dirty | awk '{print substr($1,2)}')

for name in hello helloctl; do
    DOCKER_BUILDKIT=1 docker build -f build/image/Dockerfile.$name -t fengye87/$name:$version .
done
