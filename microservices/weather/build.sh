#!/bin/bash

env GOOS=linux go build weather.go
docker build -t ll/weather .
docker run -p 8070:8070 -d ll/weather
