/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Subscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec,omitempty"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

type SubscriptionSpecDeliveryMethod struct {
	// +optional
	TemporarilyDisabled *bool `json:"temporarilyDisabled,omitempty" tf:"temporarily_disabled"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type SubscriptionSpecFilter struct {
	// +optional
	ID   *string `json:"ID,omitempty" tf:"id"`
	Type *string `json:"type" tf:"type"`
}

type SubscriptionSpec struct {
	State *SubscriptionSpecResource `json:"state,omitempty" tf:"-"`

	Resource SubscriptionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SubscriptionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Active         *bool                            `json:"active,omitempty" tf:"active"`
	DeliveryMethod []SubscriptionSpecDeliveryMethod `json:"deliveryMethod" tf:"delivery_method"`
	// +optional
	Description *string                  `json:"description,omitempty" tf:"description"`
	Events      []string                 `json:"events" tf:"events"`
	Filter      []SubscriptionSpecFilter `json:"filter" tf:"filter"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type SubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SubscriptionList is a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Subscription CRD objects
	Items []Subscription `json:"items,omitempty"`
}