package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	vstream "github.com/Be3751/v-stream"
	"github.com/Be3751/v-stream/internal/config"
	"github.com/Be3751/v-stream/internal/grpc_server"
	"github.com/Be3751/v-stream/pkg/pb"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()

	root, err := vstream.GetRootPath()
	if err != nil {
		return
	}
	config := config.ServerConfig{Root: root}

	myServer := grpc_server.NewMyServer(config)
	pb.RegisterVideoStreamServer(server, myServer)

	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server port: %d", port)
		server.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
