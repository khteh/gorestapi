#!/bin/bash
#$(aws ecr get-login --no-include-email)
go build
docker build -t gorestapi .
docker tag gorestapi:latest khteh/gorestapi:latest
docker push khteh/gorestapi:latest
