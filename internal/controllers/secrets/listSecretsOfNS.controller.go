package secrets

import (
	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Secrets) ListSecretsOfNS(ctx *gin.Context, namespace string) response.Either[error, []string, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Getting Secrets of namespace", zap.String("request id", requestId), zap.String("namespace", namespace))

	secrets, err := c.k8sClient.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		c.logger.Error("Error retrieving secrets", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, []string, response.HttpMeta](response.HttpMeta{
			StatusCode: 500,
			Message:    "Error retrieving secrets"},
			err)
	}

	secretsList := make([]string, 0)
	for _, secret := range secrets.Items {
		secretsList = append(secretsList, secret.Name)
	}

	c.logger.Info("Secrets retrieved successfully", zap.String("request id", requestId), zap.String("namespace", namespace))
	return response.Right[error, []string, response.HttpMeta](response.HttpMeta{
		StatusCode: 200,
		Message:    "Secrets retrieved successfully"},
		secretsList)
}
