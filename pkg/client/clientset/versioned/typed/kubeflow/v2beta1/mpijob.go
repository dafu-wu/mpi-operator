// Copyright 2023 The Kubeflow Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v2beta1

import (
	"context"
	"time"

	v2beta1 "github.com/dafu-wu/mpi-operator/pkg/apis/kubeflow/v2beta1"
	scheme "github.com/dafu-wu/mpi-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MPIJobsGetter has a method to return a MPIJobInterface.
// A group's client should implement this interface.
type MPIJobsGetter interface {
	MPIJobs(namespace string) MPIJobInterface
}

// MPIJobInterface has methods to work with MPIJob resources.
type MPIJobInterface interface {
	Create(ctx context.Context, mPIJob *v2beta1.MPIJob, opts v1.CreateOptions) (*v2beta1.MPIJob, error)
	Update(ctx context.Context, mPIJob *v2beta1.MPIJob, opts v1.UpdateOptions) (*v2beta1.MPIJob, error)
	UpdateStatus(ctx context.Context, mPIJob *v2beta1.MPIJob, opts v1.UpdateOptions) (*v2beta1.MPIJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta1.MPIJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta1.MPIJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.MPIJob, err error)
	MPIJobExpansion
}

// mPIJobs implements MPIJobInterface
type mPIJobs struct {
	client rest.Interface
	ns     string
}

// newMPIJobs returns a MPIJobs
func newMPIJobs(c *KubeflowV2beta1Client, namespace string) *mPIJobs {
	return &mPIJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mPIJob, and returns the corresponding mPIJob object, and an error if there is any.
func (c *mPIJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.MPIJob, err error) {
	result = &v2beta1.MPIJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mpijobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MPIJobs that match those selectors.
func (c *mPIJobs) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.MPIJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta1.MPIJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mpijobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mPIJobs.
func (c *mPIJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mpijobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mPIJob and creates it.  Returns the server's representation of the mPIJob, and an error, if there is any.
func (c *mPIJobs) Create(ctx context.Context, mPIJob *v2beta1.MPIJob, opts v1.CreateOptions) (result *v2beta1.MPIJob, err error) {
	result = &v2beta1.MPIJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mpijobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mPIJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mPIJob and updates it. Returns the server's representation of the mPIJob, and an error, if there is any.
func (c *mPIJobs) Update(ctx context.Context, mPIJob *v2beta1.MPIJob, opts v1.UpdateOptions) (result *v2beta1.MPIJob, err error) {
	result = &v2beta1.MPIJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mpijobs").
		Name(mPIJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mPIJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *mPIJobs) UpdateStatus(ctx context.Context, mPIJob *v2beta1.MPIJob, opts v1.UpdateOptions) (result *v2beta1.MPIJob, err error) {
	result = &v2beta1.MPIJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mpijobs").
		Name(mPIJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mPIJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mPIJob and deletes it. Returns an error if one occurs.
func (c *mPIJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mpijobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mPIJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mpijobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mPIJob.
func (c *mPIJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.MPIJob, err error) {
	result = &v2beta1.MPIJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mpijobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
