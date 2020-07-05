package k8s

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/fake"
	testcore "k8s.io/client-go/testing"
)

func TestMYTest(t *testing.T) {
	k8sClient, watcher := newWatcherClient(t)
	podName := "mybod"
	namespace := "mynamespace"

	pod, _ := k8sClient.GetPod(podName, namespace)

	pod.Status.Phase = v1.PodPending
	watcher.Add(pod)

	pod.Status.Phase = v1.PodRunning
	watcher.Modify(pod)

	pod.Status.Phase = v1.PodSucceeded
	watcher.Modify(pod)

	k8sClient.(*k8SClient).watchStatus(watcher, podName, namespace)

	_, err := k8sClient.GetPod(podName, namespace)
	require.Error(t, err)
	require.Equal(t, fmt.Sprintf(`pods "%s" not found`, podName), err.Error())
}

func newWatcherClient(t *testing.T) (K8S, *watch.RaceFreeFakeWatcher) {
	fakeClientset := fake.NewSimpleClientset(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:         "mybod",
			GenerateName: "mybod",
			Namespace:    "mynamespace",
			Labels: map[string]string{
				"tag": "",
			},
		},
	})

	watcher := watch.NewRaceFreeFake()

	fakeClientset.PrependWatchReactor("pods", testcore.DefaultWatchReactor(watcher, nil))
	k8sClient := newTestSimpleK8s(t).(*k8SClient)
	k8sClient.clientset = fakeClientset

	return k8sClient, watcher
}
