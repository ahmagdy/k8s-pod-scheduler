package scheduler

import (
	"time"

	"go.uber.org/zap"

	"github.com/ahmagdy/k8s-pod-scheduler/job"
)

// Add to register a job in the scheduler
func (c *CronScheduler) Add(job *job.SchedulerJob) (jobID string, err error) {
	c.log.Info("Adding job",
		zap.String("job_name", job.Name),
		zap.String("cron_expression", job.Cron))

	id, err := c.cron.AddFunc(job.Cron, func() {
		// 15:04:05
		startTime := time.Now()
		c.log.Info("The job started",
			zap.String("job_name", job.Cron),
			zap.Time("start_time", startTime),
			zap.String("image_to_execute", job.Image),
			zap.String("container_args", job.Args),
		)
		err := c.k8s.CreatePod(job.Name, "")
		if err != nil {
			c.log.Error("cron Add", zap.Error(err))
		}
		c.log.Info("The job has ended",
			zap.String("job_name", job.Name),
			zap.Duration("execution_time", time.Since(startTime).Round(time.Millisecond)),
		)
		c.k8s.WatchPod(job.Name, "")
	})
	if err != nil {
		return "", err
	}
	c.registeredJobs[job.Name] = int(id)
	return job.Name, nil
}
