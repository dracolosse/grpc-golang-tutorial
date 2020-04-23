package main

import (
	"fmt"
	"go-grpc-tutorial/download/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

type server struct {}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen to the port 50051 due to: ", err)
	}
	fmt.Println("Server started successfully")

	s := grpc.NewServer()
	proto.RegisterDownloadServiceServer(s, &server{})

	fmt.Println("Greeting service registered to gRPC server successfully")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}


/**
For demonstration purpose, you should create an file named test.txt in the current directory.
The size file should be large to let you check at the client side.
*/
func (*server) DownloadFile(req *proto.FileRequest, streamRes proto.DownloadService_DownloadFileServer) error  {
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
		resp := proto.Chunk{
			Chunk: buff[:bytesRead],
		}

		if err = streamRes.Send(&resp); err != nil {
			log.Println("error while sending chunk:", err)
			return err
		}
	}
	return nil
}
