
To build for linux

env GOOS=linux go build edge.go

docker built -t ll/edge
docker run -p 8080:8080 ll/edge


Test with:
http POST localhost:8080/report Username=Admin Level=1
http -a Admin:admin localhost:8080