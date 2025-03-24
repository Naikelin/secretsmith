package secrets

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

type Secrets struct {
	k8sClient *kubernetes.Clientset
	logger    *zap.Logger
}

func NewSecrets(logger *zap.Logger, k8sClient *kubernetes.Clientset) *Secrets {
	return &Secrets{k8sClient: k8sClient, logger: logger}
}
