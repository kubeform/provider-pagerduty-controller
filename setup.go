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
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gobuffalo/flect"
	pagerduty "github.com/terraform-providers/terraform-provider-pagerduty/pagerduty"
	auditlib "go.bytebuilders.dev/audit/lib"
	arv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	informers "k8s.io/apiextensions-apiserver/pkg/client/informers/externalversions"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
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
	slackv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/slack/v1alpha1"
	tagv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/tag/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/team/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/user/v1alpha1"
	webhookv1alpha1 "kubeform.dev/provider-pagerduty-api/apis/webhook/v1alpha1"
	controllersaddon "kubeform.dev/provider-pagerduty-controller/controllers/addon"
	controllersbusiness "kubeform.dev/provider-pagerduty-controller/controllers/business"
	controllersescalation "kubeform.dev/provider-pagerduty-controller/controllers/escalation"
	controllersevent "kubeform.dev/provider-pagerduty-controller/controllers/event"
	controllersextension "kubeform.dev/provider-pagerduty-controller/controllers/extension"
	controllersmaintenance "kubeform.dev/provider-pagerduty-controller/controllers/maintenance"
	controllersresponse "kubeform.dev/provider-pagerduty-controller/controllers/response"
	controllersruleset "kubeform.dev/provider-pagerduty-controller/controllers/ruleset"
	controllersschedule "kubeform.dev/provider-pagerduty-controller/controllers/schedule"
	controllersservice "kubeform.dev/provider-pagerduty-controller/controllers/service"
	controllersslack "kubeform.dev/provider-pagerduty-controller/controllers/slack"
	controllerstag "kubeform.dev/provider-pagerduty-controller/controllers/tag"
	controllersteam "kubeform.dev/provider-pagerduty-controller/controllers/team"
	controllersuser "kubeform.dev/provider-pagerduty-controller/controllers/user"
	controllerswebhook "kubeform.dev/provider-pagerduty-controller/controllers/webhook"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var _provider = pagerduty.Provider()

var runningControllers = struct {
	sync.RWMutex
	mp map[schema.GroupVersionKind]bool
}{mp: make(map[schema.GroupVersionKind]bool)}

func watchCRD(ctx context.Context, crdClient *clientset.Clientset, vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, stopCh <-chan struct{}, mgr manager.Manager, auditor *auditlib.EventPublisher, restrictToNamespace string) error {
	informerFactory := informers.NewSharedInformerFactory(crdClient, time.Second*30)
	i := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Informer()
	l := informerFactory.Apiextensions().V1().CustomResourceDefinitions().Lister()

	i.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			var key string
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				klog.Error(err)
				return
			}

			_, name, err := cache.SplitMetaNamespaceKey(key)
			if err != nil {
				klog.Error(err)
				return
			}

			crd, err := l.Get(name)
			if err != nil {
				klog.Error(err)
				return
			}
			if strings.Contains(crd.Spec.Group, "pagerduty.kubeform.com") {
				gvk := schema.GroupVersionKind{
					Group:   crd.Spec.Group,
					Version: crd.Spec.Versions[0].Name,
					Kind:    crd.Spec.Names.Kind,
				}

				// check whether this gvk came before, if no then start the controller
				runningControllers.RLock()
				_, ok := runningControllers.mp[gvk]
				runningControllers.RUnlock()

				if !ok {
					runningControllers.Lock()
					runningControllers.mp[gvk] = true
					runningControllers.Unlock()

					if enableValidatingWebhook {
						// add dynamic ValidatingWebhookConfiguration

						// create empty VWC if the group has come for the first time
						err := createEmptyVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						// update
						err = updateVWC(vwcClient, gvk)
						if err != nil {
							klog.Error(err)
							return
						}

						err = SetupWebhook(mgr, gvk)
						if err != nil {
							setupLog.Error(err, "unable to enable webhook")
							os.Exit(1)
						}
					}

					err = SetupManager(ctx, mgr, gvk, auditor, restrictToNamespace)
					if err != nil {
						setupLog.Error(err, "unable to start manager")
						os.Exit(1)
					}
				}
			}
		},
	})

	informerFactory.Start(stopCh)

	return nil
}

func createEmptyVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	_, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err == nil || !(errors.IsNotFound(err)) {
		return err
	}

	emptyVWC := &arv1.ValidatingWebhookConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ValidatingWebhookConfiguration",
			APIVersion: "admissionregistration.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-"),
			Labels: map[string]string{
				"app.kubernetes.io/instance": "pagerduty.kubeform.com",
				"app.kubernetes.io/part-of":  "kubeform.com",
			},
		},
	}
	_, err = vwcClient.ValidatingWebhookConfigurations().Create(context.TODO(), emptyVWC, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func updateVWC(vwcClient *admissionregistrationv1.AdmissionregistrationV1Client, gvk schema.GroupVersionKind) error {
	vwcName := strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-")
	vwc, err := vwcClient.ValidatingWebhookConfigurations().Get(context.TODO(), vwcName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	path := "/validate-" + strings.ReplaceAll(strings.ToLower(gvk.Group), ".", "-") + "-v1alpha1-" + strings.ToLower(gvk.Kind)
	fail := arv1.Fail
	sideEffects := arv1.SideEffectClassNone
	admissionReviewVersions := []string{"v1beta1"}

	rules := []arv1.RuleWithOperations{
		{
			Operations: []arv1.OperationType{
				arv1.Delete,
				arv1.Update,
			},
			Rule: arv1.Rule{
				APIGroups:   []string{strings.ToLower(gvk.Group)},
				APIVersions: []string{gvk.Version},
				Resources:   []string{strings.ToLower(flect.Pluralize(gvk.Kind))},
			},
		},
	}

	data, err := ioutil.ReadFile("/tmp/k8s-webhook-server/serving-certs/ca.crt")
	if err != nil {
		return err
	}

	name := strings.ToLower(gvk.Kind) + "." + gvk.Group
	for _, webhook := range vwc.Webhooks {
		if webhook.Name == name {
			return nil
		}
	}

	newWebhook := arv1.ValidatingWebhook{
		Name: name,
		ClientConfig: arv1.WebhookClientConfig{
			Service: &arv1.ServiceReference{
				Namespace: webhookNamespace,
				Name:      webhookName,
				Path:      &path,
			},
			CABundle: data,
		},
		Rules:                   rules,
		FailurePolicy:           &fail,
		SideEffects:             &sideEffects,
		AdmissionReviewVersions: admissionReviewVersions,
	}

	vwc.Webhooks = append(vwc.Webhooks, newWebhook)

	_, err = vwcClient.ValidatingWebhookConfigurations().Update(context.TODO(), vwc, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func SetupManager(ctx context.Context, mgr manager.Manager, gvk schema.GroupVersionKind, auditor *auditlib.EventPublisher, restrictToNamespace string) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "addon.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Addon",
	}:
		if err := (&controllersaddon.AddonReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Addon"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_addon"],
			TypeName: "pagerduty_addon",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Addon")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "business.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Service",
	}:
		if err := (&controllersbusiness.ServiceReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Service"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_business_service"],
			TypeName: "pagerduty_business_service",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Service")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "business.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ServiceSubscriber",
	}:
		if err := (&controllersbusiness.ServiceSubscriberReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("ServiceSubscriber"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_business_service_subscriber"],
			TypeName: "pagerduty_business_service_subscriber",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "ServiceSubscriber")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "escalation.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Policy",
	}:
		if err := (&controllersescalation.PolicyReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Policy"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_escalation_policy"],
			TypeName: "pagerduty_escalation_policy",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Policy")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "event.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Rule",
	}:
		if err := (&controllersevent.RuleReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Rule"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_event_rule"],
			TypeName: "pagerduty_event_rule",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Rule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "extension.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Extension",
	}:
		if err := (&controllersextension.ExtensionReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Extension"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_extension"],
			TypeName: "pagerduty_extension",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Extension")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "extension.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Servicenow",
	}:
		if err := (&controllersextension.ServicenowReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Servicenow"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_extension_servicenow"],
			TypeName: "pagerduty_extension_servicenow",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Servicenow")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "maintenance.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Window",
	}:
		if err := (&controllersmaintenance.WindowReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Window"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_maintenance_window"],
			TypeName: "pagerduty_maintenance_window",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Window")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "response.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Play",
	}:
		if err := (&controllersresponse.PlayReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Play"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_response_play"],
			TypeName: "pagerduty_response_play",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Play")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ruleset.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Ruleset",
	}:
		if err := (&controllersruleset.RulesetReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Ruleset"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_ruleset"],
			TypeName: "pagerduty_ruleset",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Ruleset")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ruleset.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Rule",
	}:
		if err := (&controllersruleset.RuleReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Rule"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_ruleset_rule"],
			TypeName: "pagerduty_ruleset_rule",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Rule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "schedule.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Schedule",
	}:
		if err := (&controllersschedule.ScheduleReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Schedule"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_schedule"],
			TypeName: "pagerduty_schedule",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Schedule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "service.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Service",
	}:
		if err := (&controllersservice.ServiceReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Service"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_service"],
			TypeName: "pagerduty_service",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Service")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "service.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Dependency",
	}:
		if err := (&controllersservice.DependencyReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Dependency"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_service_dependency"],
			TypeName: "pagerduty_service_dependency",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Dependency")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "service.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "EventRule",
	}:
		if err := (&controllersservice.EventRuleReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("EventRule"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_service_event_rule"],
			TypeName: "pagerduty_service_event_rule",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "EventRule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "service.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Integration",
	}:
		if err := (&controllersservice.IntegrationReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Integration"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_service_integration"],
			TypeName: "pagerduty_service_integration",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Integration")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "slack.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Connection",
	}:
		if err := (&controllersslack.ConnectionReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Connection"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_slack_connection"],
			TypeName: "pagerduty_slack_connection",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Connection")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "tag.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Tag",
	}:
		if err := (&controllerstag.TagReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Tag"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_tag"],
			TypeName: "pagerduty_tag",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Tag")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "tag.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Assignment",
	}:
		if err := (&controllerstag.AssignmentReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Assignment"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_tag_assignment"],
			TypeName: "pagerduty_tag_assignment",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Assignment")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "team.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Team",
	}:
		if err := (&controllersteam.TeamReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Team"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_team"],
			TypeName: "pagerduty_team",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Team")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "team.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Membership",
	}:
		if err := (&controllersteam.MembershipReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Membership"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_team_membership"],
			TypeName: "pagerduty_team_membership",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Membership")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "User",
	}:
		if err := (&controllersuser.UserReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("User"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_user"],
			TypeName: "pagerduty_user",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "User")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ContactMethod",
	}:
		if err := (&controllersuser.ContactMethodReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("ContactMethod"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_user_contact_method"],
			TypeName: "pagerduty_user_contact_method",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "ContactMethod")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "NotificationRule",
	}:
		if err := (&controllersuser.NotificationRuleReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("NotificationRule"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_user_notification_rule"],
			TypeName: "pagerduty_user_notification_rule",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "NotificationRule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "webhook.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Subscription",
	}:
		if err := (&controllerswebhook.SubscriptionReconciler{
			Client:   mgr.GetClient(),
			Log:      ctrl.Log.WithName("controllers").WithName("Subscription"),
			Scheme:   mgr.GetScheme(),
			Gvk:      gvk,
			Provider: _provider,
			Resource: _provider.ResourcesMap["pagerduty_webhook_subscription"],
			TypeName: "pagerduty_webhook_subscription",
		}).SetupWithManager(ctx, mgr, auditor, restrictToNamespace); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "Subscription")
			return err
		}

	default:
		return fmt.Errorf("Invalid CRD")
	}

	return nil
}

