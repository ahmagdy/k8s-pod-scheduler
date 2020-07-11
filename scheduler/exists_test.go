package scheduler

import (
	"testing"

	"github.com/ahmagdy/k8s-pod-scheduler/job"

	"github.com/stretchr/testify/require"
)

func TestExists(t *testing.T) {
	tests := []struct {
		name    string
		jobName string
		exist   bool
	}{
		{
			name:    "giving an existing job returns true",
			jobName: "abcd",
			exist:   true,
		},
		{
			name:    "giving an non-existing job returns false",
			jobName: "gdef",
			exist:   false,
		},
	}
	scheduler := newTestScheduler(t)
	scheduler.Add(&job.SchedulerJob{Name: "abcd", Cron: "@daily"})
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			exists := scheduler.Exists(tc.jobName)
			require.Equal(t, tc.exist, exists)
		})
	}
}
