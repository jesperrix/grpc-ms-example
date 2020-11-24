// Lightweight package to demonstrate grpc server implementation
// Implements the mymodelapi service which can be found here protos/my_model_api.proto
// The go output files are located at go/grpc/mymodelapi/
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/jesperrix/grpc-ms-example/go/grpc/mymodelapi"
	"google.golang.org/grpc"
	//"google.golang.org/protobuf/proto"
)

// MyModelApiServer implementation from go/grpc/mymodelapi/my_model_api_grpc.pb.go
type MyModelApiServer struct {
	pb.UnimplementedMyModelApiServer
	c int32
}

func (s *MyModelApiServer) Predict(ctx context.Context, input *pb.MyModelInput) (*pb.MyModelOutput, error) {
	s.c++
	log.Printf("Recieved msg #%d", s.c)
	return &pb.MyModelOutput{ModelId: "model_xyz", Y: []int32{1, 2, 3, s.c}}, nil
}

func newServer() *MyModelApiServer {
	s := &MyModelApiServer{}
	s.c = 0
	return s
}

func main() {
	lisAddress := fmt.Sprintf("localhost:%d", 56789)

	log.Println("Starting listener at: %s", lisAddress)
	lis, err := net.Listen("tcp", lisAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	// TODO: implement tls certificate

	log.Println("Registering gRPC server to listener")
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMyModelApiServer(grpcServer, newServer())

	grpcServer.Serve(lis)
}
