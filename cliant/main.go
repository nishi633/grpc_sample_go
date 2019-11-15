package main

import (
	"../server/ping" // Import ping package made from ping.proto
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
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
	if err != nil {
		log.Fatal("Hello request error:", err)
		os.Exit(1)
	}
	fmt.Printf("result:%#v \n", res)

	message2 := &ping.GoodbyRequest{}
	res2, err := client.Goodby(context.TODO(), message2)
	if err != nil {
		log.Fatal("Goodby request error:", err)
		os.Exit(1)
	}
	fmt.Printf("result:%#v \n", res2)
}
