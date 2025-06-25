package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/debianbaker/Go-Grpc/proto"
	"log"
)
const ( 
	port = ":8080"
)

func main(){
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	
	names := &pb.NamesList{
		Names: []string{"Puneet", "Alice", "Bob"},
	}

	callSayHello(client)
	callSayHelloServerStream(client, names)
}