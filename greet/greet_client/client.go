package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
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
	GreetEverybodyByBidiStreaming(client)
}

func GreetEverybodyByBidiStreaming(client greetpb.GreetServiceClient) {

	// prepare requests
	names := prepareNames()
	reqs := make([]greetpb.GreetEverybodyRequest, len(names))
	for i, name := range names {
		greet := greetpb.Greeting{FirstName: name.FirstName, LastName: name.LastName}
		reqs[i] = greetpb.GreetEverybodyRequest{Greeting: &greet}
	}

	stream, err := client.GreetEverybody(context.Background())
	if err != nil {
		log.Fatalf("Error when creating a stream to server: %v", err)
	}
	// send requests & receive responses concurrently
	waitChannel := make(chan string)

	go func() {
		for _, req := range reqs {
			fmt.Printf("Sending request: %v\n", &req)
			stream.Send(&req)
			time.Sleep(500*time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for  {
			res, err := stream.Recv()
			if err == io.EOF {
				// end of response --> signal to terminate request
				return
			}
			if err != nil {
				log.Fatalf("Error when receiving response %v", err)
				return
			}
			fmt.Println("Received: ", res.Result)
		}
		waitChannel <- "Done"
	}()

	<- waitChannel
	fmt.Println("Streaming Terminated")
}

func GreetByClientStreaming(client greetpb.GreetServiceClient) {
	names := prepareNames()
	reqs := make([]greetpb.ManyGreetRequest, len(names))
	for i, name := range names {
		greet := greetpb.Greeting{FirstName: name.FirstName, LastName: name.LastName}
		reqs[i] = greetpb.ManyGreetRequest{Greeting: &greet}
	}

	stream, err := client.ManyGreet(context.Background())
	if err != nil {
		log.Fatalf("Error when creating stream to server: %v", err)
	}

	for _, req := range reqs {
		fmt.Printf("Streaming request: %v\n", &req)
		stream.Send(&req)
		time.Sleep(500 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when closing and receiving from Server: %v", err)
	}
	fmt.Printf("ManyGreet response: %v", res.Result)
}

type Person struct {
	LastName  string
	FirstName string
}
func prepareNames() []Person {
	names := [] Person {
		{FirstName: "Bill", LastName: "Gates",},
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
	return names
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
