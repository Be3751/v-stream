package main

import (
	"context"
	"fmt"
	"log"

	vstream "github.com/Be3751/v-stream"
	"github.com/Be3751/v-stream/internal/config"
	"github.com/Be3751/v-stream/internal/grpc_client"
	"github.com/Be3751/v-stream/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := clientConnect("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewVideoStreamClient(conn)

	root, err := vstream.GetRootPath()
	if err != nil {
		log.Fatal("failed to get the root path in this project")
		return
	}
	config := config.ClientConfig{Root: root}

	ctx := context.Background()

	myClient := grpc_client.NewMyClient(config, client)
	myClient.RequestVideo(ctx, "_re")
}

func clientConnect(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}
	return conn, nil
}
