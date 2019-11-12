package main

import (
	"./ping"
	"./service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = ":19003"

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	// 作成したサービスを登録
	pingService := service.PingService{}
	ping.RegisterPingServer(server, &pingService)

	fmt.Printf("[server started] localhost%s", port)
	err = server.Serve(listenPort)
	if err != nil {
		log.Fatalln(err)
	}
}
