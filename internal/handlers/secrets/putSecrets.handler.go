package secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	v1 "k8s.io/api/core/v1"
)

func (h *Secrets) PutSecretsHandler(c *gin.Context) {
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())
	secretName := c.Param("secretname")
	secretNS := c.Param("secretns")
	var secret v1.Secret
	if err := c.ShouldBindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := UpdateSecret(c, k8sClient, secretNS, secretName, secret)
	c.JSON(http.StatusOK, res)
}
