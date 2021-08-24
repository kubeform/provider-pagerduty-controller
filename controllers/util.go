// Code generated by Kubeform. DO NOT EDIT.

package controllers

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/structs"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	tfplugin "github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/imdario/mergo"
	jsoniter "github.com/json-iterator/go"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/msgpack"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	kmapi "kmodules.xyz/client-go/api/v1"
	"kmodules.xyz/client-go/meta"
	base "kubeform.dev/apimachinery/api/v1alpha1"
	pagerduty "kubeform.dev/provider-pagerduty-api/api/provider"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	KFCFinalizer       = "pagerduty.kubeform.com/finalizer"
	UnknownIdValue     = "4856ec62-a372-11eb-bcbc-0242ac130002"
	UpdateNotSupported = "doesn't support update"
)

func StartProcess(rClient client.Client, provider *tfschema.Provider, ctx context.Context, res *tfschema.Resource, gv schema.GroupVersion, unstructuredObj *unstructured.Unstructured, tName string, jsonit jsoniter.API) error {
	err := initialUpdateStatus(rClient, ctx, gv, unstructuredObj, nil, true)
	if err != nil {
		return err
	}

	err = reconcile(rClient, provider, ctx, res, gv, unstructuredObj, tName, jsonit)
	if err != nil {
		err2 := initialUpdateStatus(rClient, ctx, gv, unstructuredObj, err, false)
		if err2 != nil {
			return err2
		}
		return err
	}

	return finalUpdateStatus(rClient, ctx, gv, unstructuredObj)
}

