package namespaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	nsController "github.com/naikelin/secretsmith/internal/controllers/namespaces"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	"github.com/naikelin/secretsmith/internal/middlewares/logger"
)

func (h *Namespaces) GetNamespacesHandler(c *gin.Context) {
	logger := logger.GetLogger(c.Request.Context())
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())

	response := nsController.NewConfigmapsController(logger, k8sClient).GetNamespaces(c)

	if response.IsLeft() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.GetLeft()})
		return
	}

	c.JSON(http.StatusOK, response.GetRight())
}
