// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubevault.dev/operator/client/clientset/versioned/typed/policy/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePolicyV1alpha1 struct {
	*testing.Fake
}

func (c *FakePolicyV1alpha1) VaultPolicies(namespace string) v1alpha1.VaultPolicyInterface {
	return &FakeVaultPolicies{c, namespace}
}

func (c *FakePolicyV1alpha1) VaultPolicyBindings(namespace string) v1alpha1.VaultPolicyBindingInterface {
	return &FakeVaultPolicyBindings{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePolicyV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
