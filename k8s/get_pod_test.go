package k8s

import (
	"testing"

	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPod(t *testing.T) {
	k8sClient := getK8sClientForGetPod(t)
	namespace := "abcd"
	podName := "mypod"

	pod, err := k8sClient.GetPod(podName, namespace)

	require.NoError(t, err)
	require.NotNil(t, pod)
	require.Equal(t, podName, pod.GetName())
}

func getK8sClientForGetPod(t *testing.T) K8S {
	k8sClient := newTestSimpleK8s(t)
	k8sClient.(*k8SClient).clientset = fake.NewSimpleClientset(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mypod",
			Namespace: "abcd",
			Labels: map[string]string{
				"tag": "",
			},
		},
	})
	return k8sClient
}
