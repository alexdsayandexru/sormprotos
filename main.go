package main

import (
	"context"
	"fmt"
	"github.com/alexdsayandexru/sormprotos/gen/go/sorm"
	"google.golang.org/grpc"
	"net"
)

type SormDataCollectionServiceImpl struct {
	sorm.UnimplementedSormDataCollectionServiceServer
}

func (SormDataCollectionServiceImpl) Send(ctx context.Context, r *sorm.SendRequest) (*sorm.SendResponse, error) {
	fmt.Println("SendRequest:", r.Text, r.Subtext)
	response := &sorm.SendResponse{Text: r.Subtext, Subtext: r.Text}
	return response, nil
}

func main() {
	fmt.Println("Initializing...")
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	server := grpc.NewServer()
	var sormDataCollectionServiceImpl SormDataCollectionServiceImpl
	sorm.RegisterSormDataCollectionServiceServer(server, sormDataCollectionServiceImpl)
	server.Serve(listen)
}
