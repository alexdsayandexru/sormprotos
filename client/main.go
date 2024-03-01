package main

import (
	"context"
	"fmt"
	"github.com/alexdsayandexru/sormprotos/gen/go/sorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	client := sorm.NewSormDataCollectionServiceClient(conn)

	response, err := client.Send(context.Background(), &sorm.SendRequest{Text: "hello", Subtext: "world"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