func SetupWebhook(mgr manager.Manager, gvk schema.GroupVersionKind) error {
	switch gvk {
	case schema.GroupVersionKind{
		Group:   "addon.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Addon",
	}:
		if err := (&addonv1alpha1.Addon{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Addon")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "business.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Service",
	}:
		if err := (&businessv1alpha1.Service{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Service")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "business.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ServiceSubscriber",
	}:
		if err := (&businessv1alpha1.ServiceSubscriber{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "ServiceSubscriber")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "escalation.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Policy",
	}:
		if err := (&escalationv1alpha1.Policy{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Policy")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "event.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Rule",
	}:
		if err := (&eventv1alpha1.Rule{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Rule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "extension.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Extension",
	}:
		if err := (&extensionv1alpha1.Extension{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Extension")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "extension.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Servicenow",
	}:
		if err := (&extensionv1alpha1.Servicenow{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Servicenow")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "maintenance.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Window",
	}:
		if err := (&maintenancev1alpha1.Window{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Window")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "response.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Play",
	}:
		if err := (&responsev1alpha1.Play{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Play")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ruleset.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Ruleset",
	}:
		if err := (&rulesetv1alpha1.Ruleset{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Ruleset")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "ruleset.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Rule",
	}:
		if err := (&rulesetv1alpha1.Rule{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Rule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "schedule.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Schedule",
	}:
		if err := (&schedulev1alpha1.Schedule{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Schedule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "service.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Service",
	}:
		if err := (&servicev1alpha1.Service{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Service")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "service.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Dependency",
	}:
		if err := (&servicev1alpha1.Dependency{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Dependency")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "service.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "EventRule",
	}:
		if err := (&servicev1alpha1.EventRule{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "EventRule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "service.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Integration",
	}:
		if err := (&servicev1alpha1.Integration{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Integration")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "slack.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Connection",
	}:
		if err := (&slackv1alpha1.Connection{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Connection")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "tag.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Tag",
	}:
		if err := (&tagv1alpha1.Tag{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Tag")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "tag.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Assignment",
	}:
		if err := (&tagv1alpha1.Assignment{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Assignment")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "team.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Team",
	}:
		if err := (&teamv1alpha1.Team{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Team")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "team.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Membership",
	}:
		if err := (&teamv1alpha1.Membership{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Membership")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "User",
	}:
		if err := (&userv1alpha1.User{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "User")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "ContactMethod",
	}:
		if err := (&userv1alpha1.ContactMethod{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "ContactMethod")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "user.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "NotificationRule",
	}:
		if err := (&userv1alpha1.NotificationRule{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "NotificationRule")
			return err
		}
	case schema.GroupVersionKind{
		Group:   "webhook.pagerduty.kubeform.com",
		Version: "v1alpha1",
		Kind:    "Subscription",
	}:
		if err := (&webhookv1alpha1.Subscription{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "Subscription")
			return err
		}

	default:
		return fmt.Errorf("Invalid Webhook")
	}

	return nil
}
