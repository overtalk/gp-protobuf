#!/bin/bash

protoc -I ./ --go_out=plugins=grpc:src/ proto/*.proto

cp -f $GOPATH/src/github.com/qinhan-shu/gp-protobuf/src/go/*.go $GOPATH/src/github.com/qinhan-shu/gp-server/protocol/