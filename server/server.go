package server

import (
	jobidl "github.com/ahmagdy/k8s-pod-scheduler/job/idl"
	"github.com/ahmagdy/k8s-pod-scheduler/k8s"
	"go.uber.org/zap"

	sc "github.com/ahmagdy/k8s-pod-scheduler/scheduler"

	"google.golang.org/grpc"
)

// K8SgRPC an implementation of the grpc server
type K8SgRPC struct {
	log       *zap.Logger
	scheduler sc.Scheduler
	k8s       k8s.K8S
}

// New instance of the GRPC server
func New(logger *zap.Logger, scheduler sc.Scheduler, k8s k8s.K8S) *grpc.Server {
	server := newGRPCServer(logger, scheduler, k8s)
	grpcServer := grpc.NewServer()
	jobidl.RegisterJobServiceServer(grpcServer, server)
	return grpcServer
}

func newGRPCServer(logger *zap.Logger, scheduler sc.Scheduler, k8s k8s.K8S) *K8SgRPC {
	return &K8SgRPC{
		log:       logger,
		scheduler: scheduler,
		k8s:       k8s,
	}
}
