package scheduler

// Start to start the cron scheduler
func (c *CronScheduler) Start() {
	c.cron.Start()
	defer c.cron.Stop()

	select {}
}
