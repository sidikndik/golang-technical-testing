package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/gofiber/recipes/fiber-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var db *sql.DB


func main() {

	// initialize db connection
	var err error
	db, err = sql.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("GRPC db connection failed: %v", err)
	}

	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}


func (s *server) Login(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}
