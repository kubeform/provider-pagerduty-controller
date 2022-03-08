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

type Service struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec   `json:"spec,omitempty"`
	Status            ServiceStatus `json:"status,omitempty"`
}

type ServiceSpecAlertGroupingParametersConfig struct {
	// +optional
	Aggregate *string `json:"aggregate,omitempty" tf:"aggregate"`
	// +optional
	Fields []string `json:"fields,omitempty" tf:"fields"`
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
}

type ServiceSpecAlertGroupingParameters struct {
	// +optional
	Config *ServiceSpecAlertGroupingParametersConfig `json:"config,omitempty" tf:"config"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServiceSpecIncidentUrgencyRuleDuringSupportHours struct {
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Urgency *string `json:"urgency,omitempty" tf:"urgency"`
}

type ServiceSpecIncidentUrgencyRuleOutsideSupportHours struct {
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Urgency *string `json:"urgency,omitempty" tf:"urgency"`
}

type ServiceSpecIncidentUrgencyRule struct {
	// +optional
	DuringSupportHours *ServiceSpecIncidentUrgencyRuleDuringSupportHours `json:"duringSupportHours,omitempty" tf:"during_support_hours"`
	// +optional
	OutsideSupportHours *ServiceSpecIncidentUrgencyRuleOutsideSupportHours `json:"outsideSupportHours,omitempty" tf:"outside_support_hours"`
	Type                *string                                            `json:"type" tf:"type"`
	// +optional
	Urgency *string `json:"urgency,omitempty" tf:"urgency"`
}

type ServiceSpecScheduledActionsAt struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServiceSpecScheduledActions struct {
	// +optional
	At []ServiceSpecScheduledActionsAt `json:"at,omitempty" tf:"at"`
	// +optional
	ToUrgency *string `json:"toUrgency,omitempty" tf:"to_urgency"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServiceSpecSupportHours struct {
	// +optional
	// +kubebuilder:validation:MaxItems=7
	DaysOfWeek []int64 `json:"daysOfWeek,omitempty" tf:"days_of_week"`
	// +optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time"`
	// +optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time"`
	// +optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServiceSpec struct {
	State *ServiceSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AcknowledgementTimeout *string `json:"acknowledgementTimeout,omitempty" tf:"acknowledgement_timeout"`
	// +optional
	AlertCreation *string `json:"alertCreation,omitempty" tf:"alert_creation"`
	// +optional
	// Deprecated
	AlertGrouping *string `json:"alertGrouping,omitempty" tf:"alert_grouping"`
	// +optional
	AlertGroupingParameters *ServiceSpecAlertGroupingParameters `json:"alertGroupingParameters,omitempty" tf:"alert_grouping_parameters"`
	// +optional
	// Deprecated
	AlertGroupingTimeout *string `json:"alertGroupingTimeout,omitempty" tf:"alert_grouping_timeout"`
	// +optional
	AutoResolveTimeout *string `json:"autoResolveTimeout,omitempty" tf:"auto_resolve_timeout"`
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// +optional
	Description      *string `json:"description,omitempty" tf:"description"`
	EscalationPolicy *string `json:"escalationPolicy" tf:"escalation_policy"`
	// +optional
	HtmlURL *string `json:"htmlURL,omitempty" tf:"html_url"`
	// +optional
	IncidentUrgencyRule *ServiceSpecIncidentUrgencyRule `json:"incidentUrgencyRule,omitempty" tf:"incident_urgency_rule"`
	// +optional
	LastIncidentTimestamp *string `json:"lastIncidentTimestamp,omitempty" tf:"last_incident_timestamp"`
	Name                  *string `json:"name" tf:"name"`
	// +optional
	ScheduledActions []ServiceSpecScheduledActions `json:"scheduledActions,omitempty" tf:"scheduled_actions"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	SupportHours *ServiceSpecSupportHours `json:"supportHours,omitempty" tf:"support_hours"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ServiceStatus struct {
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

// ServiceList is a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Service CRD objects
	Items []Service `json:"items,omitempty"`
}
