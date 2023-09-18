package main

import (
	"context"
	"grpc/invoicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("Hello World"),
		Docx: []byte("Hello World"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Failed to listen on port 8080: %v", err)
	}

	serverRegister := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegister, service)
	err = serverRegister.Serve(lis)

	if err != nil {
		log.Fatal("Failed to serve gRPC server over port 8080: %v", err)
	}
}
