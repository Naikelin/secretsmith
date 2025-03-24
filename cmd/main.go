package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/naikelin/secretsmith/internal/server"
	"github.com/naikelin/secretsmith/internal/utils/logger"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func initK8sClient(kubeconfig, masterURL string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	if kubeconfig == "" {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("Error getting in-cluster config: %v", err)
		}
	} else {
		config, err = clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
		if err != nil {
			return nil, fmt.Errorf("Error building kubeconfig: %v", err)
		}
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("Error building Kubernetes client: %v", err)
	}

	return k8sClient, nil
}

func gracefulShutdown(apiServer *http.Server, done chan bool, logger *zap.Logger) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	logger.Info("Shutting down gracefully, press Ctrl+C again to force")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := apiServer.Shutdown(shutdownCtx); err != nil {
		logger.Warn("Server forced to shutdown with error, exiting", zap.Error(err))
	}

	logger.Info("Server exiting")

	done <- true
}

func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = "/home/user/.kube/config"
	}

	masterURL := os.Getenv("MASTER_URL")

	log := logger.InitLogger()
	defer log.Sync()

	log.Info("Starting server...")

	k8sClient, err := initK8sClient(kubeconfig, masterURL)
	if err != nil {
		log.Fatal("Failed to initialize Kubernetes client", zap.Error(err))
	}

	server := server.NewServer(log, k8sClient)

	done := make(chan bool, 1)

	go gracefulShutdown(server, done, log)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("HTTP server error", zap.Error(err))
	}

	<-done
	log.Info("Graceful shutdown complete.")
}
