package main

import (
	"../maxpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct{}

func (*Server) Max(stream maxpb.MaxService_MaxServer) error {
	max := int64(0)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal("Error while reading client stream: ", err)
			return err
		}
		if msg.GetNum() > max {
			max = msg.GetNum()
			err := stream.Send(&maxpb.Response{Num: max})
			if err != nil {
				log.Fatal("Error while sending response: ", err)
				return err
			}
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	maxpb.RegisterMaxServiceServer(grpcServer, &Server{})

	log.Println("Registered grpc server on port 9090")
	grpcServer.Serve(lis)
}
