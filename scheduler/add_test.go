package scheduler

import (
	"testing"

	"github.com/ahmagdy/k8s-pod-scheduler/job"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name         string
		job          *job.SchedulerJob
		returnsError bool
	}{
		{
			name: "new job is being registered",
			job:  &job.SchedulerJob{Name: "first job", Cron: "* * * * * *"},
		},
		{
			name:         "returns error give invalid cron",
			job:          &job.SchedulerJob{Name: "another job", Cron: "* * * *"},
			returnsError: true,
		},
	}
	scheduler := newTestScheduler(t)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			jobID, err := scheduler.Add(tc.job)
			if tc.returnsError {
				require.Error(t, err)
			} else {
				require.NotEmpty(t, jobID)
				require.Equal(t, tc.job.Name, jobID)
			}
		})
	}
}