func reconcile(rClient client.Client, provider *tfschema.Provider, ctx context.Context, res *tfschema.Resource, gv schema.GroupVersion, unstructuredObj *unstructured.Unstructured, tName string, jsonit jsoniter.API) error {
	server := tfplugin.NewKubeformServer(provider)

	// Get RawSpec (including sensitive data)
	rawSpec, err := getSpecWithSensitiveData(gv, rClient, ctx, unstructuredObj, jsonit)
	if err != nil {
		return err
	}

	// Get RawStatus (including sensitive data)
	rawStatus, err := getStatusWithSensitiveData(gv, rClient, ctx, unstructuredObj, jsonit)
	if err != nil {
		return err
	}

	// Get object ID
	_, found, err := unstructured.NestedString(unstructuredObj.Object, "spec", "resource", "id")
	if err != nil {
		return err
	}

	// Set ProviderMeta
	err = setProviderMeta(rClient, provider, ctx, unstructuredObj, server, jsonit)
	if err != nil {
		return err
	}

	// validation check
	if rawSpec["id"] == nil {
		rawSpec["id"] = UnknownIdValue
	}
	rawSpecCty := HCL2ValueFromConfigValue(rawSpec)
	initialState, err := msgpack.Marshal(rawSpecCty, res.CoreConfigSchema().ImpliedType())
	if err != nil {
		return err
	}

	err = server.ValidateResourceTypeConfig(tName, initialState)
	if err != nil {
		return err
	}

	if hasFinalizer(unstructuredObj.GetFinalizers(), KFCFinalizer) {
		if unstructuredObj.GetDeletionTimestamp() != nil {
			err := updateStatus(rClient, ctx, unstructuredObj, status.TerminatingStatus)
			if err != nil {
				return err
			}
			// if not found then also delete
			if found {
				err = destroyTheObject(rawStatus, res, server, tName)
				if err != nil && !strings.Contains(err.Error(), "[404] Not found") {
					return err
				}
			}
			secretName, found, err := unstructured.NestedString(unstructuredObj.Object, "spec", "secretRef", "name")
			if err != nil {
				return err
			}
			if found {
				var secret corev1.Secret
				req := types.NamespacedName{
					Namespace: unstructuredObj.GetNamespace(),
					Name:      secretName,
				}
				if err := rClient.Get(ctx, req, &secret); err != nil {
					return err
				}
				delete(secret.Data, "state")
				if err := rClient.Update(ctx, &secret); err != nil {
					return err
				}
			}
			return removeFinalizer(ctx, rClient, unstructuredObj, KFCFinalizer)
		}
	} else {
		err := addFinalizer(ctx, rClient, unstructuredObj, KFCFinalizer)
		if err != nil {
			return err
		}
	}

	if !found {
		err := updateStatus(rClient, ctx, unstructuredObj, status.InProgressStatus)
		if err != nil {
			return err
		}
		newStateVal, intrfc, err := createTheObject(rawSpec, res, server, tName)
		if err != nil {
			return err
		}
		err = updateStatus(rClient, ctx, unstructuredObj, status.CurrentStatus)
		if err != nil {
			return err
		}
		err = updateStateField(rClient, ctx, intrfc.Raw, gv, unstructuredObj, jsonit)
		if err != nil {
			return err
		}

		// set the id value in unstructuredObj object
		err = unstructured.SetNestedField(unstructuredObj.Object, newStateVal.GetAttr("id").AsString(), "spec", "resource", "id")
		if err != nil {
			return err
		}

		// apply the update of the object
		if err = rClient.Update(ctx, unstructuredObj); err != nil {
			return err
		}
		return nil
	}

	combineRaw, err := getCombineRawAndDeepCopyRawStatus(rawStatus, rawSpec)
	if err != nil {
		return err
	}

	changed, err := hasResourceChanged(combineRaw, rawStatus, res)
	if err != nil {
		return err
	}
	if !changed {
		return nil
	}

	requireNew, priorState, proposedState, plannedState, plannedPrivate, err := checkRequireNewOrNot(combineRaw, rawStatus, res, server, tName)
	if err != nil {
		return err
	}

	if requireNew { // Resources is needed to be destroyed because one of the field needs the resource to be replaced
		updatePolicy, _, err := unstructured.NestedFieldNoCopy(unstructuredObj.Object, "spec", "updatePolicy")
		if err != nil {
			return err
		}
		if updatePolicy == string(base.UpdatePolicyDoNotDestroy) {
			return fmt.Errorf("updatePolicy is set to `DoNotDestroy`, can't destroy the object to create a new one")
		}

		err = updateStatus(rClient, ctx, unstructuredObj, status.TerminatingStatus)
		if err != nil {
			return err
		}
		err = destroyTheObject(rawStatus, res, server, tName)
		if err != nil {
			return err
		}

		err = updateStatus(rClient, ctx, unstructuredObj, status.InProgressStatus)
		if err != nil {
			return err
		}
		newStateVal, intrfc, err := createTheObject(rawSpec, res, server, tName)
		if err != nil {
			return err
		}
		err = updateStatus(rClient, ctx, unstructuredObj, status.CurrentStatus)
		if err != nil {
			return err
		}
		err = updateStateField(rClient, ctx, intrfc.Raw, gv, unstructuredObj, jsonit)
		if err != nil {
			return err
		}

		// set the id value in unstructuredObj object
		err = unstructured.SetNestedField(unstructuredObj.Object, newStateVal.GetAttr("id").AsString(), "spec", "resource", "id")
		if err != nil {
			return err
		}

		// apply the update of the object
		if err = rClient.Update(ctx, unstructuredObj); err != nil {
			return err
		}
		return nil
	}

	newStateVal, updated, err := updateTheObject(priorState, proposedState, plannedState, plannedPrivate, server, res, tName)
	if err != nil {
		return err
	}

	if updated {
		//set the id value in unstructuredObj object
		err = unstructured.SetNestedField(unstructuredObj.Object, newStateVal.GetAttr("id").AsString(), "spec", "resource", "id")
		if err != nil {
			return err
		}

		// apply the update of the object
		if err = rClient.Update(ctx, unstructuredObj); err != nil {
			return err
		}

		intrfc := terraform.NewResourceConfigShimmed(newStateVal, res.CoreConfigSchema())

		err = updateStateField(rClient, ctx, intrfc.Raw, gv, unstructuredObj, jsonit)
		if err != nil {
			return err
		}
	}

	return nil
}

