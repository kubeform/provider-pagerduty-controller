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

type Play struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlaySpec   `json:"spec,omitempty"`
	Status            PlayStatus `json:"status,omitempty"`
}

type PlaySpecResponderEscalationRuleTarget struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type PlaySpecResponderEscalationRule struct {
	// +optional
	EscalationDelayInMinutes *int64 `json:"escalationDelayInMinutes,omitempty" tf:"escalation_delay_in_minutes"`
	// +optional
	ID     *string                                 `json:"ID,omitempty" tf:"id"`
	Target []PlaySpecResponderEscalationRuleTarget `json:"target" tf:"target"`
}

type PlaySpecResponderService struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type PlaySpecResponderTeam struct {
	// +optional
	ID   *string `json:"ID,omitempty" tf:"id"`
	Type *string `json:"type" tf:"type"`
}

type PlaySpecResponder struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	EscalationRule []PlaySpecResponderEscalationRule `json:"escalationRule,omitempty" tf:"escalation_rule"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NumLoops *int64 `json:"numLoops,omitempty" tf:"num_loops"`
	// +optional
	OnCallHandoffNotifications *string `json:"onCallHandoffNotifications,omitempty" tf:"on_call_handoff_notifications"`
	// +optional
	Service []PlaySpecResponderService `json:"service,omitempty" tf:"service"`
	// +optional
	Team []PlaySpecResponderTeam `json:"team,omitempty" tf:"team"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type PlaySpecSubscriber struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type PlaySpec struct {
	State *PlaySpecResource `json:"state,omitempty" tf:"-"`

	Resource PlaySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PlaySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ConferenceNumber *string `json:"conferenceNumber,omitempty" tf:"conference_number"`
	// +optional
	ConferenceURL *string `json:"conferenceURL,omitempty" tf:"conference_url"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	From        *string `json:"from" tf:"from"`
	Name        *string `json:"name" tf:"name"`
	// +optional
	Responder []PlaySpecResponder `json:"responder,omitempty" tf:"responder"`
	// +optional
	RespondersMessage *string `json:"respondersMessage,omitempty" tf:"responders_message"`
	// +optional
	Runnability *string `json:"runnability,omitempty" tf:"runnability"`
	// +optional
	Subscriber []PlaySpecSubscriber `json:"subscriber,omitempty" tf:"subscriber"`
	// +optional
	SubscribersMessage *string `json:"subscribersMessage,omitempty" tf:"subscribers_message"`
	// +optional
	Team *string `json:"team,omitempty" tf:"team"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type PlayStatus struct {
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

// PlayList is a list of Plays
type PlayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Play CRD objects
	Items []Play `json:"items,omitempty"`
}