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

func (*server) GreetEverybody(streamReq greetpb.GreetService_GreetEverybodyServer) error  {
	for {
		req, err := streamReq.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.Greeting.GetFirstName()
		lastName := req.Greeting.GetLastName()

		result := fmt.Sprintf("Hello %v, %v from Server", firstName, lastName)
		res := greetpb.GreetEverybodyResponse{Result: result}

		err = streamReq.Send(&res)
		if err != nil {
			log.Fatalf("Error while streaming response %v, error: %v", res, err)
			return err
		}
	}
}

var balance = 1000.00
func (*server) Withdraw(ctx context.Context, req *greetpb.WithdrawRequest) (*greetpb.WithdrawResponse, error)  {

	withdraw := req.Amount
	if withdraw <= 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Withdraw amount %v must be positive", withdraw))
	}

	if withdraw > balance {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Withdraw amount %v must not be higher than available amount %v", withdraw, balance))
	}

	balance -= withdraw
	res := greetpb.WithdrawResponse{Amount:balance}
	return &res, nil
}

func (*server) HelloWithDeadline(ctx context.Context, req *greetpb.HelloWithDeadlineRequest) (*greetpb.HelloWithDeadlineResponse, error)  {
	// simulate that computation work takes 3s
	for i := 0; i < 3; i++ {
		// we verify that request's deadline not passed yet.
		if ctx.Err() == context.Canceled {
			// the client has canceled request
			fmt.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled, "Client canceled the request")
		}
		// we simulate computation time
		time.Sleep(time.Second)
	}
	firstName := req.Greeting.FirstName
	lastName := req.Greeting.LastName
	result := "Hello " + lastName + ", " + firstName
	res := &greetpb.HelloWithDeadlineResponse{Result:result}

	return res, nil
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


