package cli

func (s *EnvSettings) SetNamespace(ns string) {
	s.namespace = ns
}

func (s *EnvSettings) SetKubeConfig(kubecfg string) {
	s.KubeConfig = kubecfg
}

func (s *EnvSettings) SetKubeToken(kubetoken string) {
	s.KubeToken = kubetoken
}

func (s *EnvSettings) SetRegistryConfig(registcfg string) {
	s.RegistryConfig = registcfg
}

func (s *EnvSettings) SetRepositoryConfig(repocfg string) {
	s.RepositoryConfig = repocfg
}

func (s *EnvSettings) SetDebug(b bool) {
	s.Debug = b
}
