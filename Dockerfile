FROM golang:1.21-alpine

WORKDIR /go/src
COPY . .
RUN apk update && apk add protobuf & \
    go get google.golang.org/protobuf/cmd/protoc-gen-go@latest & \
    go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

CMD ["go", "run", "cmd/main.go"]