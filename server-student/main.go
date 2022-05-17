package main

import (
	"GRPC/database"
	"GRPC/server"
	"GRPC/studenttpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatal(err)
	}
	repository, err := database.NewPosgresRepository("postgres://postgres:Prueba123*@localhost:5432/grpc?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	newServer := server.NewServer(repository)
	s := grpc.NewServer()
	studenttpb.RegisterStudentServiceServer(s, newServer)

	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
