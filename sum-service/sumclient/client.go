package main

import (
	"../sumpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	defer cc.Close()

	r := &sumpb.Request{
		Num1: 10,
		Num2: 20,
	}

	serv := sumpb.NewSumServiceClient(cc)
	resp, err := serv.Sum(context.Background(), r)
	if err != nil {
		log.Fatal("Could Fetch data from rpc server: ", err)
	}

	fmt.Println("Resp: ", resp)
}
