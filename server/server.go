package main

import (
	"context"
	"fmt"
	"log"
	"net"

	job "github.com/Ahmad-Magdy/k8s-pod-scheduler/job"

	"google.golang.org/grpc"
)

const (
	_port = ":8080"
)

type server struct{}

func (s *server) Add(ctx context.Context, req *job.AddJobRequest) (*job.AddJobResponse, error) {
	// TODO: Validate input fields
	j := job.SchedulerJobFromJob(req.Job)
	fmt.Println(j.Name)

	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	job.RegisterJobServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
