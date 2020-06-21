package k8s

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"k8s.io/client-go/kubernetes/fake"
)

func TestCreateNamespace(t *testing.T) {
	tests := []struct {
		name          string
		namespace     string
		expectedError error
	}{
		{
			name:      "namespace is created",
			namespace: "abcd",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			k8sClient := newTestSimpleK8s(t)
			err := k8sClient.CreateNamespace(tc.namespace)
			require.Equal(t, tc.expectedError, err)
		})
	}
}

func newTestSimpleK8s(t *testing.T) K8S {
	client := k8SClient{}
	client.log = zaptest.NewLogger(t)
	client.clientset = fake.NewSimpleClientset()

	return &client
}
