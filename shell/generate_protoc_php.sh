#!/bin/bash
echo "############################################"
echo "#"
echo "# 脚本使用protoc 命令,请确保安装系统已经安装 Protocol Buffers 来生成  {PHP}  的 grpc  pb文件\n"
echo "# 下载地址: https://github.com/protocolbuffers/protobuf/releases \n"
echo "# 脚本只能在当前目录下使用请不要随意移动!!!!"
source /etc/profile
echo "# 当前版本:";
protoc --version
echo "#"
echo "#"
echo "#"
echo "#"
echo "############################################"

echo "安装grpc php 插件"
#git clone https://github.com/grpc/grpc.git
#cd grpc
#git pull --recurse-submodules && git submodule update --init --recursive
#make
#sudo make install

basePath=$(dirname $(pwd))
outPath=$basePath"/pbs"
inputPath=$basePath"/protocs/"
proto=$1
protoc --proto_path=$inputPath  --php_out=./ --grpc_out=./ --plugin=protoc-gen-grpc=/home/ryan/grpc_php_plugins/grpc/bins/opt/grpc_php_plugin  $proto  #--go_out=plugins=grpc:$outPath $proto

