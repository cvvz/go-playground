package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)
	clientset, _ := kubernetes.NewForConfig(config)

	if _, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "POD_NAME", metav1.GetOptions{}); err != nil {
		var dnsError *net.DNSError
		if errors.Is(err, &net.DNSError{}) {
			fmt.Println("Is DNS error")
		}
		fmt.Printf("Error: %s\n", dnsError)
		if errors.As(err, &dnsError) {
			fmt.Println("As DNS error")
		}
		fmt.Printf("original Error: %s\n", err)
		fmt.Printf("Error: %s\n", dnsError)
	}
}
