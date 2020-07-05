package job

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/protobuf/ptypes/wrappers"
)

func TestSchedulerJobFromJob(t *testing.T) {
	tests := []struct {
		name     string
		job      *Job
		expected *SchedulerJob
	}{
		{
			name: "job is mapped to scheduler job",
			job: &Job{
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
			job: &Job{
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
