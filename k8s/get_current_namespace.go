package k8s

func (k8s *k8SClient) GetCurrentNamespace() (string, error) {
	return getNamespace(), nil
}
