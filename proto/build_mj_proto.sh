#!/bin/bash
pwd=`pwd`
cd ./ddproto
rm -rf ./mjproto/*.pb.go
cd ../ddprotobuf/mjproto
protoc --proto_path="." --go_out=plugins=grpc:../../ddproto/mjproto *.proto
cd $pwd
