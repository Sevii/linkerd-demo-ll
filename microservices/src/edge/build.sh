#!/bin/bash

env GOOS=linux go build edge.go
docker build -t quay.io/sevii/ll-edge:latest .
#docker run -p 8080:8080 -d quay.io/sevii/ll-edge
docker push quay.io/sevii/ll-edge