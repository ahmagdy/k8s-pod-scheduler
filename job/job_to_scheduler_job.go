package job

import jobidl "github.com/ahmagdy/k8s-pod-scheduler/job/idl"

// SchedulerJob a representation for a job that will be executed
type SchedulerJob struct {
	Name     string
	Cron     string
	Image    string
	Args     []string
	Commands []string
}

// SchedulerJobFromJob map Job input into a SchedulerJob representation
func SchedulerJobFromJob(job *jobidl.Job) *SchedulerJob {
	sj := &SchedulerJob{
		Name: job.Name,
		Cron: job.Cron,
	}
	if job.Spec != nil {
		sj.Image = job.Spec.GetImage()
		sj.Args = job.Spec.GetArgs()
		sj.Commands = job.Spec.GetCommands()
	}
	return sj
}
