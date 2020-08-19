package k8s

import (
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ahmagdy/k8s-pod-scheduler/job"

	"github.com/stretchr/testify/require"
)

func TestCreateCronJob(t *testing.T) {
	type input struct {
		namespace string
		job       *job.SchedulerJob
	}
	tests := []struct {
		name              string
		input             input
		expectedError     error
		expectedNamespace string
	}{
		{
			name: "cronjob is created with provided namespace",
			input: input{
				namespace: "abcd",
				job: &job.SchedulerJob{
					Name: "my-job",
				},
			},
			expectedNamespace: "abcd",
		},
		{
			name: "cron is created with default namespace if not provided",
			input: input{
				job: &job.SchedulerJob{
					Name: "my-job",
				},
			},
			expectedNamespace: "default",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			k8sClient := newTestSimpleK8s(t)
			_, err := k8sClient.CreateCronJob(tc.input.job, tc.input.namespace)

			require.Equal(t, tc.expectedError, err)
			crons, err := k8sClient.(*k8SClient).clientset.BatchV1beta1().CronJobs("").List(v1.ListOptions{})
			require.Equal(t, tc.expectedNamespace, crons.Items[0].Namespace)

			// It should be the returned name from k8sClient.CreatePod but k8s fake client implementation doesn't evaluate the pod
			// require.Contains(t, name, tc.input.job.Name)
			require.Contains(t, crons.Items[0].GetGenerateName(), tc.input.job.Name)
		})
	}
}
