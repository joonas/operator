// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePostgresRoles implements PostgresRoleInterface
type FakePostgresRoles struct {
	Fake *FakeEngineV1alpha1
	ns   string
}

var postgresrolesResource = schema.GroupVersionResource{Group: "engine.kubevault.com", Version: "v1alpha1", Resource: "postgresroles"}

var postgresrolesKind = schema.GroupVersionKind{Group: "engine.kubevault.com", Version: "v1alpha1", Kind: "PostgresRole"}

// Get takes name of the postgresRole, and returns the corresponding postgresRole object, and an error if there is any.
func (c *FakePostgresRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PostgresRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(postgresrolesResource, c.ns, name), &v1alpha1.PostgresRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresRole), err
}

// List takes label and field selectors, and returns the list of PostgresRoles that match those selectors.
func (c *FakePostgresRoles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PostgresRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(postgresrolesResource, postgresrolesKind, c.ns, opts), &v1alpha1.PostgresRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PostgresRoleList{ListMeta: obj.(*v1alpha1.PostgresRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.PostgresRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested postgresRoles.
func (c *FakePostgresRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(postgresrolesResource, c.ns, opts))

}

// Create takes the representation of a postgresRole and creates it.  Returns the server's representation of the postgresRole, and an error, if there is any.
func (c *FakePostgresRoles) Create(ctx context.Context, postgresRole *v1alpha1.PostgresRole, opts v1.CreateOptions) (result *v1alpha1.PostgresRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(postgresrolesResource, c.ns, postgresRole), &v1alpha1.PostgresRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresRole), err
}

// Update takes the representation of a postgresRole and updates it. Returns the server's representation of the postgresRole, and an error, if there is any.
func (c *FakePostgresRoles) Update(ctx context.Context, postgresRole *v1alpha1.PostgresRole, opts v1.UpdateOptions) (result *v1alpha1.PostgresRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(postgresrolesResource, c.ns, postgresRole), &v1alpha1.PostgresRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresRole), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePostgresRoles) UpdateStatus(ctx context.Context, postgresRole *v1alpha1.PostgresRole, opts v1.UpdateOptions) (*v1alpha1.PostgresRole, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(postgresrolesResource, "status", c.ns, postgresRole), &v1alpha1.PostgresRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresRole), err
}

// Delete takes name of the postgresRole and deletes it. Returns an error if one occurs.
func (c *FakePostgresRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(postgresrolesResource, c.ns, name), &v1alpha1.PostgresRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePostgresRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(postgresrolesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PostgresRoleList{})
	return err
}

// Patch applies the patch and returns the patched postgresRole.
func (c *FakePostgresRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PostgresRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(postgresrolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PostgresRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PostgresRole), err
}
