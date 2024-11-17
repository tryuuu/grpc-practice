package main

import (
	"context"
	"fmt"
	"grpc-practice/pb"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

// 構造体
type server struct {
	pb.UnimplementedFileServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles was invoked")

	dir := "/Users/ryutsuruyoshi/grpc-practice/storage"

	paths, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name()) // スライスにファイル名を追加
		}
	}

	res := &pb.ListFilesResponse{
		Filenames: filenames,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()
	// serverをサービスに登録(ListFilesとgRPC サーバーを結びつける)
	// FileService_serviceDescに基づいてListFilesがgRPC サーバーに登録される
	pb.RegisterFileServiceServer(s, &server{})

	fmt.Println("server is running")

	// サーバーを起動
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve:%v", err)
	}
}
