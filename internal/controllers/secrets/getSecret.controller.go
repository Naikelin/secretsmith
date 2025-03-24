package secrets

import (
	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Secrets) GetSecretData(ctx *gin.Context, secretNS, secretName string) response.Either[error, *corev1.Secret, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Getting Secret", zap.String("request id", requestId), zap.String("namespace", secretNS), zap.String("name", secretName))

	secret, err := c.k8sClient.CoreV1().Secrets(secretNS).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		c.logger.Error("Error retrieving secret", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, *corev1.Secret, response.HttpMeta](response.HttpMeta{
			StatusCode: 500,
			Message:    "Error retrieving secret"},
			err)
	}

	c.logger.Info("Secret retrieved successfully", zap.String("request id", requestId), zap.String("namespace", secretNS), zap.String("name", secretName))
	return response.Right[error, *corev1.Secret, response.HttpMeta](response.HttpMeta{
		StatusCode: 200,
		Message:    "Secret retrieved successfully"},
		secret)
}
