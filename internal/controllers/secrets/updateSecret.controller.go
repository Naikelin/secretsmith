package secrets

import (
	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Secrets) UpdateSecret(ctx *gin.Context, secretNS, secretName string, secret corev1.Secret) response.Either[error, *corev1.Secret, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Updating Secret", zap.String("request id", requestId), zap.String("namespace", secretNS), zap.String("name", secretName))

	sec, err := c.k8sClient.CoreV1().Secrets(secretNS).Update(ctx, &secret, metav1.UpdateOptions{})
	if err != nil {
		c.logger.Error("Error updating secret", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, *corev1.Secret, response.HttpMeta](response.HttpMeta{
			StatusCode: 500,
			Message:    "Error updating secret"},
			err)
	}

	c.logger.Info("Secret updated successfully", zap.String("request id", requestId), zap.String("namespace", secretNS), zap.String("name", secretName))
	return response.Right[error, *corev1.Secret, response.HttpMeta](response.HttpMeta{
		StatusCode: 200,
		Message:    "Secret updated successfully"},
		sec)
}
