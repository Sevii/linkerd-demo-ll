#!/bin/bash

env GOOS=linux go build stocks.go
docker build -t ll/stocks .
docker run -p 8075:8075 -d ll/stocks
