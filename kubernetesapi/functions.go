package kubernetesapi

import (
	"context"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
)

func Connect() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func LocalConnect() *kubernetes.Clientset {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("",kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func GetNameSpaces (clientset *kubernetes.Clientset, labelMatcher string) *apiv1.NamespaceList {
	namespaces, err  := clientset.CoreV1().Namespaces().List(context.TODO(),metav1.ListOptions{
		LabelSelector: labelMatcher,
	})
	if err != nil {
		panic(err.Error())
	}
	return namespaces
}
func DeleteNamespace (clientset *kubernetes.Clientset, ns string) bool {
	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), ns,metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
	return true
}
