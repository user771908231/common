#!/bin/bash
pwd=`pwd`
cd ./ddproto
rm -rf *.pb.go
cd ../ddprotobuf
protoc --proto_path="." --go_out=plugins=grpc:../ddproto *.proto
cd $pwd
