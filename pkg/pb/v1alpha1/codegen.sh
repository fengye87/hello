#!/bin/sh
set -euo pipefail

for name in hello; do
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --openapiv2_out=. $name.proto
done
