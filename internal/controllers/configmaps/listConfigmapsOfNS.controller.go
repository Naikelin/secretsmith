package configmaps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Configmaps) ListConfigMapsOfNS(ctx *gin.Context, namespace string) response.Either[error, []string, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Getting ConfigMaps of namespace", zap.String("request id", requestId), zap.String("namespace", namespace))

	configMaps, err := c.k8sClient.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		c.logger.Error("Error retrieving ConfigMaps", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, []string, response.HttpMeta](
			response.HttpMeta{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error retrieving ConfigMaps"},
			err,
		)
	}

	configMapsList := make([]string, 0)
	for _, configMap := range configMaps.Items {
		configMapsList = append(configMapsList, configMap.Name)
	}

	c.logger.Info("ConfigMaps retrieved successfully", zap.String("request id", requestId), zap.String("namespace", namespace))
	return response.Right[error, []string, response.HttpMeta](
		response.HttpMeta{
			StatusCode: http.StatusOK,
			Message:    "ConfigMaps retrieved successfully"},
		configMapsList,
	)
}
