package main

import (
	"context"
	"log"
	"time"
	"io"

	pb "github.com/debianbaker/Go-Grpc/proto"
)

func callSayHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("Client stream started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil{
		log.Fatalf("could not send names: %v", err)
	}
	
	waitc := make(chan struct{})

	go func(){
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil{
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil{
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent the request with name: %v", name)
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}