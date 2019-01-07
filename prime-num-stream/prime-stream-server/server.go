package main

import (
	"../primepb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct{}

func (*Server) Prime(req *primepb.PrimeRequest, stream primepb.PrimeService_PrimeServer) error {
	for n, k := req.GetNum(), int32(2); n > 1; {
		if n%k == 0 {
			//send into stream
			stream.Send(&primepb.PrimeResponse{Num: k})
			time.Sleep(1 * time.Second)
			n = n / k
		} else {
			k = k + 1
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
	primepb.RegisterPrimeServiceServer(grpcServer, &Server{})

	log.Println("Registered grpc server on port 9090")
	grpcServer.Serve(lis)
}
