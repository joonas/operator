// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRoleSubjectRef) DeepCopyInto(out *AppRoleSubjectRef) {
	*out = *in
	if in.SecretIDBoundCidrs != nil {
		in, out := &in.SecretIDBoundCidrs, &out.SecretIDBoundCidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenPolicies != nil {
		in, out := &in.TokenPolicies, &out.TokenPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenBoundCidrs != nil {
		in, out := &in.TokenBoundCidrs, &out.TokenBoundCidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRoleSubjectRef.
func (in *AppRoleSubjectRef) DeepCopy() *AppRoleSubjectRef {
	if in == nil {
		return nil
	}
	out := new(AppRoleSubjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JWTSubjectRef) DeepCopyInto(out *JWTSubjectRef) {
	*out = *in
	if in.BoundAudiences != nil {
		in, out := &in.BoundAudiences, &out.BoundAudiences
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.BoundClaims != nil {
		in, out := &in.BoundClaims, &out.BoundClaims
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ClaimMappings != nil {
		in, out := &in.ClaimMappings, &out.ClaimMappings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OIDCScopes != nil {
		in, out := &in.OIDCScopes, &out.OIDCScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowedRedirectUris != nil {
		in, out := &in.AllowedRedirectUris, &out.AllowedRedirectUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenPolicies != nil {
		in, out := &in.TokenPolicies, &out.TokenPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TokenBoundCidrs != nil {
		in, out := &in.TokenBoundCidrs, &out.TokenBoundCidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JWTSubjectRef.
func (in *JWTSubjectRef) DeepCopy() *JWTSubjectRef {
	if in == nil {
		return nil
	}
	out := new(JWTSubjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesSubjectRef) DeepCopyInto(out *KubernetesSubjectRef) {
	*out = *in
	if in.ServiceAccountNames != nil {
		in, out := &in.ServiceAccountNames, &out.ServiceAccountNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceAccountNamespaces != nil {
		in, out := &in.ServiceAccountNamespaces, &out.ServiceAccountNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesSubjectRef.
func (in *KubernetesSubjectRef) DeepCopy() *KubernetesSubjectRef {
	if in == nil {
		return nil
	}
	out := new(KubernetesSubjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LdapGroupSubjectRef) DeepCopyInto(out *LdapGroupSubjectRef) {
	*out = *in
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LdapGroupSubjectRef.
func (in *LdapGroupSubjectRef) DeepCopy() *LdapGroupSubjectRef {
	if in == nil {
		return nil
	}
	out := new(LdapGroupSubjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LdapUserSubjectRef) DeepCopyInto(out *LdapUserSubjectRef) {
	*out = *in
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LdapUserSubjectRef.
func (in *LdapUserSubjectRef) DeepCopy() *LdapUserSubjectRef {
	if in == nil {
		return nil
	}
	out := new(LdapUserSubjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyIdentifier) DeepCopyInto(out *PolicyIdentifier) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyIdentifier.
func (in *PolicyIdentifier) DeepCopy() *PolicyIdentifier {
	if in == nil {
		return nil
	}
	out := new(PolicyIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountReference) DeepCopyInto(out *ServiceAccountReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountReference.
func (in *ServiceAccountReference) DeepCopy() *ServiceAccountReference {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectRef) DeepCopyInto(out *SubjectRef) {
	*out = *in
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(KubernetesSubjectRef)
		(*in).DeepCopyInto(*out)
	}
	if in.AppRole != nil {
		in, out := &in.AppRole, &out.AppRole
		*out = new(AppRoleSubjectRef)
		(*in).DeepCopyInto(*out)
	}
	if in.LdapGroup != nil {
		in, out := &in.LdapGroup, &out.LdapGroup
		*out = new(LdapGroupSubjectRef)
		(*in).DeepCopyInto(*out)
	}
	if in.LdapUser != nil {
		in, out := &in.LdapUser, &out.LdapUser
		*out = new(LdapUserSubjectRef)
		(*in).DeepCopyInto(*out)
	}
	if in.JWT != nil {
		in, out := &in.JWT, &out.JWT
		*out = new(JWTSubjectRef)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectRef.
func (in *SubjectRef) DeepCopy() *SubjectRef {
	if in == nil {
		return nil
	}
	out := new(SubjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicy) DeepCopyInto(out *VaultPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicy.
func (in *VaultPolicy) DeepCopy() *VaultPolicy {
	if in == nil {
		return nil
	}
	out := new(VaultPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicyBinding) DeepCopyInto(out *VaultPolicyBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicyBinding.
func (in *VaultPolicyBinding) DeepCopy() *VaultPolicyBinding {
	if in == nil {
		return nil
	}
	out := new(VaultPolicyBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultPolicyBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicyBindingList) DeepCopyInto(out *VaultPolicyBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultPolicyBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicyBindingList.
func (in *VaultPolicyBindingList) DeepCopy() *VaultPolicyBindingList {
	if in == nil {
		return nil
	}
	out := new(VaultPolicyBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultPolicyBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicyBindingSpec) DeepCopyInto(out *VaultPolicyBindingSpec) {
	*out = *in
	out.VaultRef = in.VaultRef
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]PolicyIdentifier, len(*in))
		copy(*out, *in)
	}
	in.SubjectRef.DeepCopyInto(&out.SubjectRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicyBindingSpec.
func (in *VaultPolicyBindingSpec) DeepCopy() *VaultPolicyBindingSpec {
	if in == nil {
		return nil
	}
	out := new(VaultPolicyBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicyBindingStatus) DeepCopyInto(out *VaultPolicyBindingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicyBindingStatus.
func (in *VaultPolicyBindingStatus) DeepCopy() *VaultPolicyBindingStatus {
	if in == nil {
		return nil
	}
	out := new(VaultPolicyBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicyList) DeepCopyInto(out *VaultPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VaultPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicyList.
func (in *VaultPolicyList) DeepCopy() *VaultPolicyList {
	if in == nil {
		return nil
	}
	out := new(VaultPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VaultPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicySpec) DeepCopyInto(out *VaultPolicySpec) {
	*out = *in
	out.VaultRef = in.VaultRef
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicySpec.
func (in *VaultPolicySpec) DeepCopy() *VaultPolicySpec {
	if in == nil {
		return nil
	}
	out := new(VaultPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultPolicyStatus) DeepCopyInto(out *VaultPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultPolicyStatus.
func (in *VaultPolicyStatus) DeepCopy() *VaultPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(VaultPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
