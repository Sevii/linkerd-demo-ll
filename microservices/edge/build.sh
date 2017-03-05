#!/bin/bash

env GOOS=linux go build edge.go
docker build -t ll/edge .
docker run -p 8080:8080 -d ll/edge
