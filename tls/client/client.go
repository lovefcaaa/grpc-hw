package main

import (
	"fmt"
	"log"

	pb "github.com/drgarcia1986/grpc-hw/tls/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	address := "localhost:50051"
	creds, err := credentials.NewClientTLSFromFile("../cert.crt", "")
	if err != nil {
		log.Fatal("error on create creds: ", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("failed to connect: ", err)
	}
	defer conn.Close()

	cli := pb.NewHelloWorldClient(conn)
	r, err := cli.Say(context.Background(), &pb.Request{Name: "World"})
	if err != nil {
		log.Fatal("error on request to server: ", err)
	}

	fmt.Println(r.Msg)
}
