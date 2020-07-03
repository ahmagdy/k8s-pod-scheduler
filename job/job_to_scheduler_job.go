package job

// SchedulerJob a representation for a job that will be executed
type SchedulerJob struct {
	Name     string
	Cron     string
	Image    string
	Args     []string
	Commands []string
}

// SchedulerJobFromJob map Job input into a SchedulerJob representation
func SchedulerJobFromJob(job *Job) *SchedulerJob {
	sj := &SchedulerJob{
		Name: job.Name.GetValue(),
		Cron: job.Cron.GetValue(),
	}
	if job.Spec != nil {
		sj.Image = job.Spec.Image.GetValue()
		sj.Args = job.Spec.GetArgs()
		sj.Commands = job.Spec.GetCommands()
	}
	return sj
}
