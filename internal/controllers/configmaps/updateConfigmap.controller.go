package configmaps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Configmaps) UpdateConfigMap(ctx *gin.Context, cmns, cmName string, configmap corev1.ConfigMap) response.Either[error, *corev1.ConfigMap, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Updating ConfigMap", zap.String("request id", requestId), zap.String("namespace", cmns), zap.String("name", cmName))

	cm, err := c.k8sClient.CoreV1().ConfigMaps(cmns).Update(ctx, &configmap, metav1.UpdateOptions{})
	if err != nil {
		c.logger.Error("Error updating ConfigMap", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, *corev1.ConfigMap, response.HttpMeta](
			response.HttpMeta{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error updating ConfigMap",
			}, err,
		)
	}

	c.logger.Info("ConfigMap updated successfully", zap.String("request id", requestId), zap.String("namespace", cmns), zap.String("name", cmName))
	return response.Right[error, *corev1.ConfigMap, response.HttpMeta](
		response.HttpMeta{
			StatusCode: http.StatusOK,
			Message:    "ConfigMap updated successfully"},
		cm,
	)
}
