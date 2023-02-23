package grpc_server

import (
	"fmt"
	"io"
	"os"

	"github.com/Be3751/v-stream/internal/config"
	"github.com/Be3751/v-stream/pkg/pb"
)

var _ pb.VideoStreamServer = (*myServer)(nil)

type myServer struct {
	pb.UnimplementedVideoStreamServer
	config config.ServerConfig
}

func NewMyServer(c config.ServerConfig) pb.VideoStreamServer {
	return &myServer{
		config: c,
	}
}

func (s *myServer) ReceiveVideo(req *pb.VideoRequest, srv pb.VideoStream_ReceiveVideoServer) error {
	fileName := fmt.Sprintf("piyo%s.mp4", req.VideoId)
	f, err := os.Open(fmt.Sprintf("%s/media/in/%s", s.config.Root, fileName))
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
