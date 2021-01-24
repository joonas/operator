// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	scheme "kubevault.dev/operator/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AzureAccessKeyRequestsGetter has a method to return a AzureAccessKeyRequestInterface.
// A group's client should implement this interface.
type AzureAccessKeyRequestsGetter interface {
	AzureAccessKeyRequests(namespace string) AzureAccessKeyRequestInterface
}

// AzureAccessKeyRequestInterface has methods to work with AzureAccessKeyRequest resources.
type AzureAccessKeyRequestInterface interface {
	Create(ctx context.Context, azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest, opts v1.CreateOptions) (*v1alpha1.AzureAccessKeyRequest, error)
	Update(ctx context.Context, azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest, opts v1.UpdateOptions) (*v1alpha1.AzureAccessKeyRequest, error)
	UpdateStatus(ctx context.Context, azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest, opts v1.UpdateOptions) (*v1alpha1.AzureAccessKeyRequest, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AzureAccessKeyRequest, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AzureAccessKeyRequestList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureAccessKeyRequest, err error)
	AzureAccessKeyRequestExpansion
}

// azureAccessKeyRequests implements AzureAccessKeyRequestInterface
type azureAccessKeyRequests struct {
	client rest.Interface
	ns     string
}

// newAzureAccessKeyRequests returns a AzureAccessKeyRequests
func newAzureAccessKeyRequests(c *EngineV1alpha1Client, namespace string) *azureAccessKeyRequests {
	return &azureAccessKeyRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureAccessKeyRequest, and returns the corresponding azureAccessKeyRequest object, and an error if there is any.
func (c *azureAccessKeyRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureAccessKeyRequests that match those selectors.
func (c *azureAccessKeyRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureAccessKeyRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzureAccessKeyRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureAccessKeyRequests.
func (c *azureAccessKeyRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a azureAccessKeyRequest and creates it.  Returns the server's representation of the azureAccessKeyRequest, and an error, if there is any.
func (c *azureAccessKeyRequests) Create(ctx context.Context, azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest, opts v1.CreateOptions) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureAccessKeyRequest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a azureAccessKeyRequest and updates it. Returns the server's representation of the azureAccessKeyRequest, and an error, if there is any.
func (c *azureAccessKeyRequests) Update(ctx context.Context, azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest, opts v1.UpdateOptions) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(azureAccessKeyRequest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureAccessKeyRequest).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *azureAccessKeyRequests) UpdateStatus(ctx context.Context, azureAccessKeyRequest *v1alpha1.AzureAccessKeyRequest, opts v1.UpdateOptions) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(azureAccessKeyRequest.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureAccessKeyRequest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the azureAccessKeyRequest and deletes it. Returns an error if one occurs.
func (c *azureAccessKeyRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureAccessKeyRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched azureAccessKeyRequest.
func (c *azureAccessKeyRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureAccessKeyRequest, err error) {
	result = &v1alpha1.AzureAccessKeyRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azureaccesskeyrequests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
