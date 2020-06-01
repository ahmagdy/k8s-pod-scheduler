package k8s

import (
	"os"

	"k8s.io/client-go/tools/clientcmd"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

type K8S interface {
	GetPod(name string, namespace string) (*v1.Pod, error)
	GetCurrentNamespace() (string, error)
	CreatePod(name string, namespace string) error
	CreateNamespace(name string) error
	DeletePod(name string, namespace string) error
}

type k8SClient struct {
	log       *zap.Logger
	clientset *kubernetes.Clientset
}

var _ K8S = (*k8SClient)(nil)

func New(logger *zap.Logger) (K8S, error) {
	config, err := clientcmd.BuildConfigFromFlags("", homeDir())
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &k8SClient{
		log:       logger,
		clientset: clientset,
	}, nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
