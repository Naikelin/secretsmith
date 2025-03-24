package configmaps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cmController "github.com/naikelin/secretsmith/internal/controllers/configmaps"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	logger "github.com/naikelin/secretsmith/internal/middlewares/logger"
	v1 "k8s.io/api/core/v1"
)

func (h *Configmaps) PutConfigmapsHandler(c *gin.Context) {
	logger := logger.GetLogger(c.Request.Context())
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())

	cmName := c.Param("cmname")
	cmNS := c.Param("cmns")
	var configMap v1.ConfigMap
	if err := c.ShouldBindJSON(&configMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response := cmController.NewConfigmapsController(logger, k8sClient).UpdateConfigMap(c, cmNS, cmName, configMap)

	if response.IsLeft() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": response.GetLeft()})
		return
	}

	c.JSON(http.StatusOK, response.GetRight())
}
