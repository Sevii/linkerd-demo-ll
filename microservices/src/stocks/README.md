
To build for linux

env GOOS=linux go build stocks.go

docker built -t ll/stocks
docker run -p 8080:8080 ll/stocks
