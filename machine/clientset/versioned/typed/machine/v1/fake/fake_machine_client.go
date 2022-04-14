// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/client-go/machine/clientset/versioned/typed/machine/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMachineV1 struct {
	*testing.Fake
}

func (c *FakeMachineV1) AWSPlacementGroups(namespace string) v1.AWSPlacementGroupInterface {
	return &FakeAWSPlacementGroups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMachineV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
