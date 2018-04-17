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
	internalinterfaces "github.com/marun/federation-v2/pkg/client/informers_generated/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// FederatedClusters returns a FederatedClusterInformer.
	FederatedClusters() FederatedClusterInformer
	// FederatedConfigMaps returns a FederatedConfigMapInformer.
	FederatedConfigMaps() FederatedConfigMapInformer
	// FederatedConfigMapOverrides returns a FederatedConfigMapOverrideInformer.
	FederatedConfigMapOverrides() FederatedConfigMapOverrideInformer
	// FederatedConfigMapPlacements returns a FederatedConfigMapPlacementInformer.
	FederatedConfigMapPlacements() FederatedConfigMapPlacementInformer
	// FederatedDeployments returns a FederatedDeploymentInformer.
	FederatedDeployments() FederatedDeploymentInformer
	// FederatedDeploymentOverrides returns a FederatedDeploymentOverrideInformer.
	FederatedDeploymentOverrides() FederatedDeploymentOverrideInformer
	// FederatedDeploymentPlacements returns a FederatedDeploymentPlacementInformer.
	FederatedDeploymentPlacements() FederatedDeploymentPlacementInformer
	// FederatedJobs returns a FederatedJobInformer.
	FederatedJobs() FederatedJobInformer
	// FederatedJobOverrides returns a FederatedJobOverrideInformer.
	FederatedJobOverrides() FederatedJobOverrideInformer
	// FederatedJobPlacements returns a FederatedJobPlacementInformer.
	FederatedJobPlacements() FederatedJobPlacementInformer
	// FederatedNamespacePlacements returns a FederatedNamespacePlacementInformer.
	FederatedNamespacePlacements() FederatedNamespacePlacementInformer
	// FederatedReplicaSets returns a FederatedReplicaSetInformer.
	FederatedReplicaSets() FederatedReplicaSetInformer
	// FederatedReplicaSetOverrides returns a FederatedReplicaSetOverrideInformer.
	FederatedReplicaSetOverrides() FederatedReplicaSetOverrideInformer
	// FederatedReplicaSetPlacements returns a FederatedReplicaSetPlacementInformer.
	FederatedReplicaSetPlacements() FederatedReplicaSetPlacementInformer
	// FederatedSecrets returns a FederatedSecretInformer.
	FederatedSecrets() FederatedSecretInformer
	// FederatedSecretOverrides returns a FederatedSecretOverrideInformer.
	FederatedSecretOverrides() FederatedSecretOverrideInformer
	// FederatedSecretPlacements returns a FederatedSecretPlacementInformer.
	FederatedSecretPlacements() FederatedSecretPlacementInformer
	// FederatedServices returns a FederatedServiceInformer.
	FederatedServices() FederatedServiceInformer
	// FederatedServicePlacements returns a FederatedServicePlacementInformer.
	FederatedServicePlacements() FederatedServicePlacementInformer
	// PropagatedVersions returns a PropagatedVersionInformer.
	PropagatedVersions() PropagatedVersionInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// FederatedClusters returns a FederatedClusterInformer.
func (v *version) FederatedClusters() FederatedClusterInformer {
	return &federatedClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// FederatedConfigMaps returns a FederatedConfigMapInformer.
func (v *version) FederatedConfigMaps() FederatedConfigMapInformer {
	return &federatedConfigMapInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedConfigMapOverrides returns a FederatedConfigMapOverrideInformer.
func (v *version) FederatedConfigMapOverrides() FederatedConfigMapOverrideInformer {
	return &federatedConfigMapOverrideInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedConfigMapPlacements returns a FederatedConfigMapPlacementInformer.
func (v *version) FederatedConfigMapPlacements() FederatedConfigMapPlacementInformer {
	return &federatedConfigMapPlacementInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedDeployments returns a FederatedDeploymentInformer.
func (v *version) FederatedDeployments() FederatedDeploymentInformer {
	return &federatedDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedDeploymentOverrides returns a FederatedDeploymentOverrideInformer.
func (v *version) FederatedDeploymentOverrides() FederatedDeploymentOverrideInformer {
	return &federatedDeploymentOverrideInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedDeploymentPlacements returns a FederatedDeploymentPlacementInformer.
func (v *version) FederatedDeploymentPlacements() FederatedDeploymentPlacementInformer {
	return &federatedDeploymentPlacementInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedJobs returns a FederatedJobInformer.
func (v *version) FederatedJobs() FederatedJobInformer {
	return &federatedJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedJobOverrides returns a FederatedJobOverrideInformer.
func (v *version) FederatedJobOverrides() FederatedJobOverrideInformer {
	return &federatedJobOverrideInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedJobPlacements returns a FederatedJobPlacementInformer.
func (v *version) FederatedJobPlacements() FederatedJobPlacementInformer {
	return &federatedJobPlacementInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedNamespacePlacements returns a FederatedNamespacePlacementInformer.
func (v *version) FederatedNamespacePlacements() FederatedNamespacePlacementInformer {
	return &federatedNamespacePlacementInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// FederatedReplicaSets returns a FederatedReplicaSetInformer.
func (v *version) FederatedReplicaSets() FederatedReplicaSetInformer {
	return &federatedReplicaSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedReplicaSetOverrides returns a FederatedReplicaSetOverrideInformer.
func (v *version) FederatedReplicaSetOverrides() FederatedReplicaSetOverrideInformer {
	return &federatedReplicaSetOverrideInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedReplicaSetPlacements returns a FederatedReplicaSetPlacementInformer.
func (v *version) FederatedReplicaSetPlacements() FederatedReplicaSetPlacementInformer {
	return &federatedReplicaSetPlacementInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedSecrets returns a FederatedSecretInformer.
func (v *version) FederatedSecrets() FederatedSecretInformer {
	return &federatedSecretInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedSecretOverrides returns a FederatedSecretOverrideInformer.
func (v *version) FederatedSecretOverrides() FederatedSecretOverrideInformer {
	return &federatedSecretOverrideInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedSecretPlacements returns a FederatedSecretPlacementInformer.
func (v *version) FederatedSecretPlacements() FederatedSecretPlacementInformer {
	return &federatedSecretPlacementInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedServices returns a FederatedServiceInformer.
func (v *version) FederatedServices() FederatedServiceInformer {
	return &federatedServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedServicePlacements returns a FederatedServicePlacementInformer.
func (v *version) FederatedServicePlacements() FederatedServicePlacementInformer {
	return &federatedServicePlacementInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PropagatedVersions returns a PropagatedVersionInformer.
func (v *version) PropagatedVersions() PropagatedVersionInformer {
	return &propagatedVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
