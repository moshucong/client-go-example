package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
    "os"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
    var daemonsetName = "gateway"
    var namespace = os.Getenv("NAMESPACE")
    fmt.Printf( "%s , %s " , namespace , daemonsetName )
    ds, err := clientset.AppsV1().DaemonSets(namespace).Get(daemonsetName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Daemonset not found\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting Daemonset %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
        fmt.Printf( "pod num : %d " ,  ds.Status.CurrentNumberScheduled )
	}
}
