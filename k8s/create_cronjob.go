package k8s

import (
	"github.com/ahmagdy/k8s-pod-scheduler/job"
	"go.uber.org/zap"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/securitycontext"
)

func (k8s *k8SClient) CreateCronJob(job *job.SchedulerJob, namespace string) (string, error) {
	if namespace == "" {
		namespace = k8s.GetCurrentNamespace()
	}
	k8s.log.Info("Creating cronjob",
		zap.String("cronjob_name", job.Name),
		zap.String("namespace", namespace),
	)
	objectMeta := metav1.ObjectMeta{
		GenerateName: job.Name,
		Labels: map[string]string{
			"name": job.Name,
			"type": "production",
		},
	}
	ttlTime := int32(0)
	jobSpec := batchv1.JobSpec{
		TTLSecondsAfterFinished: &ttlTime,
		Template: v1.PodTemplateSpec{
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
			},
		},
	}

	cronJob, err := k8s.clientset.BatchV1beta1().CronJobs(namespace).Create(&v1beta1.CronJob{
		ObjectMeta: objectMeta,
		Spec: v1beta1.CronJobSpec{
			Schedule:          job.Cron,
			ConcurrencyPolicy: v1beta1.ForbidConcurrent,
			JobTemplate: v1beta1.JobTemplateSpec{
				Spec: jobSpec,
			},
		},
	})
	if err != nil {
		return "", err
	}
	k8s.log.Info(cronJob.GetName())

	return cronJob.GetName(), nil

}
