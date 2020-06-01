package scheduler

func (c *CronScheduler) Start() {
	c.cron.Start()
	defer c.cron.Stop()

	select {}
}
