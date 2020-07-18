package k8s

import (
	"github.com/ahmagdy/k8s-pod-scheduler/job"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/securitycontext"
)

func (k8s *k8SClient) CreatePod(job *job.SchedulerJob, namespace string) (string, error) {
	if namespace == "" {
		namespace = k8s.GetCurrentNamespace()
	}
	k8s.log.Info("Creating pod",
		zap.String("pod_name", job.Name),
		zap.String("namespace", namespace),
	)
	pod, err := k8s.clientset.CoreV1().Pods(namespace).Create(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: job.Name,
			Labels: map[string]string{
				"name": job.Name,
				"type": "production",
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:                   job.Name,
					Image:                  job.Image,
					TerminationMessagePath: v1.TerminationMessagePathDefault,
					ImagePullPolicy:        v1.PullIfNotPresent,
					SecurityContext:        securitycontext.ValidSecurityContextWithContainerDefaults(),
					Command:                job.Commands,
					Args:                   job.Args,
				},
			},
			RestartPolicy: v1.RestartPolicyOnFailure,
			DNSPolicy:     v1.DNSDefault,
		}})
	if err != nil {
		return "", err
	}
	k8s.log.Info(pod.GetName())

	return pod.GetName(), nil

}
