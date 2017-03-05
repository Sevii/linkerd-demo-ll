#!/bin/bash

env GOOS=linux go build reports.go
docker build -t ll/reports .
docker run -p 8055:8055 -d ll/reports
