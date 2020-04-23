package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context,req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	greeting := fmt.Sprintf("Hello %s, %s", lastName, firstName)
	res := greetpb.GreetResponse{Result: greeting}
	return &res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, streamRes greetpb.GreetService_GreetManyTimesServer) error  {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	for i := 0; i < 20 ; i++ {
		greeting := fmt.Sprintf("Hello %s, %s %d times", lastName, firstName, i)
		res := greetpb.GreetManyTimesResponse{Result: greeting}
		streamRes.Send(&res)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (*server) ManyGreet(streamReq greetpb.GreetService_ManyGreetServer) error  {

	greet := "Say hello to"
	for {
		req, err := streamReq.Recv()
		if err == io.EOF {
			// end of receiving requests
			fmt.Printf("End of receiving requests and send response!")
			res := greetpb.ManyGreetResponse{Result:greet}
			return streamReq.SendAndClose(&res)
		}
		if err != nil {
			log.Fatalf("Error when receiving client stream: %v", err)
		}
		greet += " " + req.GetGreeting().GetFirstName() +
			" " + req.GetGreeting().GetLastName() + "\n"
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen to the port 50051 due to: ", err)
	}
	fmt.Println("Server started successfully")

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Println("Greeting service registered to gRPC server successfully")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}


