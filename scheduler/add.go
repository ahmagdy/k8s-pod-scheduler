package scheduler

import (
	"time"

	"go.uber.org/zap"

	"github.com/Ahmad-Magdy/k8s-pod-scheduler/job"
)

func (c *CronScheduler) Add(job *job.Job) error {
	id, err := c.cron.AddFunc(job.Cron, func() {
		// 15:04:05
		startTime := time.Now()
		c.log.Info("The job started", zap.String("job_name", job.Name), zap.Time("start_time", startTime),
			zap.String("image_to_execute", job.Specs.Image), zap.String("container_args", job.Specs.Args),
		)

		//time.Sleep(10 * time.Second)
		c.log.Info("The job has ended", zap.String("job_name", job.Name),
			zap.Int64("execution_time", time.Since(startTime).Milliseconds()))
	})
	if err != nil {
		return err
	}
	c.registeredJobs[job.Name] = int(id)
	return nil
}
