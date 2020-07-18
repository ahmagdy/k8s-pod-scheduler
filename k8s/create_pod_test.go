package k8s

import (
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ahmagdy/k8s-pod-scheduler/job"

	"github.com/stretchr/testify/require"
)

func TestCreatePod(t *testing.T) {
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
			name: "pod is created with provided namespace",
			input: input{
				namespace: "abcd",
				job: &job.SchedulerJob{
					Name: "my-job",
				},
			},
			expectedNamespace: "abcd",
		},
		{
			name: "pod is created with default namespace if not provided",
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
			_, err := k8sClient.CreatePod(tc.input.job, tc.input.namespace)

			require.Equal(t, tc.expectedError, err)
			pod, err := k8sClient.(*k8SClient).clientset.CoreV1().Pods("").List(v1.ListOptions{})
			require.Equal(t, tc.expectedNamespace, pod.Items[0].Namespace)

			// It should be the returned name from k8sClient.CreatePod but k8s fake client implementation doesn't evaluate the pod
			// require.Contains(t, name, tc.input.job.Name)
			require.Contains(t, pod.Items[0].GetGenerateName(), tc.input.job.Name)
		})
	}
}
