package namespaces

import (
	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Namespaces) GetNamespaces(ctx *gin.Context) response.Either[error, []corev1.Namespace] {
	c.logger.Info("Listing Namespaces")

	allNamespaces, err := c.k8sClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return response.Left[error, []corev1.Namespace](err)
	}
	return response.Right[error, []corev1.Namespace](allNamespaces.Items)
}
