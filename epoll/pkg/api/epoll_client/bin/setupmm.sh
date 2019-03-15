#!/bin/bash
## 这个脚本使用Docker在不同的网络命名空间产生多个client实例.
## 这样才能避免source port的限制，在一台机器上才能创建百万的连接.
##
## 用法: ./connect <connections> <number of clients> <server ip>
## Server IP 通常是 Docker gateway IP address, 缺省是 172.17.0.1

CONNECTIONS=$1
IP=$2

DATE=`date -d "+2 minutes" +"%FT%T %z"`

docker run -v $(pwd)/epoll_client:/client --name 1mclient_0 -d alpine ./client -conn=${CONNECTIONS} -ip=${IP}  -sm "${DATE}"