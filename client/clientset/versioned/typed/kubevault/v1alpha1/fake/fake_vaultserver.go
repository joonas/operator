// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "kubevault.dev/operator/apis/kubevault/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVaultServers implements VaultServerInterface
type FakeVaultServers struct {
	Fake *FakeKubevaultV1alpha1
	ns   string
}

var vaultserversResource = schema.GroupVersionResource{Group: "kubevault.com", Version: "v1alpha1", Resource: "vaultservers"}

var vaultserversKind = schema.GroupVersionKind{Group: "kubevault.com", Version: "v1alpha1", Kind: "VaultServer"}

// Get takes name of the vaultServer, and returns the corresponding vaultServer object, and an error if there is any.
func (c *FakeVaultServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VaultServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vaultserversResource, c.ns, name), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}

// List takes label and field selectors, and returns the list of VaultServers that match those selectors.
func (c *FakeVaultServers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VaultServerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vaultserversResource, vaultserversKind, c.ns, opts), &v1alpha1.VaultServerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VaultServerList{ListMeta: obj.(*v1alpha1.VaultServerList).ListMeta}
	for _, item := range obj.(*v1alpha1.VaultServerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vaultServers.
func (c *FakeVaultServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vaultserversResource, c.ns, opts))

}

// Create takes the representation of a vaultServer and creates it.  Returns the server's representation of the vaultServer, and an error, if there is any.
func (c *FakeVaultServers) Create(ctx context.Context, vaultServer *v1alpha1.VaultServer, opts v1.CreateOptions) (result *v1alpha1.VaultServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vaultserversResource, c.ns, vaultServer), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}

// Update takes the representation of a vaultServer and updates it. Returns the server's representation of the vaultServer, and an error, if there is any.
func (c *FakeVaultServers) Update(ctx context.Context, vaultServer *v1alpha1.VaultServer, opts v1.UpdateOptions) (result *v1alpha1.VaultServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vaultserversResource, c.ns, vaultServer), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVaultServers) UpdateStatus(ctx context.Context, vaultServer *v1alpha1.VaultServer, opts v1.UpdateOptions) (*v1alpha1.VaultServer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vaultserversResource, "status", c.ns, vaultServer), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}

// Delete takes name of the vaultServer and deletes it. Returns an error if one occurs.
func (c *FakeVaultServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vaultserversResource, c.ns, name), &v1alpha1.VaultServer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVaultServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vaultserversResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VaultServerList{})
	return err
}

// Patch applies the patch and returns the patched vaultServer.
func (c *FakeVaultServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VaultServer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vaultserversResource, c.ns, name, pt, data, subresources...), &v1alpha1.VaultServer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultServer), err
}