func initialUpdateStatus(rClient client.Client, ctx context.Context, gv schema.GroupVersion, obj *unstructured.Unstructured, er error, flag bool) error {
	objGen, _, err := unstructured.NestedInt64(obj.Object, "metadata", "generation")
	if err != nil {
		return err
	}

	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return err
	}

	typedStruct := structs.New(typedObj)
	conditionsVal := reflect.ValueOf(typedStruct.Field("Status").Field("Conditions").Value())
	conditions := conditionsVal.Interface().([]kmapi.Condition)
	if kmapi.HasCondition(conditions, "Stalled") {
		return nil
	}

	phase := status.InProgressStatus

	if flag {
		conditions = kmapi.SetCondition(conditions, kmapi.NewCondition("Reconciling", "Kubeform is currently reconciling "+obj.GetKind()+" resource", objGen))
	} else {
		conditions = kmapi.SetCondition(conditions, kmapi.NewCondition("Stalled", er.Error(), objGen))
		phase = status.FailedStatus
	}

	err = setNestedFieldNoCopy(obj.Object, conditions, "status", "conditions")
	if err != nil {
		return err
	}
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}

	return updateStatus(rClient, ctx, obj, phase)
}

func finalUpdateStatus(rClient client.Client, ctx context.Context, gv schema.GroupVersion, obj *unstructured.Unstructured) error {
	var newCondi []kmapi.Condition
	err := setNestedFieldNoCopy(obj.Object, newCondi, "status", "conditions")
	if err != nil {
		return err
	}
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}
	err = updateStatus(rClient, ctx, obj, status.CurrentStatus)
	if err != nil {
		return err
	}
	return nil
}

func updateStatus(rClient client.Client, ctx context.Context, obj *unstructured.Unstructured, phase status.Status) error {
	if phase == status.CurrentStatus {
		obsGen, _, err := unstructured.NestedInt64(obj.Object, "metadata", "generation")
		if err != nil {
			return err
		}
		err = unstructured.SetNestedField(obj.Object, obsGen, "status", "observedGeneration")
		if err != nil {
			return err
		}
	}

	err := setNestedFieldNoCopy(obj.Object, phase, "status", "phase")
	if err != nil {
		return err
	}

	// apply the status update of the object
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}
	return nil
}

func setNestedFieldNoCopy(obj map[string]interface{}, value interface{}, fields ...string) error {
	m := obj

	for i, field := range fields[:len(fields)-1] {
		if val, ok := m[field]; ok {
			if valMap, ok := val.(map[string]interface{}); ok {
				m = valMap
			} else {
				return fmt.Errorf("value cannot be set because %v is not a map[string]interface{}", jsonPath(fields[:i+1]))
			}
		} else {
			newVal := make(map[string]interface{})
			m[field] = newVal
			m = newVal
		}
	}
	m[fields[len(fields)-1]] = value
	return nil
}

func jsonPath(fields []string) string {
	return "." + strings.Join(fields, ".")
}

func getCombineRawAndDeepCopyRawStatus(rawStatus map[string]interface{}, rawSpec map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(rawSpec)
	if err != nil {
		return nil, err
	}

	var copyrawSpec map[string]interface{}
	err = dec.Decode(&copyrawSpec)
	if err != nil {
		return nil, err
	}
	if err := mergo.Merge(&copyrawSpec, rawStatus); err != nil {
		return nil, err
	}

	return copyrawSpec, nil
}

func updateTheObject(priorState []byte, proposedState, plannedState, plannedPrivate []byte, server *tfplugin.KubeformServer, res *tfschema.Resource, tName string) (cty.Value, bool, error) {
	newState, err := server.ApplyResourceChange(tName, priorState, plannedState, proposedState, plannedPrivate)
	if err != nil {
		return cty.Value{}, false, err
	}

	schma := res.CoreConfigSchema()
	newStateVal, err := msgpack.Unmarshal(newState, schma.ImpliedType())
	if err != nil {
		return cty.Value{}, false, err
	}

	return newStateVal, true, nil
}

func hasResourceChanged(combineRaw map[string]interface{}, copyrawStatus map[string]interface{}, res *tfschema.Resource) (bool, error) {
	stateVal := HCL2ValueFromConfigValue(copyrawStatus)
	proposedPlanVal := HCL2ValueFromConfigValue(combineRaw)

	diff, err := tfschema.DiffFromValues(stateVal, proposedPlanVal, stripResourceModifiers(res))
	if err != nil {
		return false, err
	}

	return diff != nil, nil
}

