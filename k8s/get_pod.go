package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k8s *k8SClient) GetPod(name string, namespace string) (*v1.Pod, error) {
	if namespace == "" {
		namespace = k8s.GetCurrentNamespace()
	}
	return k8s.clientset.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
}
