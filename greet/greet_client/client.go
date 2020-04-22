package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("Starting a client...")
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())// remove in prod code
	if err != nil {
		log.Fatal("Could not connect to server at port 50051 due to: ", err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	// greetByUnary(client)
	// GreetByServerStreaming(client)
	// GreetByClientStreaming(client)
	UploadFileByClientStreaming(client)
}

func UploadFileByClientStreaming(client greetpb.GreetServiceClient) {
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
		resPart := greetpb.Chunk{Content: buff[:bytesRead]}
		err = streamReq.Send(&resPart)
		if err != nil {
			log.Fatalf("Error when sending chunk: %v", err)
		}
	}

}

func GreetByClientStreaming(client greetpb.GreetServiceClient) {
	names := [] struct {
		LastName string
		FirstName string
	} {
		{FirstName: "Bill",LastName: "Gates",},
		{LastName: "Job", FirstName: "Steve",},
		{LastName: "Page", FirstName: "Larry",},
		{LastName: "Ritchie", FirstName: "Dennis",},
		{LastName: "Zuckerberg", FirstName: "Mark",},
		{LastName: "Ken", FirstName: "Thompson",},
		{LastName: "Torvalds", FirstName: "Linus",},
		{LastName: "Nakamoto", FirstName: "Satoshi",},
		{LastName: "Lovelace", FirstName: "Ada",},
		{LastName: "Berners-Lee", FirstName: "Tim",},
		{LastName: "Turing", FirstName: "Alan",},
	}
	reqs := make([]greetpb.ManyGreetRequest, len(names))
	for i, name := range names {
		greet := greetpb.Greeting{FirstName: name.FirstName, LastName:name.LastName}
		reqs[i] = greetpb.ManyGreetRequest{Greeting: &greet}
	}

	stream, err := client.ManyGreet(context.Background())
	if err != nil {
		log.Fatalf("Error when creating stream to server: %v", err)
	}

	for _, req := range reqs {
		fmt.Printf("Streaming request: %v\n", &req)
		stream.Send(&req)
		time.Sleep(500*time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when closing and receiving from Server: %v", err)
	}
	fmt.Printf("ManyGreet response: %v", res.Result)
}

func greetByUnary(client greetpb.GreetServiceClient) {
	greeting := greetpb.Greeting{
		FirstName: "Truc",
		LastName:  "Nguyen",
	}
	req := greetpb.GreetRequest{Greeting: &greeting}
	res, err := client.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func GreetByServerStreaming(client greetpb.GreetServiceClient) {
	greeting := greetpb.Greeting{
		FirstName: "Truc",
		LastName:  "Nguyen",
	}
	req := greetpb.GreetManyTimesRequest{Greeting: &greeting}

	resStream, err := client.GreetManyTimes(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling Server Streaming RPC: %v", err)
	}

	for {
		message, err := resStream.Recv()
		if err == io.EOF {
			// end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error when reading from response stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", message.Result)
	}
}
