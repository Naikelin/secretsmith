package configmaps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	v1 "k8s.io/api/core/v1"
)

func (h *Configmaps) PostConfigmapsHandler(c *gin.Context) {
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())
	var configMap v1.ConfigMap
	if err := c.ShouldBindJSON(&configMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := CreateConfigMap(c, k8sClient, configMap)
	c.JSON(http.StatusOK, res)
}
