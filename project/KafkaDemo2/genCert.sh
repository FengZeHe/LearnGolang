#!/bin/bash

# 创建目录
mkdir -p certs
cd certs

# 生成根证书私钥
openssl genrsa -out ca-key.pem 2048

# 生成根证书
openssl req -x509 -new -nodes -key ca-key.pem -days 3650 -out ca-cert.pem -subj "/CN=KafkaCA"

# 生成 Kafka 服务器私钥
openssl genrsa -out server-key.pem 2048

# 生成 Kafka 服务器证书签名请求
openssl req -new -key server-key.pem -out server-csr.pem -subj "/CN=kafka"

# 使用根证书签署 Kafka 服务器证书
openssl x509 -req -in server-csr.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -days 3650

# 生成客户端私钥
openssl genrsa -out client-key.pem 2048

# 生成客户端证书签名请求
openssl req -new -key client-key.pem -out client-csr.pem -subj "/CN=client"

# 使用根证书签署客户端证书
openssl x509 -req -in client-csr.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -days 3650

cd ..