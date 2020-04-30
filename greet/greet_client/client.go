package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// GreetEverybodyByBidiStreaming(client)
	// WithdrawByUnary(client)
	HelloWithDeadlineByUnary(client)
}

func HelloWithDeadlineByUnary(client greetpb.GreetServiceClient) {
	greeting := greetpb.Greeting{
		FirstName: "Truc",
		LastName:  "Nguyen",
	}
	req := greetpb.HelloWithDeadlineRequest{Greeting: &greeting}

	CallHelloWithTime(client, req, 1*time.Second)
	CallHelloWithTime(client, req, 5*time.Second)
}

func CallHelloWithTime(client greetpb.GreetServiceClient, req greetpb.HelloWithDeadlineRequest, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() //
	res, err := client.HelloWithDeadline(ctx, &req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout!")
			} else {
				fmt.Println("Unexpected error: ", statusErr)
			}
		} else {
			log.Fatalf("Error when calling Greet RPC: %v", err)
		}
		return
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func WithdrawByUnary(client greetpb.GreetServiceClient) {
	withdraw(client, 0)
	for i := 1; i <= 3; i++ {
		withdraw(client, float64(500))
	}
}

func withdraw(client greetpb.GreetServiceClient, amount float64) {
	fmt.Println("Try to withdraw an amount of: ", amount)
	res, err := client.Withdraw(context.Background(), &greetpb.WithdrawRequest{Amount:amount})

	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			// this is a user error (error defined by user at server side)
			// use can use switch statement to handle all types of errors you wish
			if resErr.Code() == codes.InvalidArgument {
				fmt.Printf("%v is not a valid amount\n", amount)
				fmt.Printf("Error: %v\n", resErr.Message())
			}
		} else {
			// framework error due to something unexpected
			log.Fatalf("Encounter a serious problem: %v\n", err)
		}
		return
	}

	fmt.Printf("Transaction sucessuful! Available amount is: %v\n", res.Amount)
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