func checkRequireNewOrNot(combineRaw map[string]interface{}, copyrawStatus map[string]interface{}, res *tfschema.Resource, server *tfplugin.KubeformServer, tName string) (bool, []byte, []byte, []byte, []byte, error) {
	stateVal := HCL2ValueFromConfigValue(copyrawStatus)
	proposedPlanVal := HCL2ValueFromConfigValue(combineRaw)

	schma := res.CoreConfigSchema()
	priorState, err := msgpack.Marshal(stateVal, schma.ImpliedType())
	if err != nil {
		return false, nil, nil, nil, nil, err
	}

	proposedState, err := msgpack.Marshal(proposedPlanVal, schma.ImpliedType())
	if err != nil {
		return false, nil, nil, nil, nil, err
	}

	plannedState, plannedPrivate, requireNew, err := server.PlanResourceChange(tName, priorState, proposedState)
	if err != nil {
		return false, nil, nil, nil, nil, err
	}

	return requireNew, priorState, proposedState, plannedState, plannedPrivate, nil
}

func createTheObject(rawSpec map[string]interface{}, res *tfschema.Resource, server *tfplugin.KubeformServer, tName string) (cty.Value, *terraform.ResourceConfig, error) {
	rawSpec["id"] = UnknownIdValue
	stateVal := HCL2ValueFromConfigValue(rawSpec)

	schma := res.CoreConfigSchema()
	priorState, err := msgpack.Marshal(cty.NullVal(schma.ImpliedType()), schma.ImpliedType())
	if err != nil {
		return cty.Value{}, nil, err
	}
	proposedState, err := msgpack.Marshal(stateVal, schma.ImpliedType())
	if err != nil {
		return cty.Value{}, nil, err
	}

	plannedState, plannedPrivate, _, err := server.PlanResourceChange(tName, priorState, proposedState)
	if err != nil {
		return cty.Value{}, nil, err
	}

	newState, err := server.ApplyResourceChange(tName, priorState, plannedState, proposedState, plannedPrivate)
	if err != nil {
		return cty.Value{}, nil, err
	}

	newStateVal, err := msgpack.Unmarshal(newState, schma.ImpliedType())
	if err != nil {
		return cty.Value{}, nil, err
	}
	intrfc := terraform.NewResourceConfigShimmed(newStateVal, res.CoreConfigSchema())

	return newStateVal, intrfc, nil
}

func destroyTheObject(rawStatus map[string]interface{}, res *tfschema.Resource, server *tfplugin.KubeformServer, tName string) error {
	stateVal := HCL2ValueFromConfigValue(rawStatus)
	schma := res.CoreConfigSchema()
	priorState, err := msgpack.Marshal(stateVal, schma.ImpliedType())
	if err != nil {
		return err
	}

	proposedState, err := msgpack.Marshal(cty.NullVal(schma.ImpliedType()), schma.ImpliedType())
	if err != nil {
		return err
	}

	plannedState, plannedPrivate, _, err := server.PlanResourceChange(tName, priorState, proposedState)
	if err != nil {
		return err
	}

	_, err = server.ApplyResourceChange(tName, priorState, plannedState, proposedState, plannedPrivate)

	return err
}

func GetJSONItr(typeEncoder map[string]jsoniter.ValEncoder, typeDecoder map[string]jsoniter.ValDecoder) jsoniter.API {
	return jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeDecoders:           typeDecoder,
		TypeEncoders:           typeEncoder,
	}.Froze()
}

