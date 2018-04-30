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
	v1beta1 "k8s.io/api/extensions/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// ReplicationControllersGetter has a method to return a ReplicationControllerInterface.
// A group's client should implement this interface.
type ReplicationControllersGetter interface {
	ReplicationControllers(namespace string) ReplicationControllerInterface
}

// ReplicationControllerInterface has methods to work with ReplicationController resources.
type ReplicationControllerInterface interface {
	Create(ctx context.Context, obj *v1.ReplicationController) (*v1.ReplicationController, error)
	Update(ctx context.Context, obj *v1.ReplicationController) (*v1.ReplicationController, error)
	UpdateStatus(ctx context.Context, obj *v1.ReplicationController) (*v1.ReplicationController, error)
	Delete(ctx context.Context, name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(ctx context.Context, name string, options meta_v1.GetOptions) (*v1.ReplicationController, error)
	List(ctx context.Context, opts meta_v1.ListOptions) (*v1.ReplicationControllerList, error)
	Watch(ctx context.Context, opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ReplicationController, err error)
	GetScale(ctx context.Context, replicationControllerName string, options meta_v1.GetOptions) (*v1beta1.Scale, error)
	UpdateScale(ctx context.Context, replicationControllerName string, scale *v1beta1.Scale) (*v1beta1.Scale, error)

	ReplicationControllerExpansion
}

// replicationControllers implements ReplicationControllerInterface
type replicationControllers struct {
	client rest.Interface
	ns     string
}

// newReplicationControllers returns a ReplicationControllers
func newReplicationControllers(c *CoreV1Client, namespace string) *replicationControllers {
	return &replicationControllers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the replicationController, and returns the corresponding replicationController object, and an error if there is any.
func (c *replicationControllers) Get(ctx context.Context, name string, options meta_v1.GetOptions) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ReplicationControllers that match those selectors.
func (c *replicationControllers) List(ctx context.Context, opts meta_v1.ListOptions) (result *v1.ReplicationControllerList, err error) {
	result = &v1.ReplicationControllerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested replicationControllers.
func (c *replicationControllers) Watch(ctx context.Context, opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Context(ctx).
		Watch()
}

// Create takes the representation of a replicationController and creates it.  Returns the server's representation of the replicationController, and an error, if there is any.
func (c *replicationControllers) Create(ctx context.Context, replicationController *v1.ReplicationController) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Body(replicationController).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Update takes the representation of a replicationController and updates it. Returns the server's representation of the replicationController, and an error, if there is any.
func (c *replicationControllers) Update(ctx context.Context, replicationController *v1.ReplicationController) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(replicationController.Name).
		Body(replicationController).
		Context(ctx).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *replicationControllers) UpdateStatus(ctx context.Context, replicationController *v1.ReplicationController) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(replicationController.Name).
		SubResource("status").
		Body(replicationController).
		Context(ctx).
		Do().
		Into(result)
	return
}

// Delete takes name of the replicationController and deletes it. Returns an error if one occurs.
func (c *replicationControllers) Delete(ctx context.Context, name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(name).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *replicationControllers) DeleteCollection(ctx context.Context, options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Context(ctx).
		Do().
		Error()
}

// Patch applies the patch and returns the patched replicationController.
func (c *replicationControllers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ReplicationController, err error) {
	result = &v1.ReplicationController{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("replicationcontrollers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Context(ctx).
		Do().
		Into(result)
	return
}

// GetScale takes name of the replicationController, and returns the corresponding v1beta1.Scale object, and an error if there is any.
func (c *replicationControllers) GetScale(ctx context.Context, replicationControllerName string, options meta_v1.GetOptions) (result *v1beta1.Scale, err error) {
	result = &v1beta1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(replicationControllerName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Context(ctx).
		Do().
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *replicationControllers) UpdateScale(ctx context.Context, replicationControllerName string, scale *v1beta1.Scale) (result *v1beta1.Scale, err error) {
	result = &v1beta1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("replicationcontrollers").
		Name(replicationControllerName).
		SubResource("scale").
		Body(scale).
		Context(ctx).
		Do().
		Into(result)
	return
}
