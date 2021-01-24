// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubevault.dev/operator/client/clientset/versioned/typed/catalog/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCatalogV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCatalogV1alpha1) VaultServerVersions() v1alpha1.VaultServerVersionInterface {
	return &FakeVaultServerVersions{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCatalogV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
