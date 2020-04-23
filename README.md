#Golang gRPC Tutorial – Part 3 – gRPC Server Streaming

This is the third part of our Golang gRPC Tutorial series. In this post, we’re gonna learn how to build a gRPC Server Streaming and 
how to invoke it at the client side in Go.

https://nguyenttz.com/golang-grpc-tutorial-part-3/

Here is an exercise for you if you want to make sure that you well understand Server Streaming RPC. We would like you to implement 
Server Streaming RPC  which accepts a Request to download a large file.

A client RPC will send a request which contains a file name. If the server cannot find out that the file does not exist, it then returns an error. 
Otherwise, it will stream that file back to client.

Finally, you should write a client RPC in Go to invoke the API
