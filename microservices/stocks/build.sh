#!/bin/bash

env GOOS=linux go build stocks.go
docker build -t quay.io/sevii/ll-stocks .
docker run -p 8075:8075 -d quay.io/sevii/ll-stocks
docker push quay.io/sevii/ll-stocks
