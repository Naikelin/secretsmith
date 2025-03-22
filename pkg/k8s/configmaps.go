package k8s

import (
	"context"

	"github.com/naikelin/secretsmith/pkg/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var log = logger.GetLogger().Sugar()

func ListConfigMaps(ctx context.Context, kubeclient *kubernetes.Clientset) ([]corev1.ConfigMap, error) {
	configMapList, err := kubeclient.CoreV1().ConfigMaps("").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Errorw("Error while listing all the configmaps", "error", err)
		return nil, err
	}
	return configMapList.Items, nil
}

func GetConfigMap(ctx context.Context, kubeclient *kubernetes.Clientset, cmns, cmName string) (*corev1.ConfigMap, error) {
	cm, err := kubeclient.CoreV1().ConfigMaps(cmns).Get(ctx, cmName, metav1.GetOptions{})
	if err != nil {
		log.Errorw("Error while getting a configmap", "namespace", cmns, "name", cmName, "error", err)
		return nil, err
	}
	return cm, nil
}

func UpdateConfigMap(ctx context.Context, kubeclient *kubernetes.Clientset, cmns, cmName string, configmap corev1.ConfigMap) (*corev1.ConfigMap, error) {
	cm, err := kubeclient.CoreV1().ConfigMaps(cmns).Update(ctx, &configmap, metav1.UpdateOptions{})
	if err != nil {
		log.Errorw("Error while updating the configmap", "namespace", cmns, "name", cmName, "error", err)
		return nil, err
	}
	return cm, nil
}

func GetNamespaces(ctx context.Context, kubeclient *kubernetes.Clientset) ([]corev1.Namespace, error) {
	allNamespaces, err := kubeclient.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Errorw("Error while getting all the namespaces", "error", err)
		return nil, err
	}
	return allNamespaces.Items, nil
}

func GetConfigMapsOfNS(ctx context.Context, kubeclient *kubernetes.Clientset, namespace string) ([]corev1.ConfigMap, error) {
	configMaps, err := kubeclient.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Errorw("Error while listing all configmaps of a namespace", "namespace", namespace, "error", err)
		return nil, err
	}
	return configMaps.Items, nil
}

func DeleteConfigMap(ctx context.Context, kubeclient *kubernetes.Clientset, namespace, name string) error {
	err := kubeclient.CoreV1().ConfigMaps(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		log.Errorw("Error while deleting the configmap", "namespace", namespace, "name", name, "error", err)
		return err
	}
	return nil
}

func CreateConfigMap(ctx context.Context, kubeclient *kubernetes.Clientset, cm corev1.ConfigMap) error {
	_, err := kubeclient.CoreV1().ConfigMaps(cm.Namespace).Create(ctx, &cm, metav1.CreateOptions{})
	if err != nil {
		log.Errorw("Error while creating the configmap", "namespace", cm.Namespace, "name", cm.Name, "error", err)
		return err
	}
	return nil
}
