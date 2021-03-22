/*
Copyright 2020 The OpenYurt Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "github.com/erda-project/erda/pkg/clientgo/apis/openyurt/v1alpha1"
	scheme "github.com/erda-project/erda/pkg/clientgo/clientset/versioned/scheme"
)

// NodePoolsGetter has a method to return a NodePoolInterface.
// A group's client should implement this interface.
type NodePoolsGetter interface {
	NodePools() NodePoolInterface
}

// NodePoolInterface has methods to work with NodePool resources.
type NodePoolInterface interface {
	Create(context.Context, *v1alpha1.NodePool) (*v1alpha1.NodePool, error)
	Update(context.Context, *v1alpha1.NodePool) (*v1alpha1.NodePool, error)
	UpdateStatus(context.Context, *v1alpha1.NodePool) (*v1alpha1.NodePool, error)
	Delete(ctx context.Context, name string, options *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(ctx context.Context, name string, options v1.GetOptions) (*v1alpha1.NodePool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NodePoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodePool, err error)
	NodePoolExpansion
}

// nodePools implements NodePoolInterface
type nodePools struct {
	client rest.Interface
}

// newNodePools returns a NodePools
func newNodePools(c *AppsV1alpha1Client) *nodePools {
	return &nodePools{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodePool, and returns the corresponding nodePool object, and an error if there is any.
func (c *nodePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Get().
		Resource("nodepools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodePools that match those selectors.
func (c *nodePools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodePoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodePoolList{}
	err = c.client.Get().
		Resource("nodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodePools.
func (c *nodePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodePool and creates it.  Returns the server's representation of the nodePool, and an error, if there is any.
func (c *nodePools) Create(ctx context.Context, nodePool *v1alpha1.NodePool) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Post().
		Resource("nodepools").
		Body(nodePool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodePool and updates it. Returns the server's representation of the nodePool, and an error, if there is any.
func (c *nodePools) Update(ctx context.Context, nodePool *v1alpha1.NodePool) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Put().
		Resource("nodepools").
		Name(nodePool.Name).
		Body(nodePool).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodePools) UpdateStatus(ctx context.Context, nodePool *v1alpha1.NodePool) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Put().
		Resource("nodepools").
		Name(nodePool.Name).
		SubResource("status").
		Body(nodePool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodePool and deletes it. Returns an error if one occurs.
func (c *nodePools) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodepools").
		Name(name).
		Body(options).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodePools) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("nodepools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodePool.
func (c *nodePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodePool, err error) {
	result = &v1alpha1.NodePool{}
	err = c.client.Patch(pt).
		Resource("nodepools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(ctx).
		Into(result)
	return
}