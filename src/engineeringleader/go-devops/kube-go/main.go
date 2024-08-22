package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	var (
		client *kubernetes.Clientset
		err    error
		labels map[string]string
	)

	ctx := context.Background()

	if client, err = getClient(); err != nil {
		fmt.Printf("error getting the kube client: %s", err)
		os.Exit(1)
	}

	fmt.Println("Got Kube client successful!")

	if labels, err = deploy(ctx, client); err != nil {
		fmt.Printf("error deploying: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Deployment finished! %v", labels)
}

func getClient() (*kubernetes.Clientset, error) {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func deploy(ctx context.Context, client *kubernetes.Clientset) (map[string]string, error) {
	var deployment *v1.Deployment

	appFile, err := os.ReadFile("app.yaml")
	if err != nil {
		return nil, fmt.Errorf("read file error: %s", err)
	}

	obj, groupVersionKind, err := scheme.Codecs.UniversalDeserializer().Decode(appFile, nil, nil)

	switch obj.(type) {
	case *v1.Deployment:
		deployment = obj.(*v1.Deployment)
	default:
		return nil, fmt.Errorf("unrecognized type: %s", err)
	}

	deploymentResponse, err := client.AppsV1().Deployments("default").Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("deployment error: %s", groupVersionKind)
	}

	return deploymentResponse.Spec.Template.Labels, nil

}
