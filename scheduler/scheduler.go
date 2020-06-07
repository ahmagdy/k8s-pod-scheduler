package scheduler

import (
	"github.com/Ahmad-Magdy/k8s-pod-scheduler/job"
	"github.com/Ahmad-Magdy/k8s-pod-scheduler/k8s"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Scheduler interface {
	Start()
	Add(job *job.Job) error
	Exists(jobName string) bool
}

type CronScheduler struct {
	log            *zap.Logger
	cron           *cron.Cron
	k8s            k8s.K8S
	registeredJobs map[string]int
}

var _ Scheduler = (*CronScheduler)(nil)

func New(logger *zap.Logger, k8s k8s.K8S) Scheduler {
	cron := cron.New(cron.WithSeconds())
	n := k8s.GetCurrentNamespace()
	logger.Info(n)
	return &CronScheduler{
		log:            logger,
		cron:           cron,
		k8s:            k8s,
		registeredJobs: make(map[string]int),
	}
}
