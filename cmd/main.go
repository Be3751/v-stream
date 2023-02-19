package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Be3751/v-stream/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()

	pb.RegisterVideoStreamServer(srv, NewMyServer())

	go func() {
		log.Printf("start gRPC server port: %d", port)
		srv.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	srv.GracefulStop()
}

type myServer struct {
	pb.UnimplementedVideoStreamServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) ReceiveVideo(*pb.SendRequest, pb.VideoStream_ReceiveVideoServer) error {
	return nil
}
