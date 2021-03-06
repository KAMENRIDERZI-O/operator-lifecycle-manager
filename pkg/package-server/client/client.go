package client

import (
	"k8s.io/client-go/rest"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/client/clientset/versioned"
)

// NewClient creates a client that can interact with the ALM resources in k8s api
func NewClient(kubeconfig string) (client versioned.Interface, err error) {
	var config *rest.Config
	config, err = GetConfig(kubeconfig)
	if err != nil {
		return
	}
	return versioned.NewForConfig(config)
}
