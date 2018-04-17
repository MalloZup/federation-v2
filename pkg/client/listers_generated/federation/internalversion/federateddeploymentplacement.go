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

package internalversion

import (
	federation "github.com/marun/federation-v2/pkg/apis/federation"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedDeploymentPlacementLister helps list FederatedDeploymentPlacements.
type FederatedDeploymentPlacementLister interface {
	// List lists all FederatedDeploymentPlacements in the indexer.
	List(selector labels.Selector) (ret []*federation.FederatedDeploymentPlacement, err error)
	// FederatedDeploymentPlacements returns an object that can list and get FederatedDeploymentPlacements.
	FederatedDeploymentPlacements(namespace string) FederatedDeploymentPlacementNamespaceLister
	FederatedDeploymentPlacementListerExpansion
}

// federatedDeploymentPlacementLister implements the FederatedDeploymentPlacementLister interface.
type federatedDeploymentPlacementLister struct {
	indexer cache.Indexer
}

// NewFederatedDeploymentPlacementLister returns a new FederatedDeploymentPlacementLister.
func NewFederatedDeploymentPlacementLister(indexer cache.Indexer) FederatedDeploymentPlacementLister {
	return &federatedDeploymentPlacementLister{indexer: indexer}
}

// List lists all FederatedDeploymentPlacements in the indexer.
func (s *federatedDeploymentPlacementLister) List(selector labels.Selector) (ret []*federation.FederatedDeploymentPlacement, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedDeploymentPlacement))
	})
	return ret, err
}

// FederatedDeploymentPlacements returns an object that can list and get FederatedDeploymentPlacements.
func (s *federatedDeploymentPlacementLister) FederatedDeploymentPlacements(namespace string) FederatedDeploymentPlacementNamespaceLister {
	return federatedDeploymentPlacementNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedDeploymentPlacementNamespaceLister helps list and get FederatedDeploymentPlacements.
type FederatedDeploymentPlacementNamespaceLister interface {
	// List lists all FederatedDeploymentPlacements in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*federation.FederatedDeploymentPlacement, err error)
	// Get retrieves the FederatedDeploymentPlacement from the indexer for a given namespace and name.
	Get(name string) (*federation.FederatedDeploymentPlacement, error)
	FederatedDeploymentPlacementNamespaceListerExpansion
}

// federatedDeploymentPlacementNamespaceLister implements the FederatedDeploymentPlacementNamespaceLister
// interface.
type federatedDeploymentPlacementNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedDeploymentPlacements in the indexer for a given namespace.
func (s federatedDeploymentPlacementNamespaceLister) List(selector labels.Selector) (ret []*federation.FederatedDeploymentPlacement, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedDeploymentPlacement))
	})
	return ret, err
}

// Get retrieves the FederatedDeploymentPlacement from the indexer for a given namespace and name.
func (s federatedDeploymentPlacementNamespaceLister) Get(name string) (*federation.FederatedDeploymentPlacement, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(federation.Resource("federateddeploymentplacement"), name)
	}
	return obj.(*federation.FederatedDeploymentPlacement), nil
}
