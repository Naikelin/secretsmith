package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naikelin/secretsmith/pkg/controllers"
	"k8s.io/client-go/kubernetes"
)

func RegisterRoutes(router *gin.Engine, kubeclient *kubernetes.Clientset) {
	// Namespaces routes
	router.GET("/namespaces", func(c *gin.Context) {
		controllers.GetNamespaces(c, kubeclient)
	})

	// ConfigMap routes
	router.POST("/configs", func(c *gin.Context) {
		controllers.CreateConfigMap(c, kubeclient)
	})
	router.PUT("/configs/:cmns/:cmname", func(c *gin.Context) {
		controllers.UpdateConfigMap(c, kubeclient)
	})

	// Secret routes
	router.POST("/secrets", func(c *gin.Context) {
		controllers.CreateSecret(c, kubeclient)
	})
	router.PUT("/secrets/:secretns/:secretname", func(c *gin.Context) {
		controllers.UpdateSecret(c, kubeclient)
	})
}
