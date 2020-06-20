package main

import (
	"log"
	"net"

	"github.com/Ahmad-Magdy/k8s-pod-scheduler/server"
)

const (
	_port = ":8080"
)

func main() {
	server := server.New(nil, nil)
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
