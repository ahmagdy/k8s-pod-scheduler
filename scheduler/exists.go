package scheduler

func (c *CronScheduler) Exists(jobName string) bool {
	_, ok := c.registeredJobs[jobName]
	return ok
}
