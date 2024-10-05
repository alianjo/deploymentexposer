package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", os.Getenv("HOME")+"/.kube/config", "path to the kubeconfig file")

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Print("Error %v in building clientcmd", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Print("Error %v in building in cluster config", err.Error())
		}

	}
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Print("Error %v in building clientset", err.Error())
	}
	ch := make(chan struct{})

	informers := informers.NewSharedInformerFactory(clientset, 2*time.Minute)

	c := newController(clientset, informers.Apps().V1().Deployments(), informers.Core().V1().Services(), informers.Networking().V1().Ingresses())
	informers.Start(ch)
	c.run(ch)
}
