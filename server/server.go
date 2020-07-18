package server

import (
	jobidl "github.com/ahmagdy/k8s-pod-scheduler/job/idl"
	"go.uber.org/zap"

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
	server := newGRPCServer(logger, scheduler)
	grpcServer := grpc.NewServer()
	jobidl.RegisterJobServiceServer(grpcServer, server)
	return grpcServer
}

func newGRPCServer(logger *zap.Logger, scheduler sc.Scheduler) *K8SgRPC {
	return &K8SgRPC{
		log:       logger,
		scheduler: scheduler,
	}
}
