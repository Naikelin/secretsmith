package main

import (
	"flag"

	"github.com/naikelin/secretsmith/internal/server"
	"github.com/naikelin/secretsmith/internal/utils/logger"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to your kubeconfig file.")
	flag.StringVar(&masterURL, "masterurl", "", "URL of your kube-apiserver.")
}

func main() {
	logger := logger.InitLogger()
	defer logger.Sync()

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		logger.Fatal("Error building kubeconfig", zap.Error(err))
	}
	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatal("Error building k8s client", zap.Error(err))
	}

	server.Start(logger, k8sClient)
}
