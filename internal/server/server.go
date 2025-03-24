package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
)

type Server struct {
	logger    *zap.Logger
	k8sClient *kubernetes.Clientset
	port      int
}

func NewServer(logger *zap.Logger, k8sClient *kubernetes.Clientset) *http.Server {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, _ := strconv.Atoi(portStr)

	NewServer := &Server{
		port:      port,
		logger:    logger,
		k8sClient: k8sClient,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("Server created!", zap.Int("port", NewServer.port))

	return server
}
