package k8s

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func (k8s *k8SClient) DeletePod(name string, namespace string) error {
	if namespace == "" {
		namespace = k8s.GetCurrentNamespace()
	}
	err := k8s.clientset.CoreV1().Pods(namespace).Delete(name, &metav1.DeleteOptions{})
	return err
}
