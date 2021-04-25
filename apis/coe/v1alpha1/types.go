/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// COEClusterParameters are the configurable fields of a COECluster.
type COEClusterParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// COEClusterObservation are the observable fields of a COECluster.
type COEClusterObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A COEClusterSpec defines the desired state of a COECluster.
type COEClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       COEClusterParameters `json:"forProvider"`
}

// A COEClusterStatus represents the observed state of a COECluster.
type COEClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          COEClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A COECluster is an example API type
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type COECluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   COEClusterSpec   `json:"spec"`
	Status COEClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// COEClusterList contains a list of COECluster
type COEClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []COECluster `json:"items"`
}
