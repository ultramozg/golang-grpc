package main

import (
	"../maxpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	cc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	defer cc.Close()

	serv := maxpb.NewMaxServiceClient(cc)
	stream, err := serv.Max(context.Background())

	if err != nil {
		log.Fatal("Couldn't get data: ", err)
	}

	data := []int64{3, 8, 1, 12, 11, 24, 30}
	done := make(chan struct{})

	go func() {
		for _, v := range data {
			fmt.Println("Sendind data :", v)
			stream.Send(&maxpb.Request{Num: v})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("Error while receiving: ", err)
				break
			}
			fmt.Println("Receiving data: ", res)
		}
		close(done)
	}()

	<-done
}
