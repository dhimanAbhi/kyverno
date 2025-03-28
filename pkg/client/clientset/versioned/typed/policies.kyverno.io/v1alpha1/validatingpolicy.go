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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kyverno/kyverno/api/policies.kyverno.io/v1alpha1"
	scheme "github.com/kyverno/kyverno/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ValidatingPoliciesGetter has a method to return a ValidatingPolicyInterface.
// A group's client should implement this interface.
type ValidatingPoliciesGetter interface {
	ValidatingPolicies() ValidatingPolicyInterface
}

// ValidatingPolicyInterface has methods to work with ValidatingPolicy resources.
type ValidatingPolicyInterface interface {
	Create(ctx context.Context, validatingPolicy *v1alpha1.ValidatingPolicy, opts v1.CreateOptions) (*v1alpha1.ValidatingPolicy, error)
	Update(ctx context.Context, validatingPolicy *v1alpha1.ValidatingPolicy, opts v1.UpdateOptions) (*v1alpha1.ValidatingPolicy, error)
	UpdateStatus(ctx context.Context, validatingPolicy *v1alpha1.ValidatingPolicy, opts v1.UpdateOptions) (*v1alpha1.ValidatingPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ValidatingPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ValidatingPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ValidatingPolicy, err error)
	ValidatingPolicyExpansion
}

// validatingPolicies implements ValidatingPolicyInterface
type validatingPolicies struct {
	client rest.Interface
}

// newValidatingPolicies returns a ValidatingPolicies
func newValidatingPolicies(c *PoliciesV1alpha1Client) *validatingPolicies {
	return &validatingPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the validatingPolicy, and returns the corresponding validatingPolicy object, and an error if there is any.
func (c *validatingPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ValidatingPolicy, err error) {
	result = &v1alpha1.ValidatingPolicy{}
	err = c.client.Get().
		Resource("validatingpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ValidatingPolicies that match those selectors.
func (c *validatingPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ValidatingPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ValidatingPolicyList{}
	err = c.client.Get().
		Resource("validatingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested validatingPolicies.
func (c *validatingPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("validatingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a validatingPolicy and creates it.  Returns the server's representation of the validatingPolicy, and an error, if there is any.
func (c *validatingPolicies) Create(ctx context.Context, validatingPolicy *v1alpha1.ValidatingPolicy, opts v1.CreateOptions) (result *v1alpha1.ValidatingPolicy, err error) {
	result = &v1alpha1.ValidatingPolicy{}
	err = c.client.Post().
		Resource("validatingpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(validatingPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a validatingPolicy and updates it. Returns the server's representation of the validatingPolicy, and an error, if there is any.
func (c *validatingPolicies) Update(ctx context.Context, validatingPolicy *v1alpha1.ValidatingPolicy, opts v1.UpdateOptions) (result *v1alpha1.ValidatingPolicy, err error) {
	result = &v1alpha1.ValidatingPolicy{}
	err = c.client.Put().
		Resource("validatingpolicies").
		Name(validatingPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(validatingPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *validatingPolicies) UpdateStatus(ctx context.Context, validatingPolicy *v1alpha1.ValidatingPolicy, opts v1.UpdateOptions) (result *v1alpha1.ValidatingPolicy, err error) {
	result = &v1alpha1.ValidatingPolicy{}
	err = c.client.Put().
		Resource("validatingpolicies").
		Name(validatingPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(validatingPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the validatingPolicy and deletes it. Returns an error if one occurs.
func (c *validatingPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("validatingpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *validatingPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("validatingpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched validatingPolicy.
func (c *validatingPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ValidatingPolicy, err error) {
	result = &v1alpha1.ValidatingPolicy{}
	err = c.client.Patch(pt).
		Resource("validatingpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
