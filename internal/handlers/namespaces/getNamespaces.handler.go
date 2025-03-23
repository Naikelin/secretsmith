package namespaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	k8sClient "github.com/naikelin/secretsmith/internal/middlewares/k8s"
)

func (h *Namespaces) GetNamespacesHandler(c *gin.Context) {
	k8sClient := k8sClient.GetK8sClient(c.Request.Context())
	res, err := GetNamespaces(c, k8sClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
