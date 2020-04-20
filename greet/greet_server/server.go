package main

import (
	"context"
	"fmt"
	"go-grpc-tutorial/greet/greetpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
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

/**
For demonstration purpose, you should create an file named test.txt in the current directory.
The size file should be large to let you check at the client side.
 */
func (*server) DownloadFile(req *greetpb.FileRequest, streamRes greetpb.GreetService_DownloadFileServer) error  {
	chunkSize := 64 * 1024 // 64kiB recommended chunk size
	absPath, _ := filepath.Abs("../go-grpc-tutorial/greet/greet_server")
	file, err := os.Open(absPath + "/" + req.GetFileName()) // only test.txt accepted here.
	if err != nil {
		fmt.Printf("File not found or err: %v", err)
		return err
	}
	defer file.Close()

	buff := make([]byte, chunkSize)
	for {
		bytesRead, err := file.Read(buff)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		resp := greetpb.Chunk{
			Chunk: buff[:bytesRead],
		}

		if err = streamRes.Send(&resp); err != nil {
			log.Println("error while sending chunk:", err)
			return err
		}
	}
	return nil
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


