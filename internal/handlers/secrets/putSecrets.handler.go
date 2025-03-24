package secrets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sController "github.com/naikelin/secretsmith/internal/controllers/secrets"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	"github.com/naikelin/secretsmith/internal/middlewares/logger"
	v1 "k8s.io/api/core/v1"
)

func (h *Secrets) PutSecretsHandler(c *gin.Context) {
	logger := logger.GetLogger(c.Request.Context())
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())
	secretName := c.Param("secretname")
	secretNS := c.Param("secretns")
	var secret v1.Secret
	if err := c.ShouldBindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := sController.NewSecrets(logger, k8sClient).UpdateSecret(c, secretNS, secretName, secret)

	if response.IsLeft() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.GetLeft()})
		return
	}

	c.JSON(http.StatusOK, response.GetRight())
}
