
To build for linux

env GOOS=linux go build weather.go

docker built -t ll/weather
docker run -p 8080:8080 ll/weather
