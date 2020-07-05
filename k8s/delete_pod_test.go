package k8s

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes/fake"
)

func TestDeletePod(t *testing.T) {
	tests := []struct {
		name          string
		namespace     string
		podName       string
		expectedError error
	}{
		{
			name:      "pod is deleted providing a namespace",
			podName:   "my-bod",
			namespace: "ab-cd",
		},
		{
			// delete the pod from the default namespace if a specific namespace is not provided
			name:    "pod is deleted providing empty namespace",
			podName: "my-bod",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			k8sClient := getK8sClientForDeletePods(t)
			err := k8sClient.DeletePod(tc.podName, tc.namespace)
			require.Equal(t, tc.expectedError, err)
		})
	}
}

func getK8sClientForDeletePods(t *testing.T) K8S {
	k8sClient := newTestSimpleK8s(t)
	k8sClient.(*k8SClient).clientset = fake.NewSimpleClientset(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-bod",
			Namespace: "ab-cd",
			Labels: map[string]string{
				"tag": "",
			},
		},
	},
		&v1.Pod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-bod",
				Namespace: "default",
				Labels: map[string]string{
					"tag": "",
				},
			},
		})
	return k8sClient
}
