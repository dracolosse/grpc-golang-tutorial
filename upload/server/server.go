package main

import (
	"fmt"
	"go-grpc-tutorial/upload/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

type server struct {}

func (*server) Upload(streamReq proto.UploadService_UploadServer) error {
	// create a file to handle uploading file
	file, err := os.Create("./outpput")
	if err != nil {
		log.Fatalf("Error when creating a new file")
		streamReq.SendAndClose(
			&proto.UploadStatus{
				Result:"Failed: Cannot create file!",
				Code: proto.UploadStatusCode(1),
			})
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {panic(err)}
	}()


	for {
		res, err := streamReq.Recv()
		if err == io.EOF {
			streamReq.SendAndClose(
				&proto.UploadStatus{
					Result:"File uploaded successfully!",
					Code: proto.UploadStatusCode(0),
				})
			return nil
		}
		if err != nil {
			log.Fatalf("Error when receiving file %v", err)
			streamReq.SendAndClose(
				&proto.UploadStatus{
					Result:"Failed: Error when receiving file!",
					Code: proto.UploadStatusCode(1),
				})
		}
		_, err = file.Write(res.Content)
		if err != nil {
			log.Fatalf("Error when writing file %v", err)
			streamReq.SendAndClose(
				&proto.UploadStatus{
					Result:"Failed: Error when writing file!",
					Code: proto.UploadStatusCode(1),
				})
			return err
		}
	}
}


func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen to the port 50051 due to: ", err)
	}
	fmt.Println("Server started successfully")

	s := grpc.NewServer()
	proto.RegisterUploadServiceServer(s, &server{})

	fmt.Println("Greeting service registered to gRPC server successfully")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
