/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime/schema"
	addonv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/addon/v1alpha1"
	businessv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/business/v1alpha1"
	escalationv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/escalation/v1alpha1"
	eventv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/event/v1alpha1"
	extensionv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/extension/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-pagerduty-api/apis/maintenance/v1alpha1"
	responsev1alpha1 "kubeform.dev/provider-pagerduty-api/apis/response/v1alpha1"
	rulesetv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/ruleset/v1alpha1"
	schedulev1alpha1 "kubeform.dev/provider-pagerduty-api/apis/schedule/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-pagerduty-api/apis/service/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/team/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/user/v1alpha1"
	"kubeform.dev/provider-pagerduty-controller/controllers"
)

type Data struct {
	JsonIt       jsoniter.API
	ResourceType string
}

var (
	allJsonIt = map[schema.GroupVersionResource]Data{
		{
			Group:    "addon.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "addons",
		}: {
			JsonIt:       controllers.GetJSONItr(addonv1alpha1.GetEncoder(), addonv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_addon",
		},
		{
			Group:    "business.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "services",
		}: {
			JsonIt:       controllers.GetJSONItr(businessv1alpha1.GetEncoder(), businessv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_business_service",
		},
		{
			Group:    "escalation.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "policies",
		}: {
			JsonIt:       controllers.GetJSONItr(escalationv1alpha1.GetEncoder(), escalationv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_escalation_policy",
		},
		{
			Group:    "event.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "rules",
		}: {
			JsonIt:       controllers.GetJSONItr(eventv1alpha1.GetEncoder(), eventv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_event_rule",
		},
		{
			Group:    "extension.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "extensions",
		}: {
			JsonIt:       controllers.GetJSONItr(extensionv1alpha1.GetEncoder(), extensionv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_extension",
		},
		{
			Group:    "extension.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "servicenows",
		}: {
			JsonIt:       controllers.GetJSONItr(extensionv1alpha1.GetEncoder(), extensionv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_extension_servicenow",
		},
		{
			Group:    "maintenance.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "windows",
		}: {
			JsonIt:       controllers.GetJSONItr(maintenancev1alpha1.GetEncoder(), maintenancev1alpha1.GetDecoder()),
			ResourceType: "pagerduty_maintenance_window",
		},
		{
			Group:    "response.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "plays",
		}: {
			JsonIt:       controllers.GetJSONItr(responsev1alpha1.GetEncoder(), responsev1alpha1.GetDecoder()),
			ResourceType: "pagerduty_response_play",
		},
		{
			Group:    "ruleset.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "rulesets",
		}: {
			JsonIt:       controllers.GetJSONItr(rulesetv1alpha1.GetEncoder(), rulesetv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_ruleset",
		},
		{
			Group:    "ruleset.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "rules",
		}: {
			JsonIt:       controllers.GetJSONItr(rulesetv1alpha1.GetEncoder(), rulesetv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_ruleset_rule",
		},
		{
			Group:    "schedule.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "schedules",
		}: {
			JsonIt:       controllers.GetJSONItr(schedulev1alpha1.GetEncoder(), schedulev1alpha1.GetDecoder()),
			ResourceType: "pagerduty_schedule",
		},
		{
			Group:    "service.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "services",
		}: {
			JsonIt:       controllers.GetJSONItr(servicev1alpha1.GetEncoder(), servicev1alpha1.GetDecoder()),
			ResourceType: "pagerduty_service",
		},
		{
			Group:    "service.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "dependencies",
		}: {
			JsonIt:       controllers.GetJSONItr(servicev1alpha1.GetEncoder(), servicev1alpha1.GetDecoder()),
			ResourceType: "pagerduty_service_dependency",
		},
		{
			Group:    "service.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "eventrules",
		}: {
			JsonIt:       controllers.GetJSONItr(servicev1alpha1.GetEncoder(), servicev1alpha1.GetDecoder()),
			ResourceType: "pagerduty_service_event_rule",
		},
		{
			Group:    "service.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "integrations",
		}: {
			JsonIt:       controllers.GetJSONItr(servicev1alpha1.GetEncoder(), servicev1alpha1.GetDecoder()),
			ResourceType: "pagerduty_service_integration",
		},
		{
			Group:    "team.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "teams",
		}: {
			JsonIt:       controllers.GetJSONItr(teamv1alpha1.GetEncoder(), teamv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_team",
		},
		{
			Group:    "team.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "memberships",
		}: {
			JsonIt:       controllers.GetJSONItr(teamv1alpha1.GetEncoder(), teamv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_team_membership",
		},
		{
			Group:    "user.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "users",
		}: {
			JsonIt:       controllers.GetJSONItr(userv1alpha1.GetEncoder(), userv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_user",
		},
		{
			Group:    "user.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "contactmethods",
		}: {
			JsonIt:       controllers.GetJSONItr(userv1alpha1.GetEncoder(), userv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_user_contact_method",
		},
		{
			Group:    "user.pagerduty.kubeform.com",
			Version:  "v1alpha1",
			Resource: "notificationrules",
		}: {
			JsonIt:       controllers.GetJSONItr(userv1alpha1.GetEncoder(), userv1alpha1.GetDecoder()),
			ResourceType: "pagerduty_user_notification_rule",
		},
	}
)

func getJsonItAndResType(gvr schema.GroupVersionResource) Data {
	return allJsonIt[gvr]
}