func getStatusWithSensitiveData(gv schema.GroupVersion, rClient client.Client, ctx context.Context, obj *unstructured.Unstructured, jsonit jsoniter.API) (map[string]interface{}, error) {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return nil, err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return nil, err
	}

	typedStruct := structs.New(typedObj)
	status := reflect.ValueOf(typedStruct.Field("Spec").Field("State").Value())
	statusType := reflect.TypeOf(typedStruct.Field("Spec").Field("State").Value())
	statusValue := reflect.New(statusType)
	statusValue.Elem().Set(status)

	secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
	if err != nil {
		return nil, err
	}

	secretData := make(map[string]interface{})
	if secretRef != nil {
		secretName := typedStruct.Field("Spec").Field("SecretRef").Field("Name").Value()

		if secretName != nil {
			var secret corev1.Secret
			req := types.NamespacedName{
				Namespace: obj.GetNamespace(),
				Name:      secretName.(string),
			}
			if err := rClient.Get(ctx, req, &secret); err != nil {
				return nil, err
			}

			if _, ok := secret.Data["state"]; ok {
				err = json.Unmarshal(secret.Data["state"], &secretData)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	str, err := jsonit.Marshal(statusValue.Interface())
	if err != nil {
		return nil, err
	}
	rawStatus := make(map[string]interface{})
	err = json.Unmarshal(str, &rawStatus)
	if err != nil {
		return nil, err
	}

	if err := mergo.Merge(&rawStatus, secretData); err != nil {
		return nil, err
	}

	return rawStatus, nil
}

func getSpecWithSensitiveData(gv schema.GroupVersion, rClient client.Client, ctx context.Context, obj *unstructured.Unstructured, jsonit jsoniter.API) (map[string]interface{}, error) {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return nil, err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return nil, err
	}

	typedStruct := structs.New(typedObj)
	spec := reflect.ValueOf(typedStruct.Field("Spec").Field("Resource").Value())
	specType := reflect.TypeOf(typedStruct.Field("Spec").Field("Resource").Value())
	specValue := reflect.New(specType)
	specValue.Elem().Set(spec)

	secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
	if err != nil {
		return nil, err
	}

	secretData := make(map[string]interface{})
	if secretRef != nil {
		secretName := typedStruct.Field("Spec").Field("SecretRef").Field("Name").Value()

		if secretName != nil {
			var secret corev1.Secret
			req := types.NamespacedName{
				Namespace: obj.GetNamespace(),
				Name:      secretName.(string),
			}
			if err := rClient.Get(ctx, req, &secret); err != nil {
				return nil, err
			}

			if _, ok := secret.Data["resource"]; ok {
				err = json.Unmarshal(secret.Data["resource"], &secretData)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	str, err := jsonit.Marshal(specValue.Interface())
	if err != nil {
		return nil, err
	}
	rawSpec := make(map[string]interface{})
	err = json.Unmarshal(str, &rawSpec)
	if err != nil {
		return nil, err
	}

	if err := mergo.Merge(&rawSpec, secretData); err != nil {
		return nil, err
	}

	return rawSpec, nil
}

func getProviderSecretData(rClient client.Client, ctx context.Context, obj *unstructured.Unstructured) (map[string][]byte, error) {
	providerRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "providerRef", "name")
	if err != nil {
		return nil, err
	}
	configData := make(map[string][]byte)
	if providerRef != nil {
		var secret corev1.Secret
		req := types.NamespacedName{
			Namespace: obj.GetNamespace(),
			Name:      providerRef.(string),
		}
		if err := rClient.Get(ctx, req, &secret); err != nil {
			return nil, err
		}

		configData = secret.Data
	}
	return configData, nil
}

func hasFinalizer(finalizers []string, finalizer string) bool {
	for _, f := range finalizers {
		if f == finalizer {
			return true
		}
	}
	return false
}

func removeFinalizer(ctx context.Context, rClient client.Client, u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for i, v := range finalizers {
		if v == finalizer {
			finalizers = append(finalizers[:i], finalizers[i+1:]...)
			break
		}
	}
	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}

	err = rClient.Update(ctx, u)
	return err
}

func addFinalizer(ctx context.Context, rClient client.Client, u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for _, v := range finalizers {
		if v == finalizer {
			return nil
		}
	}
	finalizers = append(finalizers, finalizer)
	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}
	err = rClient.Update(ctx, u)
	return err
}

func updateStateField(rClient client.Client, ctx context.Context, intrfc map[string]interface{}, gv schema.GroupVersion, obj *unstructured.Unstructured, jsonit jsoniter.API) error {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return err
	}

	var raw []byte
	jsonByte, err := json.Marshal(intrfc)
	if err != nil {
		return err
	}

	raw = append(raw, []byte(`{"spec":{ "resource":`)...)
	raw = append(raw, jsonByte...)
	raw = append(raw, []byte(`}}`)...)

	err = jsonit.Unmarshal(raw, &typedObj)
	if err != nil {
		return err
	}

	s := structs.New(typedObj)
	hasAnySensitiveField := false
	secretData, _, err := processSensitiveFields(reflect.TypeOf(s.Field("Spec").Field("Resource").Value()), reflect.ValueOf(s.Field("Spec").Field("Resource").Value()), &hasAnySensitiveField)
	if err != nil {
		return err
	}

	if hasAnySensitiveField {
		var secretName string

		secretRef, _, err := unstructured.NestedFieldNoCopy(obj.Object, "spec", "secretRef")
		if err != nil {
			return err
		}

		if secretRef != nil {
			secretName = s.Field("Spec").Field("SecretRef").Field("Name").Value().(string)
		} else {
			secretName = obj.GetName() + "-" + "pagerduty" + "-" + "sensitive"
		}

		var secret corev1.Secret
		req := types.NamespacedName{
			Namespace: obj.GetNamespace(),
			Name:      secretName,
		}
		tr := true
		if err := rClient.Get(ctx, req, &secret); err != nil {
			if errors.ReasonForError(err) == metav1.StatusReasonNotFound {
				err = rClient.Create(ctx, &corev1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      secretName,
						Namespace: obj.GetNamespace(),
						OwnerReferences: []metav1.OwnerReference{
							{
								APIVersion: obj.GetAPIVersion(),
								Kind:       obj.GetKind(),
								Name:       obj.GetName(),
								Controller: &tr,
								UID:        obj.GetUID(),
							},
						},
					},
				})
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if err := rClient.Get(ctx, req, &secret); err != nil {
			return err
		}
		if secret.Data == nil {
			secret.Data = make(map[string][]byte)
		}

		secretByte, err := json.Marshal(secretData)
		if err != nil {
			return err
		}
		secret.Data["state"] = secretByte

		// apply the update of the object
		if err = rClient.Update(ctx, &secret); err != nil {
			return err
		}
	}

	output := s.Field("Spec").Field("Resource").Value()
	specByte, err := json.Marshal(output)
	if err != nil {
		return err
	}

	var specMap map[string]interface{}
	err = json.Unmarshal(specByte, &specMap)
	if err != nil {
		return err
	}

	err = unstructured.SetNestedField(obj.Object, specMap, "spec", "state")
	if err != nil {
		return err
	}

	if err = rClient.Update(ctx, obj); err != nil {
		return err
	}
	return nil
}

