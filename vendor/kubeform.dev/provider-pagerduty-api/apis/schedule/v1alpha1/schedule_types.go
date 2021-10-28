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

type Schedule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduleSpec   `json:"spec,omitempty"`
	Status            ScheduleStatus `json:"status,omitempty"`
}

type ScheduleSpecLayerRestriction struct {
	DurationSeconds *int64 `json:"durationSeconds" tf:"duration_seconds"`
	// +optional
	StartDayOfWeek *int64  `json:"startDayOfWeek,omitempty" tf:"start_day_of_week"`
	StartTimeOfDay *string `json:"startTimeOfDay" tf:"start_time_of_day"`
	Type           *string `json:"type" tf:"type"`
}

type ScheduleSpecLayer struct {
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Restriction               []ScheduleSpecLayerRestriction `json:"restriction,omitempty" tf:"restriction"`
	RotationTurnLengthSeconds *int64                         `json:"rotationTurnLengthSeconds" tf:"rotation_turn_length_seconds"`
	RotationVirtualStart      *string                        `json:"rotationVirtualStart" tf:"rotation_virtual_start"`
	Start                     *string                        `json:"start" tf:"start"`
	// +kubebuilder:validation:MinItems=1
	Users []string `json:"users" tf:"users"`
}

type ScheduleSpec struct {
	State *ScheduleSpecResource `json:"state,omitempty" tf:"-"`

	Resource ScheduleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ScheduleSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description *string             `json:"description,omitempty" tf:"description"`
	Layer       []ScheduleSpecLayer `json:"layer" tf:"layer"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Overflow *bool   `json:"overflow,omitempty" tf:"overflow"`
	TimeZone *string `json:"timeZone" tf:"time_zone"`
}

type ScheduleStatus struct {
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

// ScheduleList is a list of Schedules
type ScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Schedule CRD objects
	Items []Schedule `json:"items,omitempty"`
}
