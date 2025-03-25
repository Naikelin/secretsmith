package secrets

import (
	"github.com/gin-gonic/gin"
	sController "github.com/naikelin/secretsmith/internal/controllers/secrets"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	logger "github.com/naikelin/secretsmith/internal/middlewares/logger"
)

func (h *Secrets) DeleteSecretHandler(c *gin.Context) {
	logger := logger.GetLogger(c.Request.Context())
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())

	sNS := c.Param("secretns")
	sN := c.Param("secretname")

	response := sController.NewSecrets(logger, k8sClient).DeleteSecret(c, sNS, sN)

	statusCode := response.GetMeta().StatusCode

	if response.IsLeft() {
		c.JSON(statusCode, gin.H{"error": response.GetLeft()})
		return
	}

	c.JSON(statusCode, response.GetRight())
}
