package scheduler

// Exists to verify if the job already registered
func (c *CronScheduler) Exists(jobName string) bool {
	_, ok := c.registeredJobs[jobName]
	return ok
}
