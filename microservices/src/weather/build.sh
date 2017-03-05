#!/bin/bash

env GOOS=linux go build weather.go
docker build -t quay.io/sevii/ll-weather:latest .
#docker run -p 8070:8070 -d quay.io/sevii/ll-weather
docker push quay.io/sevii/ll-weather