package main

import (
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Starting a client...")
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())// remove in prod code
	if err != nil {
		log.Fatal("Could not connect to server at port 50051 due to: ", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Client connect to server successfully: %f", c)

}
