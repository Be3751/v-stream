package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	vstream "github.com/Be3751/v-stream"
	"github.com/Be3751/v-stream/pkg/pb"
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

func (s *myServer) ReceiveVideo(req *pb.VideoRequest, srv pb.VideoStream_ReceiveVideoServer) error {
	root, err := vstream.GetRootPath()
	if err != nil {
		return err
	}
	fmt.Println(root)
	fileName := fmt.Sprintf("piyo%s.mp4", req.VideoId)
	f, err := os.Open(fmt.Sprintf("%s/%s", root, fileName))
	if err != nil {
		return fmt.Errorf("failed to open a file: %w", err)
	}
	for {
		videoBytes := make([]byte, 1024)
		_, err := f.Read(videoBytes)
		if err == io.EOF {
			srv.Send(&pb.VideoResponse{
				Name:  "",
				Video: videoBytes,
			})
			break
		} else if err != nil {
			return fmt.Errorf("failed to read video bytes: %w", err)
		}
		srv.Send(&pb.VideoResponse{
			Name:  "",
			Video: videoBytes,
		})
	}
	return nil
}
