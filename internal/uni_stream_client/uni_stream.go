package uni_stream_client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	vstream "github.com/Be3751/v-stream"
	"github.com/Be3751/v-stream/pkg/pb"
)

type MyClient interface {
	RequestVideo(ctx context.Context, videoId string)
}

var _ MyClient = (*myClient)(nil)

type myClient struct {
	vclient pb.VideoStreamClient
}

func NewMyClient(v pb.VideoStreamClient) MyClient {
	return &myClient{vclient: v}
}

func (c *myClient) RequestVideo(ctx context.Context, videoId string) {
	req := &pb.VideoRequest{
		VideoId: videoId,
	}
	stream, err := c.vclient.ReceiveVideo(ctx, req)
	if err != nil {
		return
	}

	root, err := vstream.GetRootPath()
	if err != nil {
		return
	}
	fileName := "download.mp4"
	f, err := os.Create(fmt.Sprintf("%s/media/out/%s", root, fileName))
	if err != nil {
		return
	}

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("all the responses have already received.")
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}

		_, err = f.Write(res.Video)
		if err != nil {
			return
		}
		fmt.Println(res)
	}

	return
}
