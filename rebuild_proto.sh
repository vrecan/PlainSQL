#!/bin/bash
export GOPATH=~/go/
export PATH=$PATH:$GOPATH/bin
protoc -I pkg/request/ pkg/request/request.proto --go_out=plugins=grpc:pkg/request
