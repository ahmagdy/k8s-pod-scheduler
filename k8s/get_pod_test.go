package k8s

import (
	"testing"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPod(t *testing.T) {

	tests := []struct {
		name      string
		podName   string
		namespace string
	}{
		{
			name:      "returns mypod given valid pod and namespace",
			podName:   "mypod",
			namespace: "abcd",
		},
		{
			name:    "uses default namespace given pod name only",
			podName: "anotherpod",
		},
	}
	k8sClient := getK8sClientForGetPod(t)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pod, err := k8sClient.GetPod(tc.podName, tc.namespace)

			require.NoError(t, err)
			require.NotNil(t, pod)
			require.Contains(t, tc.podName, pod.GetGenerateName())
		})
	}
}

func getK8sClientForGetPod(t *testing.T) K8S {
	k8sClient := newTestSimpleK8s(t)
	k8sClient.(*k8SClient).clientset = fake.NewSimpleClientset(
		&v1.Pod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:         "mypod",
				GenerateName: "mypod",
				Namespace:    "abcd",
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
				Name:         "anotherpod",
				GenerateName: "anotherpod",
				Namespace:    "default",
				Labels: map[string]string{
					"tag": "",
				},
			},
		},
	)
	return k8sClient
}
