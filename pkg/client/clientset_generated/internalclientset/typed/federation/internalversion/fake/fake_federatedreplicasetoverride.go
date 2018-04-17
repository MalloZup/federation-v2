/*
Copyright 2018 The Federation v2 Authors.

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
package fake

import (
	federation "github.com/marun/federation-v2/pkg/apis/federation"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedReplicaSetOverrides implements FederatedReplicaSetOverrideInterface
type FakeFederatedReplicaSetOverrides struct {
	Fake *FakeFederation
	ns   string
}

var federatedreplicasetoverridesResource = schema.GroupVersionResource{Group: "federation.k8s.io", Version: "", Resource: "federatedreplicasetoverrides"}

var federatedreplicasetoverridesKind = schema.GroupVersionKind{Group: "federation.k8s.io", Version: "", Kind: "FederatedReplicaSetOverride"}

// Get takes name of the federatedReplicaSetOverride, and returns the corresponding federatedReplicaSetOverride object, and an error if there is any.
func (c *FakeFederatedReplicaSetOverrides) Get(name string, options v1.GetOptions) (result *federation.FederatedReplicaSetOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(federatedreplicasetoverridesResource, c.ns, name), &federation.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federation.FederatedReplicaSetOverride), err
}

// List takes label and field selectors, and returns the list of FederatedReplicaSetOverrides that match those selectors.
func (c *FakeFederatedReplicaSetOverrides) List(opts v1.ListOptions) (result *federation.FederatedReplicaSetOverrideList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(federatedreplicasetoverridesResource, federatedreplicasetoverridesKind, c.ns, opts), &federation.FederatedReplicaSetOverrideList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &federation.FederatedReplicaSetOverrideList{}
	for _, item := range obj.(*federation.FederatedReplicaSetOverrideList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedReplicaSetOverrides.
func (c *FakeFederatedReplicaSetOverrides) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(federatedreplicasetoverridesResource, c.ns, opts))

}

// Create takes the representation of a federatedReplicaSetOverride and creates it.  Returns the server's representation of the federatedReplicaSetOverride, and an error, if there is any.
func (c *FakeFederatedReplicaSetOverrides) Create(federatedReplicaSetOverride *federation.FederatedReplicaSetOverride) (result *federation.FederatedReplicaSetOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(federatedreplicasetoverridesResource, c.ns, federatedReplicaSetOverride), &federation.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federation.FederatedReplicaSetOverride), err
}

// Update takes the representation of a federatedReplicaSetOverride and updates it. Returns the server's representation of the federatedReplicaSetOverride, and an error, if there is any.
func (c *FakeFederatedReplicaSetOverrides) Update(federatedReplicaSetOverride *federation.FederatedReplicaSetOverride) (result *federation.FederatedReplicaSetOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(federatedreplicasetoverridesResource, c.ns, federatedReplicaSetOverride), &federation.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federation.FederatedReplicaSetOverride), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedReplicaSetOverrides) UpdateStatus(federatedReplicaSetOverride *federation.FederatedReplicaSetOverride) (*federation.FederatedReplicaSetOverride, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(federatedreplicasetoverridesResource, "status", c.ns, federatedReplicaSetOverride), &federation.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federation.FederatedReplicaSetOverride), err
}

// Delete takes name of the federatedReplicaSetOverride and deletes it. Returns an error if one occurs.
func (c *FakeFederatedReplicaSetOverrides) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(federatedreplicasetoverridesResource, c.ns, name), &federation.FederatedReplicaSetOverride{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedReplicaSetOverrides) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(federatedreplicasetoverridesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &federation.FederatedReplicaSetOverrideList{})
	return err
}

// Patch applies the patch and returns the patched federatedReplicaSetOverride.
func (c *FakeFederatedReplicaSetOverrides) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *federation.FederatedReplicaSetOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(federatedreplicasetoverridesResource, c.ns, name, data, subresources...), &federation.FederatedReplicaSetOverride{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federation.FederatedReplicaSetOverride), err
}
