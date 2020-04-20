package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
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

	client := greetpb.NewGreetServiceClient(conn)

	// greetByUnary(client)
	// GreetByServerStreaming(client)
	DownloadFileStreaming(client)
}

func DownloadFileStreaming(client greetpb.GreetServiceClient) {
	req := greetpb.FileRequest{FileName:"test.txt"}

	resStream, err := client.DownloadFile(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when downloading file: %v", err)
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
