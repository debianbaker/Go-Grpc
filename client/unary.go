package main

import (
	"time"
	"context"
	"log"

	pb "github.com/debianbaker/Go-Grpc/proto"
)

func callSayHello(client pb.GreetServiceClient){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil{
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%v", res.Message)
}