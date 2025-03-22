package main

import (
	"flag"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naikelin/secretsmith/pkg/logger"
	"github.com/naikelin/secretsmith/pkg/routes"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL  string
	kubeconfig string
	kubeclient *kubernetes.Clientset
)

func ZapLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		logger.Info("Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
		)

		c.Next()

		duration := time.Since(start)
		logger.Info("Response",
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
		)
	}
}

func main() {
	log := logger.GetLogger().Sugar()
	defer log.Sync()

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Errorw("Error while building config from flag", "error", err)
	}
	kubeclient, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorw("Error while getting clientset from config", "error", err)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(ZapLogger(log))
	routes.RegisterRoutes(router, kubeclient)

	hostPort := ":8000"
	log.Infow("Server running on", "hostPort", hostPort)
	err = router.Run(hostPort)
	if err != nil {
		log.Fatalw("Error while running server", "error", err)
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to your kubeconfig file.")
	flag.StringVar(&masterURL, "masterurl", "", "URL of your kube-apiserver.")

	logger.InitLogger()
}
