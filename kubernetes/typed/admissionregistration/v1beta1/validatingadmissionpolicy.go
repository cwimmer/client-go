/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1beta1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// ValidatingAdmissionPoliciesGetter has a method to return a ValidatingAdmissionPolicyInterface.
// A group's client should implement this interface.
type ValidatingAdmissionPoliciesGetter interface {
	ValidatingAdmissionPolicies() ValidatingAdmissionPolicyInterface
}

// ValidatingAdmissionPolicyInterface has methods to work with ValidatingAdmissionPolicy resources.
type ValidatingAdmissionPolicyInterface interface {
	Create(ctx context.Context, validatingAdmissionPolicy *v1beta1.ValidatingAdmissionPolicy, opts v1.CreateOptions) (*v1beta1.ValidatingAdmissionPolicy, error)
	Update(ctx context.Context, validatingAdmissionPolicy *v1beta1.ValidatingAdmissionPolicy, opts v1.UpdateOptions) (*v1beta1.ValidatingAdmissionPolicy, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, validatingAdmissionPolicy *v1beta1.ValidatingAdmissionPolicy, opts v1.UpdateOptions) (*v1beta1.ValidatingAdmissionPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ValidatingAdmissionPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ValidatingAdmissionPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ValidatingAdmissionPolicy, err error)
	Apply(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1beta1.ValidatingAdmissionPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.ValidatingAdmissionPolicy, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, validatingAdmissionPolicy *admissionregistrationv1beta1.ValidatingAdmissionPolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.ValidatingAdmissionPolicy, err error)
	ValidatingAdmissionPolicyExpansion
}

// validatingAdmissionPolicies implements ValidatingAdmissionPolicyInterface
type validatingAdmissionPolicies struct {
	*gentype.ClientWithListAndApply[*v1beta1.ValidatingAdmissionPolicy, *v1beta1.ValidatingAdmissionPolicyList, *admissionregistrationv1beta1.ValidatingAdmissionPolicyApplyConfiguration]
}

// newValidatingAdmissionPolicies returns a ValidatingAdmissionPolicies
func newValidatingAdmissionPolicies(c *AdmissionregistrationV1beta1Client) *validatingAdmissionPolicies {
	return &validatingAdmissionPolicies{
		gentype.NewClientWithListAndApply[*v1beta1.ValidatingAdmissionPolicy, *v1beta1.ValidatingAdmissionPolicyList, *admissionregistrationv1beta1.ValidatingAdmissionPolicyApplyConfiguration](
			"validatingadmissionpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.ValidatingAdmissionPolicy { return &v1beta1.ValidatingAdmissionPolicy{} },
			func() *v1beta1.ValidatingAdmissionPolicyList { return &v1beta1.ValidatingAdmissionPolicyList{} }),
	}
}
