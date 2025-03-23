package namespaces

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNamespaces(ctx context.Context, kubeclient *kubernetes.Clientset) ([]corev1.Namespace, error) {
	allNamespaces, err := kubeclient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return allNamespaces.Items, nil
}
