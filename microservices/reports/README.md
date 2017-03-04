
To build for linux

env GOOS=linux go build reports.go

docker built -t ll/reports
docker run -p 8080:8080 ll/reports


Test with:
http POST localhost:8080/report Username=Admin Level=1
