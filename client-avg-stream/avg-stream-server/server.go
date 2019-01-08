package main

import (
	"../avgpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct{}

func (*Server) Avg(stream avgpb.AvgService_AvgServer) error {
	result := int32(0)
	n := int32(0)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&avgpb.AvgResponse{Num: result / n})
		}
		if err != nil {
			log.Fatal("Error while reading client stream: ", err)
		}
		//main computation here
		result += msg.GetNum()
		n++
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	avgpb.RegisterAvgServiceServer(grpcServer, &Server{})

	log.Println("Registered grpc server on port 9090")
	grpcServer.Serve(lis)
}
