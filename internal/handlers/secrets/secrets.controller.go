package secrets

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetSecretsOfNS(ctx context.Context, kubeclient *kubernetes.Clientset, namespace string) ([]corev1.Secret, error) {
	secrets, err := kubeclient.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return secrets.Items, nil
}

func GetSecretData(ctx context.Context, kubeclient *kubernetes.Clientset, secreteNS, secretName string) corev1.Secret {
	secret, err := kubeclient.CoreV1().Secrets(secreteNS).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
	}

	return *secret
}

func UpdateSecret(ctx context.Context, kubeclient *kubernetes.Clientset, secretNS, secretName string, secret corev1.Secret) *corev1.Secret {
	sec, err := kubeclient.CoreV1().Secrets(secretNS).Update(ctx, &secret, metav1.UpdateOptions{})
	if err != nil {
	}

	return sec
}

func DeleteSecret(ctx context.Context, kubeclient *kubernetes.Clientset, secretNS, secretName string) error {
	err := kubeclient.CoreV1().Secrets(secretNS).Delete(ctx, secretName, metav1.DeleteOptions{})
	if err != nil {
	}

	return err
}

func CreateSecret(ctx context.Context, kubeclient *kubernetes.Clientset, secret corev1.Secret) error {
	_, err := kubeclient.CoreV1().Secrets(secret.Namespace).Create(ctx, &secret, metav1.CreateOptions{})
	if err != nil {
	}

	return err
}
