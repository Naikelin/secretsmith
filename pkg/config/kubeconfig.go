package config

import (
	"github.com/naikelin/secretsmith/pkg/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var log = logger.GetLogger().Sugar()

func InitKubeClient(masterURL, kubeconfig string) (*kubernetes.Clientset, error) {

	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Errorw("Error while building config from flag", "error", err)
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorw("Error while getting clientset from config", "error", err)
		return nil, err
	}
	return clientset, nil
}
