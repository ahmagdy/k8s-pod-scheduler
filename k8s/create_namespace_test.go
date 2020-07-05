package k8s

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"k8s.io/client-go/kubernetes/fake"
)

func TestCreateNamespace(t *testing.T) {
	k8sClient := newTestSimpleK8s(t)
	namespace := "abcd"

	err := k8sClient.CreateNamespace(namespace)
	require.Equal(t, nil, err)
}

func newTestSimpleK8s(t *testing.T) K8S {
	client := k8SClient{}
	client.log = zaptest.NewLogger(t)
	client.clientset = fake.NewSimpleClientset()

	return &client
}
