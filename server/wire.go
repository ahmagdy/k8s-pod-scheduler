//+build wireinject

package server

import (
	"github.com/ahmagdy/k8s-pod-scheduler/k8s"
	"github.com/ahmagdy/k8s-pod-scheduler/logger"
	"github.com/ahmagdy/k8s-pod-scheduler/scheduler"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeServer() (*grpc.Server, error) {
	wire.Build(logger.New, k8s.NewClientset, k8s.New, scheduler.New, New)
	return &grpc.Server{}, nil
}
