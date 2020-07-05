package k8s

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_returns_namespace_value_given_valid_namespace(t *testing.T) {
	namespace := "my-special-namespace"

	k8sClient := newTestSimpleK8s(t)

	os.Setenv("POD_NAMESPACE", namespace)
	defer os.Unsetenv("POD_NAMESPACE")

	namespaceResult := k8sClient.GetCurrentNamespace()

	require.Equal(t, namespace, namespaceResult)
}

func Test_returns_default_given_unset_env_and_secret(t *testing.T) {
	namespace := "default"

	k8sClient := newTestSimpleK8s(t)

	namespaceResult := k8sClient.GetCurrentNamespace()

	require.Equal(t, namespace, namespaceResult)
}

func Test_returns_k8s_secret_when_env_var_not_set(t *testing.T) {
	namespace := "mynamespace"

	k8sClient := newTestSimpleK8s(t)
	readFile = func(filePath string) ([]byte, error) {
		return []byte(namespace), nil
	}
	// Because this monkey patching will affect other tests.
	// readFile is shared with the other tests in this directory
	defer func() { readFile = ioutil.ReadFile }()

	namespaceResult := k8sClient.GetCurrentNamespace()

	require.Equal(t, namespace, namespaceResult)
}
