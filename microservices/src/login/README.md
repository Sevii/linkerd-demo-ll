
To build for linux

env GOOS=linux go build login.go

docker built -t ll/login
docker run -p 8080:8080 ll/login


Test with:
http POST localhost:8050/login Username=Stern Password=mrpaulstern
