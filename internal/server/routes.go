package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naikelin/secretsmith/internal/handlers/configmaps"
	"github.com/naikelin/secretsmith/internal/handlers/namespaces"
	"github.com/naikelin/secretsmith/internal/handlers/secrets"
	K8sMiddleware "github.com/naikelin/secretsmith/internal/middlewares/k8s"
	LoggerMiddleware "github.com/naikelin/secretsmith/internal/middlewares/logger"
	UUIDMiddleware "github.com/naikelin/secretsmith/internal/middlewares/uuid"
)

func (s *Server) RegisterRoutes() http.Handler {
	s.logger.Info("Registering Routes...")

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(UUIDMiddleware.RequestIDMiddleware())
	r.Use(LoggerMiddleware.LoggerMiddleware(s.logger))
	r.Use(K8sMiddleware.K8sClientMiddleware(s.k8sClient))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/namespaces", namespaces.NewNamespaces().GetNamespacesHandler)

	r.GET("/configmaps/:cmns", configmaps.NewConfigmaps().GetConfigmapsHandler)
	r.POST("/configmaps", configmaps.NewConfigmaps().PostConfigmapsHandler)
	r.PUT("/configmaps/:cmns/:cmname", configmaps.NewConfigmaps().PutConfigmapsHandler)

	r.GET("/secrets/:secretns", secrets.NewSecrets().GetSecretsHandler)
	r.POST("/secrets", secrets.NewSecrets().PostSecretsHandler)
	r.PUT("/secrets/:secretns/:secretname", secrets.NewSecrets().PutSecretsHandler)

	return r
}
