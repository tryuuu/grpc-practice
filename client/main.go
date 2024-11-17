package main

import (
	"context"
	"fmt"
	"grpc-practice/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect server:%v", err)
	}
	defer conn.Close()

	// FileServiceClientを生成
	client := pb.NewFileServiceClient(conn)
	callListFiles(client)
}

func callListFiles(client pb.FileServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.GetFilenames())
}
