package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/upload/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Starting a client...")
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())// remove in prod code
	if err != nil {
		log.Fatal("Could not connect to server at port 50051 due to: ", err)
	}
	defer conn.Close()

	client := proto.NewUploadServiceClient(conn)
	UploadFileByClientStreaming(client)
}

func UploadFileByClientStreaming(client proto.UploadServiceClient) {
	chunkSize := 64 * 1024 // 64kiB recommended chunk size
	// get file to upload
	absPath, _ := filepath.Abs("../go-grpc-tutorial/greet/greet_server")
	file, err := os.Open(absPath + "/test.txt") // only test.txt accepted here.
	if err != nil {
		log.Fatalf("File not found or err: %v", err)
		return
	}
	defer func() {
		path, _ := os.Getwd()
		fmt.Println("File saved at: ", path)
		file.Close()
	}()
	streamReq, err := client.Upload(context.Background())
	if err != nil {
		log.Fatalf("Error when creating a stream to server: %v", err)
		return
	}
	buff := make([]byte, chunkSize)
	for  {
		bytesRead, err := file.Read(buff)
		if err == io.EOF {
			streamReq.CloseAndRecv()
			return
		}
		if err != nil {
			log.Fatalf("Error when reading file: %v", err)
			return
		}
		resPart := proto.Chunk{Content: buff[:bytesRead]}
		err = streamReq.Send(&resPart)
		if err != nil {
			log.Fatalf("Error when sending chunk: %v", err)
		}
	}

}