func processSensitiveFields(r reflect.Type, v reflect.Value, hasAnySensitiveField *bool) (map[string]interface{}, bool, error) {
	var err error
	out := make(map[string]interface{})
	n := r.NumField()
	tr := false
	for i := 0; i < n; i++ {
		field := r.Field(i)
		value := v.Field(i)
		tag := strings.ReplaceAll(field.Tag.Get("tf"), ",omitempty", "")
		if tag == "-" {
			continue
		}

		if field.Tag.Get("sensitive") == "true" && value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.String && value.Elem().String() != "" {
			out[tag] = value.Elem().String()
			*hasAnySensitiveField = true
			tr = true
		} else if field.Tag.Get("sensitive") == "true" && value.Kind() == reflect.Map && value.Interface().(map[string]string) != nil && len(value.Interface().(map[string]string)) != 0 {
			secretJson, err := json.Marshal(value.Interface())
			*hasAnySensitiveField = true
			tr = true
			if err != nil {
				return nil, tr, err
			} else {
				out[tag] = string(secretJson)
			}
		}
		if value.Kind() == reflect.Struct {
			val, isItSen, err := processSensitiveFields(value.Type(), value, hasAnySensitiveField)
			if err != nil {
				return nil, tr, err
			}
			if isItSen {
				out[tag] = val
				tr = true
			}
		}
		if value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct {
			val, isItSen, err := processSensitiveFields(value.Elem().Type(), value.Elem(), hasAnySensitiveField)
			if err != nil {
				return nil, tr, err
			}
			if isItSen {
				out[tag] = val
				tr = true
			}
		}

		if value.Kind() == reflect.Slice {
			n := value.Len()
			tempMap := make([]map[string]interface{}, n)
			tempBool := false
			var isItSen bool
			for i := 0; i < n; i++ {
				if value.Index(i).Kind() == reflect.Struct {
					tempMap[i], isItSen, err = processSensitiveFields(value.Index(i).Type(), value.Index(i), hasAnySensitiveField)
					if err != nil {
						return nil, false, err
					}
					if isItSen {
						tempBool = true
						tr = true
					}
				}
			}
			if tempBool {
				out[tag] = tempMap
				tr = true
			}
		}
	}
	return out, tr, nil
}

