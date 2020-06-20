package server

import (
	"context"

	"github.com/Ahmad-Magdy/k8s-pod-scheduler/job"
)

func (s *K8SgRPC) Add(ctx context.Context, req *job.AddJobRequest) (*job.AddJobResponse, error) {
	// TODO: Validate input fields
	j := job.SchedulerJobFromJob(req.Job)
	id, err := s.scheduler.Add(j)
	if err != nil {
		return nil, err
	}
	return &job.AddJobResponse{Id: id}, nil
}
