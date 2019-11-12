package main

import (
	"../server/ping" // Import ping package made from ping.proto
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := ping.NewPingClient(conn)
	message := &ping.HelloRequest{ToMessage: "はろー"}
	res, err := client.Hello(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
