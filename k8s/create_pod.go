package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/securitycontext"
)

func (k8s *k8SClient) CreatePod(name string, namespace string) error {
	pod, err := k8s.clientset.CoreV1().Pods(namespace).Create(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				"name": name,
				"type": "production",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:                   name,
					Image:                  "node",
					TerminationMessagePath: v1.TerminationMessagePathDefault,
					ImagePullPolicy:        v1.PullIfNotPresent,
					SecurityContext:        securitycontext.ValidSecurityContextWithContainerDefaults(),
					Command:                []string{},
					Args:                   []string{},
				},
			},
			RestartPolicy: v1.RestartPolicyOnFailure,
			DNSPolicy:     v1.DNSDefault,
		}})
	if err != nil {
		return err
	}
	k8s.log.Info(pod.Name)
	return nil

}
