/*
Copyright 2018 The Knative Authors

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
	v1alpha1 "github.com/knative/observability/pkg/apis/sink/v1alpha1"
	scheme "github.com/knative/observability/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterLogSinksGetter has a method to return a ClusterLogSinkInterface.
// A group's client should implement this interface.
type ClusterLogSinksGetter interface {
	ClusterLogSinks(namespace string) ClusterLogSinkInterface
}

// ClusterLogSinkInterface has methods to work with ClusterLogSink resources.
type ClusterLogSinkInterface interface {
	Create(*v1alpha1.ClusterLogSink) (*v1alpha1.ClusterLogSink, error)
	Update(*v1alpha1.ClusterLogSink) (*v1alpha1.ClusterLogSink, error)
	UpdateStatus(*v1alpha1.ClusterLogSink) (*v1alpha1.ClusterLogSink, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterLogSink, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterLogSinkList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterLogSink, err error)
	ClusterLogSinkExpansion
}

// clusterLogSinks implements ClusterLogSinkInterface
type clusterLogSinks struct {
	client rest.Interface
	ns     string
}

// newClusterLogSinks returns a ClusterLogSinks
func newClusterLogSinks(c *ObservabilityV1alpha1Client, namespace string) *clusterLogSinks {
	return &clusterLogSinks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterLogSink, and returns the corresponding clusterLogSink object, and an error if there is any.
func (c *clusterLogSinks) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterLogSink, err error) {
	result = &v1alpha1.ClusterLogSink{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterlogsinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterLogSinks that match those selectors.
func (c *clusterLogSinks) List(opts v1.ListOptions) (result *v1alpha1.ClusterLogSinkList, err error) {
	result = &v1alpha1.ClusterLogSinkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterlogsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterLogSinks.
func (c *clusterLogSinks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterlogsinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterLogSink and creates it.  Returns the server's representation of the clusterLogSink, and an error, if there is any.
func (c *clusterLogSinks) Create(clusterLogSink *v1alpha1.ClusterLogSink) (result *v1alpha1.ClusterLogSink, err error) {
	result = &v1alpha1.ClusterLogSink{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterlogsinks").
		Body(clusterLogSink).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterLogSink and updates it. Returns the server's representation of the clusterLogSink, and an error, if there is any.
func (c *clusterLogSinks) Update(clusterLogSink *v1alpha1.ClusterLogSink) (result *v1alpha1.ClusterLogSink, err error) {
	result = &v1alpha1.ClusterLogSink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterlogsinks").
		Name(clusterLogSink.Name).
		Body(clusterLogSink).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterLogSinks) UpdateStatus(clusterLogSink *v1alpha1.ClusterLogSink) (result *v1alpha1.ClusterLogSink, err error) {
	result = &v1alpha1.ClusterLogSink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterlogsinks").
		Name(clusterLogSink.Name).
		SubResource("status").
		Body(clusterLogSink).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterLogSink and deletes it. Returns an error if one occurs.
func (c *clusterLogSinks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterlogsinks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterLogSinks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterlogsinks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterLogSink.
func (c *clusterLogSinks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterLogSink, err error) {
	result = &v1alpha1.ClusterLogSink{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterlogsinks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
