package configmaps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Configmaps) ListConfigMaps(ctx *gin.Context) response.Either[error, []string, response.HttpMeta] {
	requestId := ctx.GetString("RequestID")
	c.logger.Info("Listing ConfigMaps", zap.String("request id", requestId))

	configMapList, err := c.k8sClient.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})
	if err != nil {
		c.logger.Error("Error listing ConfigMaps", zap.String("request id", requestId), zap.Error(err))
		return response.Left[error, []string, response.HttpMeta](
			response.HttpMeta{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error retrieving ConfigMaps"},
			err,
		)
	}

	var configMaps []string
	for _, cm := range configMapList.Items {
		configMaps = append(configMaps, cm.Name)
	}

	if len(configMaps) == 0 {
		c.logger.Error("No ConfigMaps found", zap.String("request id", requestId))
		return response.Right[error, []string, response.HttpMeta](
			response.HttpMeta{
				StatusCode: http.StatusNotFound,
				Message:    "No ConfigMaps found"},
			[]string{},
		)
	}

	c.logger.Info("ConfigMaps retrieved successfully", zap.String("request id", requestId))
	return response.Right[error, []string, response.HttpMeta](
		response.HttpMeta{
			StatusCode: http.StatusOK,
			Message:    "ConfigMaps retrieved successfully"},
		configMaps,
	)
}
