package job

import (
	"testing"

	jobidl "github.com/ahmagdy/k8s-pod-scheduler/job/idl"

	"github.com/stretchr/testify/require"

	"github.com/golang/protobuf/ptypes/wrappers"
)

func TestSchedulerJobFromJob(t *testing.T) {
	tests := []struct {
		name     string
		job      *jobidl.Job
		expected *SchedulerJob
	}{
		{
			name: "job is mapped to scheduler job",
			job: &jobidl.Job{
				Name: &wrappers.StringValue{Value: "XYZ"},
				Cron: &wrappers.StringValue{Value: "* * * * * *"},
				Spec: &Spec{
					Image: &wrappers.StringValue{Value: "magdy.dev/xyz:version1"},
					Args:  []string{"--yz"},
				},
			},
			expected: &SchedulerJob{
				Name:  "XYZ",
				Cron:  "* * * * * *",
				Image: "magdy.dev/xyz:version1",
				Args:  []string{"--yz"},
			},
		},
		{
			name: "given job without specs, it should map it without specs properties",
			job: &jobidl.Job{
				Name: &wrappers.StringValue{Value: "XYZ"},
				Cron: &wrappers.StringValue{Value: "* * * * * *"},
			},
			expected: &SchedulerJob{
				Name: "XYZ",
				Cron: "* * * * * *",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := SchedulerJobFromJob(tc.job)
			require.Equal(t, tc.expected, res)
		})
	}
}
