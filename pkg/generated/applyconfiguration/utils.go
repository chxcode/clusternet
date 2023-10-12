/*
Copyright The Clusternet Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	v1beta1 "github.com/clusternet/clusternet/pkg/apis/clusters/v1beta1"
	appsv1alpha1 "github.com/clusternet/clusternet/pkg/generated/applyconfiguration/apps/v1alpha1"
	clustersv1beta1 "github.com/clusternet/clusternet/pkg/generated/applyconfiguration/clusters/v1beta1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=apps.clusternet.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("AggregatedStatus"):
		return &appsv1alpha1.AggregatedStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Base"):
		return &appsv1alpha1.BaseApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("BaseSpec"):
		return &appsv1alpha1.BaseSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ChartPullSecret"):
		return &appsv1alpha1.ChartPullSecretApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ChartReference"):
		return &appsv1alpha1.ChartReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Description"):
		return &appsv1alpha1.DescriptionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DescriptionSpec"):
		return &appsv1alpha1.DescriptionSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DescriptionStatus"):
		return &appsv1alpha1.DescriptionStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DividingScheduling"):
		return &appsv1alpha1.DividingSchedulingApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("DynamicDividing"):
		return &appsv1alpha1.DynamicDividingApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Feed"):
		return &appsv1alpha1.FeedApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FeedInventory"):
		return &appsv1alpha1.FeedInventoryApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FeedInventorySpec"):
		return &appsv1alpha1.FeedInventorySpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FeedOrder"):
		return &appsv1alpha1.FeedOrderApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FeedStatus"):
		return &appsv1alpha1.FeedStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("FeedStatusPerCluster"):
		return &appsv1alpha1.FeedStatusPerClusterApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Globalization"):
		return &appsv1alpha1.GlobalizationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GlobalizationSpec"):
		return &appsv1alpha1.GlobalizationSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HelmChart"):
		return &appsv1alpha1.HelmChartApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HelmChartSpec"):
		return &appsv1alpha1.HelmChartSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HelmChartStatus"):
		return &appsv1alpha1.HelmChartStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HelmOptions"):
		return &appsv1alpha1.HelmOptionsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HelmRelease"):
		return &appsv1alpha1.HelmReleaseApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HelmReleaseSpec"):
		return &appsv1alpha1.HelmReleaseSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("HelmReleaseStatus"):
		return &appsv1alpha1.HelmReleaseStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Localization"):
		return &appsv1alpha1.LocalizationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("LocalizationSpec"):
		return &appsv1alpha1.LocalizationSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Manifest"):
		return &appsv1alpha1.ManifestApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManifestStatus"):
		return &appsv1alpha1.ManifestStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("OverrideConfig"):
		return &appsv1alpha1.OverrideConfigApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ReplicaRequirements"):
		return &appsv1alpha1.ReplicaRequirementsApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ReplicaStatus"):
		return &appsv1alpha1.ReplicaStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubGroupStrategy"):
		return &appsv1alpha1.SubGroupStrategyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Subscriber"):
		return &appsv1alpha1.SubscriberApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Subscription"):
		return &appsv1alpha1.SubscriptionApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubscriptionSpec"):
		return &appsv1alpha1.SubscriptionSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SubscriptionStatus"):
		return &appsv1alpha1.SubscriptionStatusApplyConfiguration{}

		// Group=clusters.clusternet.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("ClusterRegistrationRequest"):
		return &clustersv1beta1.ClusterRegistrationRequestApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterRegistrationRequestSpec"):
		return &clustersv1beta1.ClusterRegistrationRequestSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ClusterRegistrationRequestStatus"):
		return &clustersv1beta1.ClusterRegistrationRequestStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ManagedCluster"):
		return &clustersv1beta1.ManagedClusterApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ManagedClusterSpec"):
		return &clustersv1beta1.ManagedClusterSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ManagedClusterStatus"):
		return &clustersv1beta1.ManagedClusterStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("NodeStatistics"):
		return &clustersv1beta1.NodeStatisticsApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("PodStatistics"):
		return &clustersv1beta1.PodStatisticsApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ResourceUsage"):
		return &clustersv1beta1.ResourceUsageApplyConfiguration{}

	}
	return nil
}