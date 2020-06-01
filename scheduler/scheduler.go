package scheduler

import (
	"github.com/Ahmad-Magdy/k8s-pod-scheduler/job"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Scheduler interface {
	Start()
	Add(job *job.JobDescription) error
}

type CronScheduler struct {
	log            *zap.Logger
	cron           *cron.Cron
	registeredJobs map[string]int
}

var _ Scheduler = (*CronScheduler)(nil)

func New(logger *zap.Logger) Scheduler {
	cron := cron.New(cron.WithSeconds())

	return &CronScheduler{
		log:            logger,
		cron:           cron,
		registeredJobs: make(map[string]string),
	}
}
