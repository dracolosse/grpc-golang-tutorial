protoc calculation/proto/calculation.proto --go_out=plugins=grpc:.
go run greet/greet_server/server.go
go run greet/greet_client/client.go
