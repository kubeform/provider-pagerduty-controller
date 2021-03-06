//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleList) DeepCopyInto(out *RuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleList.
func (in *RuleList) DeepCopy() *RuleList {
	if in == nil {
		return nil
	}
	out := new(RuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpec) DeepCopyInto(out *RuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RuleSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpec.
func (in *RuleSpec) DeepCopy() *RuleSpec {
	if in == nil {
		return nil
	}
	out := new(RuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActions) DeepCopyInto(out *RuleSpecActions) {
	*out = *in
	if in.Annotate != nil {
		in, out := &in.Annotate, &out.Annotate
		*out = make([]RuleSpecActionsAnnotate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventAction != nil {
		in, out := &in.EventAction, &out.EventAction
		*out = make([]RuleSpecActionsEventAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extractions != nil {
		in, out := &in.Extractions, &out.Extractions
		*out = make([]RuleSpecActionsExtractions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = make([]RuleSpecActionsPriority, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = make([]RuleSpecActionsRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = make([]RuleSpecActionsSeverity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Suppress != nil {
		in, out := &in.Suppress, &out.Suppress
		*out = make([]RuleSpecActionsSuppress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = make([]RuleSpecActionsSuspend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActions.
func (in *RuleSpecActions) DeepCopy() *RuleSpecActions {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActionsAnnotate) DeepCopyInto(out *RuleSpecActionsAnnotate) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActionsAnnotate.
func (in *RuleSpecActionsAnnotate) DeepCopy() *RuleSpecActionsAnnotate {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActionsAnnotate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActionsEventAction) DeepCopyInto(out *RuleSpecActionsEventAction) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActionsEventAction.
func (in *RuleSpecActionsEventAction) DeepCopy() *RuleSpecActionsEventAction {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActionsEventAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActionsExtractions) DeepCopyInto(out *RuleSpecActionsExtractions) {
	*out = *in
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActionsExtractions.
func (in *RuleSpecActionsExtractions) DeepCopy() *RuleSpecActionsExtractions {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActionsExtractions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActionsPriority) DeepCopyInto(out *RuleSpecActionsPriority) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActionsPriority.
func (in *RuleSpecActionsPriority) DeepCopy() *RuleSpecActionsPriority {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActionsPriority)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActionsRoute) DeepCopyInto(out *RuleSpecActionsRoute) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActionsRoute.
func (in *RuleSpecActionsRoute) DeepCopy() *RuleSpecActionsRoute {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActionsRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActionsSeverity) DeepCopyInto(out *RuleSpecActionsSeverity) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActionsSeverity.
func (in *RuleSpecActionsSeverity) DeepCopy() *RuleSpecActionsSeverity {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActionsSeverity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActionsSuppress) DeepCopyInto(out *RuleSpecActionsSuppress) {
	*out = *in
	if in.ThresholdTimeAmount != nil {
		in, out := &in.ThresholdTimeAmount, &out.ThresholdTimeAmount
		*out = new(int64)
		**out = **in
	}
	if in.ThresholdTimeUnit != nil {
		in, out := &in.ThresholdTimeUnit, &out.ThresholdTimeUnit
		*out = new(string)
		**out = **in
	}
	if in.ThresholdValue != nil {
		in, out := &in.ThresholdValue, &out.ThresholdValue
		*out = new(int64)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActionsSuppress.
func (in *RuleSpecActionsSuppress) DeepCopy() *RuleSpecActionsSuppress {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActionsSuppress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecActionsSuspend) DeepCopyInto(out *RuleSpecActionsSuspend) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecActionsSuspend.
func (in *RuleSpecActionsSuspend) DeepCopy() *RuleSpecActionsSuspend {
	if in == nil {
		return nil
	}
	out := new(RuleSpecActionsSuspend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecConditions) DeepCopyInto(out *RuleSpecConditions) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Subconditions != nil {
		in, out := &in.Subconditions, &out.Subconditions
		*out = make([]RuleSpecConditionsSubconditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecConditions.
func (in *RuleSpecConditions) DeepCopy() *RuleSpecConditions {
	if in == nil {
		return nil
	}
	out := new(RuleSpecConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecConditionsSubconditions) DeepCopyInto(out *RuleSpecConditionsSubconditions) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make([]RuleSpecConditionsSubconditionsParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecConditionsSubconditions.
func (in *RuleSpecConditionsSubconditions) DeepCopy() *RuleSpecConditionsSubconditions {
	if in == nil {
		return nil
	}
	out := new(RuleSpecConditionsSubconditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecConditionsSubconditionsParameter) DeepCopyInto(out *RuleSpecConditionsSubconditionsParameter) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecConditionsSubconditionsParameter.
func (in *RuleSpecConditionsSubconditionsParameter) DeepCopy() *RuleSpecConditionsSubconditionsParameter {
	if in == nil {
		return nil
	}
	out := new(RuleSpecConditionsSubconditionsParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecResource) DeepCopyInto(out *RuleSpecResource) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = new(RuleSpecActions)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new(RuleSpecConditions)
		(*in).DeepCopyInto(*out)
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
	if in.Position != nil {
		in, out := &in.Position, &out.Position
		*out = new(int64)
		**out = **in
	}
	if in.Ruleset != nil {
		in, out := &in.Ruleset, &out.Ruleset
		*out = new(string)
		**out = **in
	}
	if in.TimeFrame != nil {
		in, out := &in.TimeFrame, &out.TimeFrame
		*out = new(RuleSpecTimeFrame)
		(*in).DeepCopyInto(*out)
	}
	if in.Variable != nil {
		in, out := &in.Variable, &out.Variable
		*out = make([]RuleSpecVariable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecResource.
func (in *RuleSpecResource) DeepCopy() *RuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(RuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecTimeFrame) DeepCopyInto(out *RuleSpecTimeFrame) {
	*out = *in
	if in.ActiveBetween != nil {
		in, out := &in.ActiveBetween, &out.ActiveBetween
		*out = make([]RuleSpecTimeFrameActiveBetween, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScheduledWeekly != nil {
		in, out := &in.ScheduledWeekly, &out.ScheduledWeekly
		*out = make([]RuleSpecTimeFrameScheduledWeekly, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecTimeFrame.
func (in *RuleSpecTimeFrame) DeepCopy() *RuleSpecTimeFrame {
	if in == nil {
		return nil
	}
	out := new(RuleSpecTimeFrame)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecTimeFrameActiveBetween) DeepCopyInto(out *RuleSpecTimeFrameActiveBetween) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(int64)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecTimeFrameActiveBetween.
func (in *RuleSpecTimeFrameActiveBetween) DeepCopy() *RuleSpecTimeFrameActiveBetween {
	if in == nil {
		return nil
	}
	out := new(RuleSpecTimeFrameActiveBetween)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecTimeFrameScheduledWeekly) DeepCopyInto(out *RuleSpecTimeFrameScheduledWeekly) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(int64)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(int64)
		**out = **in
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
	if in.Weekdays != nil {
		in, out := &in.Weekdays, &out.Weekdays
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecTimeFrameScheduledWeekly.
func (in *RuleSpecTimeFrameScheduledWeekly) DeepCopy() *RuleSpecTimeFrameScheduledWeekly {
	if in == nil {
		return nil
	}
	out := new(RuleSpecTimeFrameScheduledWeekly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecVariable) DeepCopyInto(out *RuleSpecVariable) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]RuleSpecVariableParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecVariable.
func (in *RuleSpecVariable) DeepCopy() *RuleSpecVariable {
	if in == nil {
		return nil
	}
	out := new(RuleSpecVariable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleSpecVariableParameters) DeepCopyInto(out *RuleSpecVariableParameters) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleSpecVariableParameters.
func (in *RuleSpecVariableParameters) DeepCopy() *RuleSpecVariableParameters {
	if in == nil {
		return nil
	}
	out := new(RuleSpecVariableParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleStatus) DeepCopyInto(out *RuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleStatus.
func (in *RuleStatus) DeepCopy() *RuleStatus {
	if in == nil {
		return nil
	}
	out := new(RuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ruleset) DeepCopyInto(out *Ruleset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ruleset.
func (in *Ruleset) DeepCopy() *Ruleset {
	if in == nil {
		return nil
	}
	out := new(Ruleset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Ruleset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesetList) DeepCopyInto(out *RulesetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Ruleset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesetList.
func (in *RulesetList) DeepCopy() *RulesetList {
	if in == nil {
		return nil
	}
	out := new(RulesetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RulesetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesetSpec) DeepCopyInto(out *RulesetSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RulesetSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesetSpec.
func (in *RulesetSpec) DeepCopy() *RulesetSpec {
	if in == nil {
		return nil
	}
	out := new(RulesetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesetSpecResource) DeepCopyInto(out *RulesetSpecResource) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RoutingKeys != nil {
		in, out := &in.RoutingKeys, &out.RoutingKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = new(RulesetSpecTeam)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesetSpecResource.
func (in *RulesetSpecResource) DeepCopy() *RulesetSpecResource {
	if in == nil {
		return nil
	}
	out := new(RulesetSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesetSpecTeam) DeepCopyInto(out *RulesetSpecTeam) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesetSpecTeam.
func (in *RulesetSpecTeam) DeepCopy() *RulesetSpecTeam {
	if in == nil {
		return nil
	}
	out := new(RulesetSpecTeam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesetStatus) DeepCopyInto(out *RulesetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesetStatus.
func (in *RulesetStatus) DeepCopy() *RulesetStatus {
	if in == nil {
		return nil
	}
	out := new(RulesetStatus)
	in.DeepCopyInto(out)
	return out
}
