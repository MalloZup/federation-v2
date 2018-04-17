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

// This file was automatically generated by informer-gen

package internalversion

import (
	federation "github.com/marun/federation-v2/pkg/apis/federation"
	internalclientset "github.com/marun/federation-v2/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/marun/federation-v2/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/marun/federation-v2/pkg/client/listers_generated/federation/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// FederatedDeploymentPlacementInformer provides access to a shared informer and lister for
// FederatedDeploymentPlacements.
type FederatedDeploymentPlacementInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.FederatedDeploymentPlacementLister
}

type federatedDeploymentPlacementInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedDeploymentPlacementInformer constructs a new informer for FederatedDeploymentPlacement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedDeploymentPlacementInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedDeploymentPlacementInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedDeploymentPlacementInformer constructs a new informer for FederatedDeploymentPlacement type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedDeploymentPlacementInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Federation().FederatedDeploymentPlacements(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Federation().FederatedDeploymentPlacements(namespace).Watch(options)
			},
		},
		&federation.FederatedDeploymentPlacement{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedDeploymentPlacementInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedDeploymentPlacementInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedDeploymentPlacementInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federation.FederatedDeploymentPlacement{}, f.defaultInformer)
}

func (f *federatedDeploymentPlacementInformer) Lister() internalversion.FederatedDeploymentPlacementLister {
	return internalversion.NewFederatedDeploymentPlacementLister(f.Informer().GetIndexer())
}
