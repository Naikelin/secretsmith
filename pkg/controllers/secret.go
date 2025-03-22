package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naikelin/secretsmith/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateSecret(c *gin.Context, kubeclient *kubernetes.Clientset) {
	ctx := context.TODO()
	var secret v1.Secret
	if err := c.ShouldBindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := k8s.CreateSecret(ctx, kubeclient, secret)
	c.JSON(http.StatusOK, res)
}

func UpdateSecret(c *gin.Context, kubeclient *kubernetes.Clientset) {
	ctx := context.TODO()
	secretName := c.Param("secretname")
	secretNS := c.Param("secretns")

	var secret v1.Secret
	if err := c.ShouldBindJSON(&secret); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := k8s.UpdateSecret(ctx, kubeclient, secretNS, secretName, secret)
	c.JSON(http.StatusOK, res)
}
