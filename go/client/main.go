package main

import (
	"context"
	"fmt"
	pb "github.com/jesperrix/grpc-ms-example/go/grpc/mymodelapi"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	server_addr := "localhost"
	server_port := 56789

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", server_addr, server_port), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewMyModelApiClient(conn)

	for i := 0; i <= 1000; i++ {
		//ctx, cancel = con
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		//mymodel_input := pb.MyModelInput{}
		//mymodel_input.X =
		model_out, err := client.Predict(ctx, &pb.MyModelInput{X: []int32{1, 2, 3}, Y: []int32{4, 5, 6}})
		if err != nil {
			log.Fatalf("%v.Predict(_) = _, %v: ", client, err)
		}
		if i%100 == 0 {
			log.Printf("sent/received %d messages.", i)
			log.Println(model_out)
		}

	}

	//client.Predict

}
