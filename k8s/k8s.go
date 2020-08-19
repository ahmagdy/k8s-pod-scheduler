package k8s

//go:generate mockgen -source=k8s.go -package=k8s -destination=k8s_mock.go

import (
	"os"
	"path/filepath"

	"github.com/ahmagdy/k8s-pod-scheduler/job"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// K8S service interface, that should abstract all k8s sdk details away
type K8S interface {
	GetPod(name string, namespace string) (*v1.Pod, error)
	GetCurrentNamespace() string
	CreateCronJob(job *job.SchedulerJob, namespace string) (string, error)
	CreatePod(job *job.SchedulerJob, namespace string) (string, error)
	CreateNamespace(name string) error
	DeletePod(name string, namespace string) error
	WatchPod(name string, namespace string) error
}

// k8SClient implementation of K8S interface
type k8SClient struct {
	log       *zap.Logger
	clientset kubernetes.Interface
}

var _ K8S = (*k8SClient)(nil)

// New instance of K8S concrete implementation
func New(logger *zap.Logger, clientset kubernetes.Interface) (K8S, error) {
	return &k8SClient{
		log:       logger,
		clientset: clientset,
	}, nil
}

// NewClientset create kubernetes clientset
func NewClientset(isInCluster bool) (kubernetes.Interface, error) {
	var config *rest.Config
	var err error
	if !isInCluster {
		var kubeconf string
		if home := homeDir(); home != "" {
			kubeconf = filepath.Join(home, ".kube", "config")
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconf)
		if err != nil {
			return nil, err
		}

	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
