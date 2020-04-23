package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/download/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Starting a client...")
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())// remove in prod code
	if err != nil {
		log.Fatal("Could not connect to server at port 50051 due to: ", err)
	}
	defer conn.Close()

	client := proto.NewDownloadServiceClient(conn)

	// greetByUnary(client)
	// GreetByServerStreaming(client)
	DownloadFileStreaming(client)
}


func DownloadFileStreaming(client proto.DownloadServiceClient) {
	req := proto.FileRequest{FileName:"test.txt"}

	resStream, err := client.DownloadFile(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when downloading file: %v", err)
		return
	}

	var blob []byte
	for {
		message, err := resStream.Recv()
		if err == io.EOF {
			log.Println("received all chunks")
			break
		}
		if err != nil {
			log.Fatalf("Error when reading file: %v", err)
			panic(err)
		}
		blob = append(blob, message.Chunk...)
	}

	fmt.Printf("End of downloading file")
}
