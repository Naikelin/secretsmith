package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naikelin/secretsmith/pkg/k8s"
	"k8s.io/client-go/kubernetes"
)

func GetNamespaces(c *gin.Context, kubeclient *kubernetes.Clientset) {
	ctx := context.TODO()
	res, err := k8s.GetNamespaces(ctx, kubeclient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
