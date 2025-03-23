package k8sclient

import (
	"context"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

type contextKey string

const k8sClientKey contextKey = "k8sClient"

func K8sClientMiddleware(k8sClient *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), k8sClientKey, k8sClient)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func GetK8sClient(ctx context.Context) *kubernetes.Clientset {
	client, ok := ctx.Value(k8sClientKey).(*kubernetes.Clientset)
	if !ok {
		return nil
	}
	return client
}