func setProviderMeta(rClient client.Client, provider *tfschema.Provider, ctx context.Context, unstructuredObj *unstructured.Unstructured, server *tfplugin.KubeformServer, jsonit jsoniter.API) error {
	jsonit = GetJSONItr(pagerduty.GetEncoder(), pagerduty.GetDecoder())
	providerSecretData, err := getProviderSecretData(rClient, ctx, unstructuredObj)
	if err != nil {
		return err
	}
	providerSchema, err := provider.GetSchema(&terraform.ProviderSchemaRequest{})
	if err != nil {
		return err
	}

	if providerSchema.Provider == nil {
		return fmt.Errorf("missing provider schema")
	}

	providerSpec := &pagerduty.PagerdutySpec{}
	err = jsonit.Unmarshal(providerSecretData["provider"], providerSpec)
	if err != nil {
		return err
	}

	providerDataByte, err := jsonit.Marshal(providerSpec)
	if err != nil {
		return err
	}

	mapData := make(map[string]interface{})
	err = jsonit.Unmarshal(providerDataByte, &mapData)
	if err != nil {
		return err
	}

	configRaw := HCL2ValueFromConfigValue(mapData)
	configPlan, err := msgpack.Marshal(configRaw, providerSchema.Provider.ImpliedType())
	if err != nil {
		return err
	}

	prepareConfigResp, err := server.PrepareProviderConfig(configPlan)
	if err != nil {
		return err
	}

	return server.Configure(prepareConfigResp)
}

func HCL2ValueFromConfigValue(v interface{}) cty.Value {
	if v == nil || v == UnknownIdValue {
		return cty.NullVal(cty.DynamicPseudoType)
	}

	switch tv := v.(type) {
	case bool:
		return cty.BoolVal(tv)
	case string:
		return cty.StringVal(tv)
	case int:
		return cty.NumberIntVal(int64(tv))
	case float64:
		return cty.NumberFloatVal(tv)
	case []interface{}:
		vals := make([]cty.Value, len(tv))
		for i, ev := range tv {
			vals[i] = HCL2ValueFromConfigValue(ev)
		}
		return cty.TupleVal(vals)
	case map[string]interface{}:
		vals := map[string]cty.Value{}
		for k, ev := range tv {
			vals[k] = HCL2ValueFromConfigValue(ev)
		}
		return cty.ObjectVal(vals)
	case map[string]string:
		vals := map[string]cty.Value{}
		for k, ev := range tv {
			vals[k] = HCL2ValueFromConfigValue(ev)
		}
		return cty.ObjectVal(vals)
	default:
		// HCL/HIL should never generate anything that isn't caught by
		// the above, so if we get here something has gone very wrong.
		panic(fmt.Errorf("can't convert %#v to cty.Value", v))
	}
}

// stripResourceModifiers takes a *schema.Resource and returns a deep copy with all
// StateFuncs and CustomizeDiffs removed. This will be used during apply to
// create a diff from a planned state where the diff modifications have already
// been applied.
func stripResourceModifiers(r *tfschema.Resource) *tfschema.Resource {
	if r == nil {
		return nil
	}
	// start with a shallow copy
	newResource := new(tfschema.Resource)
	*newResource = *r

	newResource.CustomizeDiff = nil
	newResource.Schema = map[string]*tfschema.Schema{}

	for k, s := range r.Schema {
		newResource.Schema[k] = stripSchema(s)
	}

	return newResource
}

func stripSchema(s *tfschema.Schema) *tfschema.Schema {
	if s == nil {
		return nil
	}
	// start with a shallow copy
	newSchema := new(tfschema.Schema)
	*newSchema = *s

	newSchema.StateFunc = nil

	switch e := newSchema.Elem.(type) {
	case *tfschema.Schema:
		newSchema.Elem = stripSchema(e)
	case *tfschema.Resource:
		newSchema.Elem = stripResourceModifiers(e)
	}

	return newSchema
}
