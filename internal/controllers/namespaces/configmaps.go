package namespaces

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

type Namespaces struct {
	k8sClient *kubernetes.Clientset
	logger    *zap.Logger
}

func NewConfigmapsController(logger *zap.Logger, k8sClient *kubernetes.Clientset) *Namespaces {
	return &Namespaces{k8sClient: k8sClient, logger: logger}
}
