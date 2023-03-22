#!/bin/bash

# 监听UDP端口
nc -u -l -p 12345 |
# 读取数据并输出到终端上
while read data; do
  echo "Received: $data"
done
