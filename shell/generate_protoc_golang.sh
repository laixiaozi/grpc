#!/bin/bash
echo "############################################"
echo "#"
echo "# 脚本使用protoc 命令,请确保安装系统已经安装 Protocol Buffers 来生成  {golang}  的 grpc  pb文件\n"
echo "# 下载地址: https://github.com/protocolbuffers/protobuf/releases \n"
echo "# 下载链接(供参考): https://github.com/protocolbuffers/protobuf/releases/download/v3.11.3/protobuf-all-3.11.3.tar.gz \n"
echo "# 下载链接(源码): https://github.com/protocolbuffers/protobuf/archive/v3.11.3.tar.gz \n"
echo "# 参考文档: https://developers.google.com/protocol-buffers/docs/proto3 \n"
echo "# 脚本只能在当前目录下使用请不要随意移动!!!!"
source /etc/profile
echo "# 当前版本:";
protoc --version
echo "#"
echo "# 当前脚本参考: https://github.com/grpc/grpc-go/blob/master/codegen.sh"
echo "# grpc-go插件下载地址:https://github.com/grpc/grpc-go"
echo "#"
echo "#"
echo "############################################"

basePath=$(dirname $(pwd))
outPath=$basePath"/pbs"
inputPath=$basePath"/protocs/"
proto=$1
protoc --proto_path=$inputPath  --go_out=plugins=grpc:$outPath $proto

