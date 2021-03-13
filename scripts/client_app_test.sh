#!/bin/bash

export GRPC_CONNECTION='localhost:8080'
export USER_IDS='1,2,3,4,5,6,7,8,9,10'
go build -o client cmd/grpc/client/main.go
./client