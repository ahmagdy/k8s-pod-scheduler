package k8s

import (
	"time"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
)

func (k8s *k8SClient) WatchPod(podName string, namespace string) error {
	for {
		newPod, err := k8s.GetPod(podName, namespace)

		if err != nil {
			return err
		}
		status := newPod.Status.Phase

		k8s.log.Info("Checking pod status.", zap.String("status", string(status)),
			zap.String("message", newPod.Status.Message), zap.String("reason", newPod.Status.Reason))

		if status == v1.PodRunning {
			k8s.log.Info("Deleting the pod", zap.String("pod_name", newPod.Name))
			err := k8s.DeletePod("my-pod", namespace)
			if err != nil {
				return err
			}
			break
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}
