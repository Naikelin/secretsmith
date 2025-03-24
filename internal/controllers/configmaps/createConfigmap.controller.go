package configmaps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Configmaps) CreateConfigMap(ctx *gin.Context, cm corev1.ConfigMap) response.Either[error, string, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Creating ConfigMap", zap.String("request id", requestId), zap.String("namespace", cm.Namespace), zap.String("name", cm.Name))

	_, err := c.k8sClient.CoreV1().ConfigMaps(cm.Namespace).Create(ctx, &cm, metav1.CreateOptions{})
	if err != nil {
		c.logger.Error("Error creating ConfigMap", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, string, response.HttpMeta](
			response.HttpMeta{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error creating ConfigMap"},
			err,
		)
	}

	c.logger.Info("ConfigMap created successfully", zap.String("request id", requestId), zap.String("namespace", cm.Namespace), zap.String("name", cm.Name))
	return response.Right[error, string, response.HttpMeta](
		response.HttpMeta{StatusCode: http.StatusCreated, Message: "ConfigMap created successfully"}, "ConfigMap created successfully",
	)
}
