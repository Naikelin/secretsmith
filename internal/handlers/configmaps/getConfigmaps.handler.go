package configmaps

import (
	"github.com/gin-gonic/gin"
	cmController "github.com/naikelin/secretsmith/internal/controllers/configmaps"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	logger "github.com/naikelin/secretsmith/internal/middlewares/logger"
)

func (h *Configmaps) GetConfigmapsHandler(c *gin.Context) {
	logger := logger.GetLogger(c.Request.Context())
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())

	cmNS := c.Param("cmns")

	response := cmController.NewConfigmapsController(logger, k8sClient).ListConfigMapsOfNS(c, cmNS)

	statusCode := response.GetMeta().StatusCode

	if response.IsLeft() {
		c.JSON(statusCode, gin.H{"error": response.GetLeft()})
		return
	}

	c.JSON(statusCode, response.GetRight())
}
