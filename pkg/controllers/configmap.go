package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naikelin/secretsmith/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateConfigMap(c *gin.Context, kubeclient *kubernetes.Clientset) {
	ctx := context.TODO()
	var configMap v1.ConfigMap
	if err := c.ShouldBindJSON(&configMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := k8s.CreateConfigMap(ctx, kubeclient, configMap)
	c.JSON(http.StatusOK, res)
}

func UpdateConfigMap(c *gin.Context, kubeclient *kubernetes.Clientset) {
	ctx := context.TODO()
	cmName := c.Param("cmname")
	cmNS := c.Param("cmns")

	var configMap v1.ConfigMap
	if err := c.ShouldBindJSON(&configMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := k8s.UpdateConfigMap(ctx, kubeclient, cmNS, cmName, configMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
