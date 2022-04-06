FROM golang:1.17-alpine

RUN apk add protobuf-dev curl

RUN mkdir -p /usr/include/google/api && \
    for name in annotations http; do \
        curl -L https://raw.githubusercontent.com/googleapis/googleapis/536a2d8a2435fadffbe6dc120219623c4e54b9a2/google/api/$name.proto -o /usr/include/google/api/$name.proto; \
    done

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go install \
        github.com/golang/mock/mockgen \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
