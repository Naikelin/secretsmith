package namespaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Namespaces) ListNamespaces(ctx *gin.Context) response.Either[error, []string, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Listing Namespaces", zap.String("request id", requestId))

	allNamespaces, err := c.k8sClient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		c.logger.Error("Error listing Namespaces", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, []string, response.HttpMeta](response.HttpMeta{
			StatusCode: 500,
			Message:    "Error retrieving Namespaces",
		}, err)
	}

	nameSpaces := []string{}
	for _, ns := range allNamespaces.Items {
		if ns.Name[:4] == "kube" || ns.Name[:3] == "k8s" {
			continue
		}
		nameSpaces = append(nameSpaces, ns.Name)
	}

	if len(nameSpaces) == 0 {
		c.logger.Error("No Namespaces found", zap.String("request id", requestId))
		return response.Right[error, []string, response.HttpMeta](response.HttpMeta{
			StatusCode: http.StatusNotFound,
			Message:    "No Namespaces found",
		}, []string{})
	}

	c.logger.Info("Namespaces retrieved successfully", zap.String("request id", requestId))
	return response.Right[error, []string, response.HttpMeta](response.HttpMeta{
		StatusCode: http.StatusOK,
		Message:    "Namespaces retrieved successfully",
	}, nameSpaces)
}
