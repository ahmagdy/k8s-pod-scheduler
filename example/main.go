package main

import (
	"flag"
	"net"

	"github.com/ahmagdy/k8s-pod-scheduler/logger"
	"go.uber.org/zap"

	"github.com/ahmagdy/k8s-pod-scheduler/server"
)

const (
	_port = ":8080"
)

var (
	isInCluster = flag.Bool("is-in-cluster", true, "Is this pod running inside the cluster")
)

func main() {
	flag.Parse()

	log := logger.New()

	server, err := server.InitializeServer(*isInCluster)
	if err != nil {
		log.Fatal("Error from init server", zap.Error(err))
	}
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		log.Fatal("failed to listen", zap.Error(err))
	}

	if err := server.Serve(lis); err != nil {
		log.Fatal("failed to serve", zap.Error(err))
	}
	log.Info("Serving on port", zap.String("PORT", _port))
}
