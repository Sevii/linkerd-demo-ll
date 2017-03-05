#!/bin/bash

env GOOS=linux go build login.go
docker build -t ll/login .
docker run -p 8050:8050 -d ll/login
