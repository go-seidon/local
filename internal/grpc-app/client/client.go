package main

import (
	"context"
	"io"
	"os"

	"github.com/go-seidon/local/internal/grpc-app/pb"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, _ := grpc.Dial("localhost:8005", opts...)
	client := pb.NewUploadServiceClient(conn)
	grpcClient := &Client{client: client}
	grpcClient.Upload(context.Background())
	defer conn.Close()
}

type Client struct {
	client pb.UploadServiceClient
}

func (o *Client) Upload(ctx context.Context) (pb.UploadService_UploadClient, error) {
	file, err := os.Open("/Users/yudi_s/Downloads/grpc.png")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stream, err := o.client.Upload(ctx)
	defer stream.CloseSend()

	// Allocate a buffer with `chunkSize` as the capacity
	// and length (making a 0 array of the size of `chunkSize`)
	writing := true
	buf := make([]byte, 250)
	for writing {
		// put as many bytes as `chunkSize` into the
		// buf array.
		n, err := file.Read(buf)

		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}

			return nil, nil
		}

		err = stream.Send(&pb.UploadParam{
			Content:       buf[:n],
			Filename:      "grpc.png",
			FileExtension: "png",
			FileSize:      1000,
		})

		if err != nil {
			return nil, nil
		}
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return nil, nil
	}
	return nil, nil
}
