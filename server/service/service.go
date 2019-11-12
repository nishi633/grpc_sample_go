package service

import (
	"../ping"
	"fmt"
	"golang.org/x/net/context"
)

type PingService struct {
}

func (s *PingService) Hello(ctx context.Context, req *ping.HelloRequest) (*ping.HelloResponse, error) {
	// reqでリクエストを受ける

	toMessage := req.GetToMessage() // メッセージを取得する
	fmt.Println(toMessage)
	// 返却メッセージの作成
	response := ping.HelloResponse{
		ResMessage: "I hear " + toMessage,
	}

	return &response, nil
}
