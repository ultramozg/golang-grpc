package main

import (
	"../primepb"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	defer cc.Close()

	r := &primepb.PrimeRequest{
		Num: 10575,
	}

	serv := primepb.NewPrimeServiceClient(cc)
	stream, err := serv.Prime(context.Background(), r)
	if err != nil {
		log.Fatal("Couldn't get data: ", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while reading stream: ", err)
		}
		log.Println("Response: ", msg)
	}
}
