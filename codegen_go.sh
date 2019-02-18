#!/bin/bash

protoc -I ./ --go_out=plugins=grpc:src/go/ proto/v1/*.proto
cp -f $GOPATH/src/github.com/QHasaki/protobuf/src/go/v1/*.go $GOPATH/src/github.com/QHasaki/Server/protocol/v1/