package main

import (
	"../sumpb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{}

func (*Server) Sum(ctx context.Context, req *sumpb.Request) (*sumpb.Response, error) {

	res := &sumpb.Response{
		Sum: req.GetNum1() + req.GetNum2(),
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	sumpb.RegisterSumServiceServer(grpcServer, &Server{})

	log.Println("Registered grpc server on port 9090")
	grpcServer.Serve(lis)
}
