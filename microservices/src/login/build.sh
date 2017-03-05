#!/bin/bash

env GOOS=linux go build login.go
docker build -t quay.io/sevii/ll-login:latest .
#docker run -p 8050:8050 -d quay.io/sevii/ll-login
docker push quay.io/sevii/ll-login
