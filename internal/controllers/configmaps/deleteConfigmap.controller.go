package configmaps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Configmaps) DeleteConfigMap(ctx *gin.Context, namespace, name string) response.Either[error, string, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Deleting ConfigMap", zap.String("request id", requestId), zap.String("namespace", namespace), zap.String("name", name))

	err := c.k8sClient.CoreV1().ConfigMaps(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		c.logger.Error("Error deleting ConfigMap", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, string, response.HttpMeta](
			response.HttpMeta{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error deleting ConfigMap"},
			err,
		)
	}
	c.logger.Info("ConfigMap deleted successfully", zap.String("request id", requestId), zap.String("namespace", namespace), zap.String("name", name))
	return response.Right[error, string, response.HttpMeta](
		response.HttpMeta{StatusCode: http.StatusOK, Message: "ConfigMap deleted successfully"}, "ConfigMap deleted successfully",
	)
}
