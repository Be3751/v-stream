package main

import (
	"context"
	"log"

	"github.com/Be3751/v-stream/internal/uni_stream_client"
	"github.com/Be3751/v-stream/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	client := pb.NewVideoStreamClient(conn)
	myClient := uni_stream_client.NewMyClient(client)
	myClient.RequestVideo(context.Background(), "_re")
}
