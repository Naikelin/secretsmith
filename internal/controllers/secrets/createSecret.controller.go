package secrets

import (
	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Secrets) CreateSecret(ctx *gin.Context, secret corev1.Secret) response.Either[error, string, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Creating Secret", zap.String("request id", requestId), zap.String("namespace", secret.Namespace), zap.String("name", secret.Name))

	_, err := c.k8sClient.CoreV1().Secrets(secret.Namespace).Create(ctx, &secret, metav1.CreateOptions{})
	if err != nil {
		c.logger.Error("Error creating secret", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, string, response.HttpMeta](response.HttpMeta{
			StatusCode: 500,
			Message:    "Error creating secret"},
			err)
	}

	c.logger.Info("Secret created successfully", zap.String("request id", requestId), zap.String("namespace", secret.Namespace), zap.String("name", secret.Name))
	return response.Right[error, string, response.HttpMeta](response.HttpMeta{
		StatusCode: 201,
		Message:    "Secret created successfully"},
		"Secret created successfully")
}
