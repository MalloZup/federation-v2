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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/marun/federation-v2/pkg/apis/federation/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedReplicaSetOverrideLister helps list FederatedReplicaSetOverrides.
type FederatedReplicaSetOverrideLister interface {
	// List lists all FederatedReplicaSetOverrides in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedReplicaSetOverride, err error)
	// FederatedReplicaSetOverrides returns an object that can list and get FederatedReplicaSetOverrides.
	FederatedReplicaSetOverrides(namespace string) FederatedReplicaSetOverrideNamespaceLister
	FederatedReplicaSetOverrideListerExpansion
}

// federatedReplicaSetOverrideLister implements the FederatedReplicaSetOverrideLister interface.
type federatedReplicaSetOverrideLister struct {
	indexer cache.Indexer
}

// NewFederatedReplicaSetOverrideLister returns a new FederatedReplicaSetOverrideLister.
func NewFederatedReplicaSetOverrideLister(indexer cache.Indexer) FederatedReplicaSetOverrideLister {
	return &federatedReplicaSetOverrideLister{indexer: indexer}
}

// List lists all FederatedReplicaSetOverrides in the indexer.
func (s *federatedReplicaSetOverrideLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedReplicaSetOverride, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedReplicaSetOverride))
	})
	return ret, err
}

// FederatedReplicaSetOverrides returns an object that can list and get FederatedReplicaSetOverrides.
func (s *federatedReplicaSetOverrideLister) FederatedReplicaSetOverrides(namespace string) FederatedReplicaSetOverrideNamespaceLister {
	return federatedReplicaSetOverrideNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedReplicaSetOverrideNamespaceLister helps list and get FederatedReplicaSetOverrides.
type FederatedReplicaSetOverrideNamespaceLister interface {
	// List lists all FederatedReplicaSetOverrides in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FederatedReplicaSetOverride, err error)
	// Get retrieves the FederatedReplicaSetOverride from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FederatedReplicaSetOverride, error)
	FederatedReplicaSetOverrideNamespaceListerExpansion
}

// federatedReplicaSetOverrideNamespaceLister implements the FederatedReplicaSetOverrideNamespaceLister
// interface.
type federatedReplicaSetOverrideNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedReplicaSetOverrides in the indexer for a given namespace.
func (s federatedReplicaSetOverrideNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FederatedReplicaSetOverride, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FederatedReplicaSetOverride))
	})
	return ret, err
}

// Get retrieves the FederatedReplicaSetOverride from the indexer for a given namespace and name.
func (s federatedReplicaSetOverrideNamespaceLister) Get(name string) (*v1alpha1.FederatedReplicaSetOverride, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("federatedreplicasetoverride"), name)
	}
	return obj.(*v1alpha1.FederatedReplicaSetOverride), nil
}
