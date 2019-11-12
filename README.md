# grpc_test_go

Go gRPC sample.  

`server/` is gRPC server application.  
First start this application.  

`cliant/` is cliant application.  
When you start this application, it call gRPC of server application.  



## Install protoc

http://google.github.io/proto-lens/installing-protoc.html

```
brew install protobuf
```

## Install gRPC package

```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
github.com/kainlite/grpc-ping/ping
```

## Create Go file from .proto file

```
cd server/ping/
protoc --go_out=plugins=grpc:./ ping.proto
```

## Run applications

1. Run server application

```
cd ./server
go run main.go
```

2. Run cliant application

```
cd ./cliant
go run main.go
```

## References
- [gRPC入門](https://www.slideshare.net/KenjiroKubota/grpc-141520916)
- [Goで始めるgRPC入門](https://qiita.com/marnie_ms4/items/4582a1a0db363fe246f3)
- [gRPCって何？](https://qiita.com/oohira/items/63b5ccb2bf1a913659d6)
