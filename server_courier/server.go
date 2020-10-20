package main

import (
	"log"
	"net"

	"github.com/digutier/finsd/courier"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := courier.Server{}

	grpcServer := grpc.NewServer()

	courier.RegisterCourierServiceServer(grpcServer, &s)
	//Instanciar variables de la estructura Server
	s.MapaSeguimiento = make(map[int64]string)
	s.OrdenesP = make([]*courier.OrdenPyme, 0)
	s.OrdenesR = make([]*courier.OrdenRetail, 0)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
