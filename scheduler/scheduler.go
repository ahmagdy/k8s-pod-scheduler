//go:generate mockgen -source=scheduler.go -package=scheduler -destination=scheduler_mock.go
package scheduler

import (
	"github.com/ahmagdy/k8s-pod-scheduler/job"
	"github.com/ahmagdy/k8s-pod-scheduler/k8s"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// Scheduler an interface for the scheduler operations
type Scheduler interface {
	Add(job *job.SchedulerJob) (jobID string, err error)
	Exists(jobName string) bool
}

// CronScheduler Implementation for Scheduler interface contains the needed properties
type CronScheduler struct {
	log            *zap.Logger
	cron           *cron.Cron
	k8s            k8s.K8S
	registeredJobs map[string]int
}

var _ Scheduler = (*CronScheduler)(nil)

// New instance of Scheduler
func New(logger *zap.Logger, k8s k8s.K8S) Scheduler {
	cron := cron.New(cron.WithSeconds())
	n := k8s.GetCurrentNamespace()
	logger.Info(n)
	s := &CronScheduler{
		log:            logger,
		cron:           cron,
		k8s:            k8s,
		registeredJobs: make(map[string]int),
	}
	cron.Start()
	return s
}
