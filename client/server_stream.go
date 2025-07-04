package main

import (
	"context"
	"log"
	"io"

	pb "github.com/debianbaker/Go-Grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Streaming started")
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil{
		log.Fatalf("could not send names: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("streaming finished")
}
 