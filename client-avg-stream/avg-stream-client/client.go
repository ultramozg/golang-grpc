package main

import (
	"../avgpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	cc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	defer cc.Close()

	serv := avgpb.NewAvgServiceClient(cc)
	stream, err := serv.Avg(context.Background())

	if err != nil {
		log.Fatal("Couldn't get data: ", err)
	}

	request := []*avgpb.AvgRequest{
		&avgpb.AvgRequest{Num: 10},
		&avgpb.AvgRequest{Num: 20},
		&avgpb.AvgRequest{Num: 30},
		&avgpb.AvgRequest{Num: 40},
		&avgpb.AvgRequest{Num: 50},
		&avgpb.AvgRequest{Num: 60},
		&avgpb.AvgRequest{Num: 70},
		&avgpb.AvgRequest{Num: 80},
	}

	for _, v := range request {
		fmt.Println("Sendind data :", v)
		stream.Send(v)
		time.Sleep(1 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error while recieveng response from server", err)
	}

	fmt.Println("Avg is:", response)
}
