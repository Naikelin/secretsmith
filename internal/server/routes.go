package server

import (
	"github.com/gin-gonic/gin"
	"github.com/naikelin/secretsmith/internal/handlers/configmaps"
	"github.com/naikelin/secretsmith/internal/handlers/namespaces"
	"github.com/naikelin/secretsmith/internal/handlers/secrets"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/namespaces", namespaces.NewNamespaces().GetNamespacesHandler)

	r.POST("/configmaps", configmaps.NewConfigmaps().PostConfigmapsHandler)
	r.PUT("/configmaps/:cmns/:cmname", configmaps.NewConfigmaps().PutConfigmapsHandler)

	r.POST("/secrets", secrets.NewSecrets().PostSecretsHandler)
	r.PUT("/secrets/:secretns/:secretname", secrets.NewSecrets().PutSecretsHandler)
}
