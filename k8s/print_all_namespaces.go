package k8s

import (
	"fmt"
	"os"
	"text/tabwriter"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k8s *k8SClient) PrintAllNamespaces() error {
	pods, err := k8s.clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Namespace\tName")

	for _, pod := range pods.Items {
		fmt.Fprintln(writer, fmt.Sprintf("%s\t%s", pod.Namespace, pod.Name))
	}
	writer.Flush()
	return nil
}
