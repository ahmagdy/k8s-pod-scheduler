package k8s

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k8s *k8SClient) WatchPod(name string, namespace string) error {
	k8s.log.Info("Creating a watcher")
	watch, err := k8s.clientset.CoreV1().Pods(namespace).Watch(metav1.ListOptions{
		FieldSelector: fields.Set{"metadata.name": name}.String(),
		LabelSelector: labels.Everything().String(),
	})
	k8s.log.Info("WATCHER")
	if err != nil {
		return err
	}
	func() {
		for {
			k8s.log.Info("HERE")
			select {
			case event, ok := <-watch.ResultChan():
				if !ok {
					k8s.log.Error("NOT OK")
					return
				}
				k8s.log.Info("", zap.String("type", string(event.Type)))

				resp := event.Object.(*v1.Pod)
				k8s.log.Info("Pod status:", zap.String("status", string(resp.Status.Phase)))
				//status = resp.Status
				// one way of doing it ama
				if resp.Status.Phase != v1.PodPending {
					watch.Stop()
				}
				status := resp.Status.Phase
				if status != v1.PodPending && status != v1.PodRunning {
					if status == v1.PodFailed {
						k8s.log.Error("Pod failed")

					} else if status == v1.PodSucceeded {
						k8s.log.Info("Pod succeeded")
					}
					watch.Stop()
					err = k8s.DeletePod(name, namespace)
					if err != nil {
						k8s.log.Error("Couldn't delete the pod", zap.Error(err))
					}
				}
			case <-time.After(5 * time.Minute):
				fmt.Println("timeout to wait for pod active")
				watch.Stop()
			}
		}
	}()
	return nil

	for {
		newPod, err := k8s.GetPod(name, namespace)

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
