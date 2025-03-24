package configmaps

import (
	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Configmaps) ListConfigMaps(ctx *gin.Context) response.Either[error, []corev1.ConfigMap] {
	c.logger.Info("Listing ConfigMaps")

	configMapList, err := c.k8sClient.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})
	if err != nil {
		c.logger.Error("Error listing ConfigMaps", zap.Error(err))
		return response.Left[error, []corev1.ConfigMap](err)
	}
	c.logger.Info("ConfigMaps listed successfully")
	return response.Right[error, []corev1.ConfigMap](configMapList.Items)
}

func (c *Configmaps) GetConfigMap(ctx *gin.Context, cmns, cmName string) response.Either[error, *corev1.ConfigMap] {
	c.logger.Info("Getting ConfigMap", zap.String("namespace", cmns), zap.String("name", cmName))

	cm, err := c.k8sClient.CoreV1().ConfigMaps(cmns).Get(ctx, cmName, metav1.GetOptions{})
	if err != nil {
		return response.Left[error, *corev1.ConfigMap](err)
	}
	return response.Right[error, *corev1.ConfigMap](cm)
}

func (c *Configmaps) UpdateConfigMap(ctx *gin.Context, cmns, cmName string, configmap corev1.ConfigMap) response.Either[error, *corev1.ConfigMap] {
	c.logger.Info("Updating ConfigMap", zap.String("namespace", cmns), zap.String("name", cmName))

	cm, err := c.k8sClient.CoreV1().ConfigMaps(cmns).Update(ctx, &configmap, metav1.UpdateOptions{})
	if err != nil {
		return response.Left[error, *corev1.ConfigMap](err)
	}
	return response.Right[error, *corev1.ConfigMap](cm)

}

func (c *Configmaps) GetConfigMapsOfNS(ctx *gin.Context, namespace string) response.Either[error, []corev1.ConfigMap] {
	c.logger.Info("Getting ConfigMaps of namespace", zap.String("namespace", namespace))

	configMaps, err := c.k8sClient.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return response.Left[error, []corev1.ConfigMap](err)
	}
	return response.Right[error, []corev1.ConfigMap](configMaps.Items)
}

func (c *Configmaps) DeleteConfigMap(ctx *gin.Context, namespace, name string) response.Either[error, string] {
	c.logger.Info("Deleting ConfigMap", zap.String("namespace", namespace), zap.String("name", name))

	err := c.k8sClient.CoreV1().ConfigMaps(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return response.Left[error, string](err)
	}
	return response.Right[error, string]("ConfigMap deleted successfully")
}

func (c *Configmaps) CreateConfigMap(ctx *gin.Context, cm corev1.ConfigMap) response.Either[error, string] {
	c.logger.Info("Creating ConfigMap", zap.String("namespace", cm.Namespace), zap.String("name", cm.Name))

	_, err := c.k8sClient.CoreV1().ConfigMaps(cm.Namespace).Create(ctx, &cm, metav1.CreateOptions{})
	if err != nil {
		return response.Left[error, string](err)
	}
	return response.Right[error, string]("ConfigMap created successfully")
}
