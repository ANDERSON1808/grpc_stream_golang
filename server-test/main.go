package main

import (
	"GRPC/database"
	"GRPC/server"
	"GRPC/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":5070")
	if err != nil {
		log.Fatal(err)
	}
	repository, err := database.NewPosgresRepository("postgres://postgres:Prueba123*@localhost:5432/grpc?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	testServer := server.NewTestServer(repository)
	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, testServer)

	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
