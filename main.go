package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/lomins/demo-grpc-tutorial/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Cannot create listener: %v", err)
	}

	fmt.Println("Server is listening...")

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatal("impossible to serve: %v", err)
	}
}

func (s *myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}
