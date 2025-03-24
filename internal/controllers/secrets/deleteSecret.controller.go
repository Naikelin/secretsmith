package secrets

import (
	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Secrets) DeleteSecret(ctx *gin.Context, secretNS, secretName string) response.Either[error, string, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Deleting Secret", zap.String("request id", requestId), zap.String("namespace", secretNS), zap.String("name", secretName))

	err := c.k8sClient.CoreV1().Secrets(secretNS).Delete(ctx, secretName, metav1.DeleteOptions{})
	if err != nil {
		c.logger.Error("Error deleting Secret", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, string, response.HttpMeta](response.HttpMeta{
			StatusCode: 500,
			Message:    "Error deleting secret"},
			err)
	}

	c.logger.Info("Secret deleted successfully", zap.String("request id", requestId), zap.String("namespace", secretNS), zap.String("name", secretName))
	return response.Right[error, string, response.HttpMeta](response.HttpMeta{
		StatusCode: 200,
		Message:    "Secret deleted successfully"},
		"Secret deleted successfully")
}
