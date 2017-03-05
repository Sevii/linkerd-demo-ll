#!/bin/bash

env GOOS=linux go build reports.go
docker build -t quay.io/sevii/ll-report .
docker run -p 8055:8055 -d quay.io/sevii/ll-report
docker push quay.io/sevii/ll-report

