#! /bin/bash
go mod init grpc-practice
protoc -I. --go_out=. --go-grpc_out=. proto/*.proto 
go mod tidy