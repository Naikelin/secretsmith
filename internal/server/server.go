package server

import (
	"github.com/gin-gonic/gin"
	k8sClientMiddleware "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	loggerMiddleware "github.com/naikelin/secretsmith/internal/middlewares/logger"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

func Start(logger *zap.Logger, k8sClient *kubernetes.Clientset) {
	r := gin.Default()

	r.Use(loggerMiddleware.LoggerMiddleware(logger))
	r.Use(k8sClientMiddleware.K8sClientMiddleware(k8sClient))

	RegisterRoutes(r)

	port := ":8000"
	err := r.Run(port)
	if err != nil {
		logger.Fatal("Error starting server", zap.Error(err))
	}
	logger.Info("Server started", zap.String("port", port))
}
