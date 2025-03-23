package secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	v1 "k8s.io/api/core/v1"
)

func (h *Secrets) PostSecretsHandler(c *gin.Context) {
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())
	var secret v1.Secret
	if err := c.ShouldBindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := CreateSecret(c, k8sClient, secret)
	c.JSON(http.StatusOK, res)
}
