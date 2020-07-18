package scheduler

import (
	"time"

	"go.uber.org/zap"

	"github.com/ahmagdy/k8s-pod-scheduler/job"
)

// Add to register a job in the scheduler
func (c *CronScheduler) Add(job *job.SchedulerJob) (jobID string, err error) {
	log := c.log.With(zap.String("job_name", job.Name))

	log.Info("Adding job",
		zap.String("cron_expression", job.Cron))

	id, err := c.cron.AddFunc(job.Cron, func() {
		// 15:04:05
		startTime := time.Now()

		log.Info("The job started",
			zap.Time("start_time", startTime),
			zap.String("image_to_execute", job.Image),
			zap.Strings("container_args", job.Args),
		)

		name, err := c.k8s.CreatePod(job, "")
		if err != nil {
			c.log.Error("cron Add", zap.Error(err))
		}

		log.Info("The job has ended",
			zap.Duration("execution_time", time.Since(startTime).Round(time.Millisecond)),
		)

		c.k8s.WatchPod(name, "")
	})
	if err != nil {
		return "", err
	}
	c.registeredJobs[job.Name] = int(id)
	return job.Name, nil
}
