package configmaps

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

type Configmaps struct {
	k8sClient *kubernetes.Clientset
	logger    *zap.Logger
}

func NewConfigmapsController(logger *zap.Logger, k8sClient *kubernetes.Clientset) *Configmaps {
	return &Configmaps{k8sClient: k8sClient, logger: logger}
}
