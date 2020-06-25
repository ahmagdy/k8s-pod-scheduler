package server

import (
	"go.uber.org/zap"

	job "github.com/ahmagdy/k8s-pod-scheduler/job"
	sc "github.com/ahmagdy/k8s-pod-scheduler/scheduler"

	"google.golang.org/grpc"
)

// K8SgRPC an implementation of the grpc server
type K8SgRPC struct {
	log       *zap.Logger
	scheduler sc.Scheduler
}

// New instance of the GRPC server
func New(logger *zap.Logger, scheduler sc.Scheduler) *grpc.Server {
	server := &K8SgRPC{
		log:       logger,
		scheduler: scheduler,
	}
	grpcServer := grpc.NewServer()
	job.RegisterJobServiceServer(grpcServer, server)
	return grpcServer
}
