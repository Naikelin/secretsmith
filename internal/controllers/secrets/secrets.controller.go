package secrets

import (
	"github.com/gin-gonic/gin"
	response "github.com/naikelin/secretsmith/internal/utils/responses"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Secrets) GetSecretsOfNS(ctx *gin.Context, namespace string) response.Either[error, []corev1.Secret] {
	c.logger.Info("Getting Secrets of namespace", zap.String("namespace", namespace))

	secrets, err := c.k8sClient.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return response.Left[error, []corev1.Secret](err)
	}
	return response.Right[error, []corev1.Secret](secrets.Items)
}

func (c *Secrets) GetSecretData(ctx *gin.Context, secreteNS, secretName string) response.Either[error, *corev1.Secret] {
	c.logger.Info("Getting Secret", zap.String("namespace", secreteNS), zap.String("name", secretName))

	secret, err := c.k8sClient.CoreV1().Secrets(secreteNS).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		return response.Left[error, *corev1.Secret](err)
	}
	return response.Right[error, *corev1.Secret](secret)
}

func (c *Secrets) UpdateSecret(ctx *gin.Context, secretNS, secretName string, secret corev1.Secret) response.Either[error, *corev1.Secret] {
	c.logger.Info("Updating Secret", zap.String("namespace", secretNS), zap.String("name", secretName))

	sec, err := c.k8sClient.CoreV1().Secrets(secretNS).Update(ctx, &secret, metav1.UpdateOptions{})
	if err != nil {
		return response.Left[error, *corev1.Secret](err)
	}
	return response.Right[error, *corev1.Secret](sec)
}

func (c *Secrets) DeleteSecret(ctx *gin.Context, secretNS, secretName string) response.Either[error, string] {
	c.logger.Info("Deleting Secret", zap.String("namespace", secretNS), zap.String("name", secretName))

	err := c.k8sClient.CoreV1().Secrets(secretNS).Delete(ctx, secretName, metav1.DeleteOptions{})
	if err != nil {
		return response.Left[error, string](err)
	}

	return response.Right[error, string]("Secret deleted successfully")
}

func (c *Secrets) CreateSecret(ctx *gin.Context, secret corev1.Secret) response.Either[error, string] {
	c.logger.Info("Creating Secret", zap.String("namespace", secret.Namespace), zap.String("name", secret.Name))

	_, err := c.k8sClient.CoreV1().Secrets(secret.Namespace).Create(ctx, &secret, metav1.CreateOptions{})
	if err != nil {
		return response.Left[error, string](err)
	}
	return response.Right[error, string]("Secret created successfully")
}
