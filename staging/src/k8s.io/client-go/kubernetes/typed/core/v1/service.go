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

package v1

import (
	context "context"

	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// ServicesGetter has a method to return a ServiceInterface.
// A group's client should implement this interface.
type ServicesGetter interface {
	Services(namespace string) ServiceInterface
}

// ServiceInterface has methods to work with Service resources.
type ServiceInterface interface {
	Create(ctx context.Context, obj *v1.Service) (*v1.Service, error)
	Update(ctx context.Context, obj *v1.Service) (*v1.Service, error)
	UpdateStatus(ctx context.Context, obj *v1.Service) (*v1.Service, error)
	Delete(ctx context.Context, name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(ctx context.Context, name string, options meta_v1.GetOptions) (*v1.Service, error)
	List(ctx context.Context, opts meta_v1.ListOptions) (*v1.ServiceList, error)
	Watch(ctx context.Context, opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Service, err error)
	ServiceExpansion
}

// services implements ServiceInterface
type services struct {
	client rest.Interface
	ns     string
}

// newServices returns a Services
func newServices(c *CoreV1Client, namespace string) *services {
	return &services{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the service, and returns the corresponding service object, and an error if there is any.
func (c *services) Get(ctx context.Context, name string, options meta_v1.GetOptions) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("services").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Services that match those selectors.
func (c *services) List(ctx context.Context, opts meta_v1.ListOptions) (result *v1.ServiceList, err error) {
	result = &v1.ServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("services").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested services.
func (c *services) Watch(ctx context.Context, opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("services").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Watch()
}

// Create takes the representation of a service and creates it.  Returns the server's representation of the service, and an error, if there is any.
func (c *services) Create(ctx context.Context, service *v1.Service) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("services").
		Body(service).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Update takes the representation of a service and updates it. Returns the server's representation of the service, and an error, if there is any.
func (c *services) Update(ctx context.Context, service *v1.Service) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("services").
		Name(service.Name).
		Body(service).
		Context(ctx).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *services) UpdateStatus(ctx context.Context, service *v1.Service) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("services").
		Name(service.Name).
		SubResource("status").
		Body(service).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Delete takes name of the service and deletes it. Returns an error if one occurs.
func (c *services) Delete(ctx context.Context, name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("services").
		Name(name).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *services) DeleteCollection(ctx context.Context, options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("services").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// Patch applies the patch and returns the patched service.
func (c *services) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("services").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Context(ctx).
		Do().
		Into(result)
	return
}
