// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GCPRoleLister helps list GCPRoles.
type GCPRoleLister interface {
	// List lists all GCPRoles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GCPRole, err error)
	// GCPRoles returns an object that can list and get GCPRoles.
	GCPRoles(namespace string) GCPRoleNamespaceLister
	GCPRoleListerExpansion
}

// gCPRoleLister implements the GCPRoleLister interface.
type gCPRoleLister struct {
	indexer cache.Indexer
}

// NewGCPRoleLister returns a new GCPRoleLister.
func NewGCPRoleLister(indexer cache.Indexer) GCPRoleLister {
	return &gCPRoleLister{indexer: indexer}
}

// List lists all GCPRoles in the indexer.
func (s *gCPRoleLister) List(selector labels.Selector) (ret []*v1alpha1.GCPRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GCPRole))
	})
	return ret, err
}

// GCPRoles returns an object that can list and get GCPRoles.
func (s *gCPRoleLister) GCPRoles(namespace string) GCPRoleNamespaceLister {
	return gCPRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GCPRoleNamespaceLister helps list and get GCPRoles.
type GCPRoleNamespaceLister interface {
	// List lists all GCPRoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GCPRole, err error)
	// Get retrieves the GCPRole from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GCPRole, error)
	GCPRoleNamespaceListerExpansion
}

// gCPRoleNamespaceLister implements the GCPRoleNamespaceLister
// interface.
type gCPRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GCPRoles in the indexer for a given namespace.
func (s gCPRoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GCPRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GCPRole))
	})
	return ret, err
}

// Get retrieves the GCPRole from the indexer for a given namespace and name.
func (s gCPRoleNamespaceLister) Get(name string) (*v1alpha1.GCPRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gcprole"), name)
	}
	return obj.(*v1alpha1.GCPRole), nil
}
