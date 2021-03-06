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
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParameters{}).Type1()):                ServiceSpecAlertGroupingParametersCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParametersConfig{}).Type1()):          ServiceSpecAlertGroupingParametersConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRule{}).Type1()):                    ServiceSpecIncidentUrgencyRuleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleDuringSupportHours{}).Type1()):  ServiceSpecIncidentUrgencyRuleDuringSupportHoursCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}).Type1()): ServiceSpecIncidentUrgencyRuleOutsideSupportHoursCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecSupportHours{}).Type1()):                           ServiceSpecSupportHoursCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DependencySpecDependency{}).Type1()):                          DependencySpecDependencyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecActions{}).Type1()):                              EventRuleSpecActionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecConditions{}).Type1()):                           EventRuleSpecConditionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecTimeFrame{}).Type1()):                            EventRuleSpecTimeFrameCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParameters{}).Type1()):                ServiceSpecAlertGroupingParametersCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParametersConfig{}).Type1()):          ServiceSpecAlertGroupingParametersConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRule{}).Type1()):                    ServiceSpecIncidentUrgencyRuleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleDuringSupportHours{}).Type1()):  ServiceSpecIncidentUrgencyRuleDuringSupportHoursCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}).Type1()): ServiceSpecIncidentUrgencyRuleOutsideSupportHoursCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecSupportHours{}).Type1()):                           ServiceSpecSupportHoursCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DependencySpecDependency{}).Type1()):                          DependencySpecDependencyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecActions{}).Type1()):                              EventRuleSpecActionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecConditions{}).Type1()):                           EventRuleSpecConditionsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecTimeFrame{}).Type1()):                            EventRuleSpecTimeFrameCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type ServiceSpecAlertGroupingParametersCodec struct {
}

func (ServiceSpecAlertGroupingParametersCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecAlertGroupingParameters)(ptr) == nil
}

func (ServiceSpecAlertGroupingParametersCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecAlertGroupingParameters)(ptr)
	var objs []ServiceSpecAlertGroupingParameters
	if obj != nil {
		objs = []ServiceSpecAlertGroupingParameters{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParameters{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecAlertGroupingParametersCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecAlertGroupingParameters)(ptr) = ServiceSpecAlertGroupingParameters{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecAlertGroupingParameters

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParameters{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecAlertGroupingParameters)(ptr) = objs[0]
			} else {
				*(*ServiceSpecAlertGroupingParameters)(ptr) = ServiceSpecAlertGroupingParameters{}
			}
		} else {
			*(*ServiceSpecAlertGroupingParameters)(ptr) = ServiceSpecAlertGroupingParameters{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServiceSpecAlertGroupingParameters

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParameters{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServiceSpecAlertGroupingParameters)(ptr) = obj
		} else {
			*(*ServiceSpecAlertGroupingParameters)(ptr) = ServiceSpecAlertGroupingParameters{}
		}
	default:
		iter.ReportError("decode ServiceSpecAlertGroupingParameters", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecAlertGroupingParametersConfigCodec struct {
}

func (ServiceSpecAlertGroupingParametersConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecAlertGroupingParametersConfig)(ptr) == nil
}

func (ServiceSpecAlertGroupingParametersConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecAlertGroupingParametersConfig)(ptr)
	var objs []ServiceSpecAlertGroupingParametersConfig
	if obj != nil {
		objs = []ServiceSpecAlertGroupingParametersConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParametersConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecAlertGroupingParametersConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecAlertGroupingParametersConfig)(ptr) = ServiceSpecAlertGroupingParametersConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecAlertGroupingParametersConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParametersConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecAlertGroupingParametersConfig)(ptr) = objs[0]
			} else {
				*(*ServiceSpecAlertGroupingParametersConfig)(ptr) = ServiceSpecAlertGroupingParametersConfig{}
			}
		} else {
			*(*ServiceSpecAlertGroupingParametersConfig)(ptr) = ServiceSpecAlertGroupingParametersConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServiceSpecAlertGroupingParametersConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecAlertGroupingParametersConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServiceSpecAlertGroupingParametersConfig)(ptr) = obj
		} else {
			*(*ServiceSpecAlertGroupingParametersConfig)(ptr) = ServiceSpecAlertGroupingParametersConfig{}
		}
	default:
		iter.ReportError("decode ServiceSpecAlertGroupingParametersConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecIncidentUrgencyRuleCodec struct {
}

func (ServiceSpecIncidentUrgencyRuleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecIncidentUrgencyRule)(ptr) == nil
}

func (ServiceSpecIncidentUrgencyRuleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecIncidentUrgencyRule)(ptr)
	var objs []ServiceSpecIncidentUrgencyRule
	if obj != nil {
		objs = []ServiceSpecIncidentUrgencyRule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecIncidentUrgencyRuleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecIncidentUrgencyRule)(ptr) = ServiceSpecIncidentUrgencyRule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecIncidentUrgencyRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecIncidentUrgencyRule)(ptr) = objs[0]
			} else {
				*(*ServiceSpecIncidentUrgencyRule)(ptr) = ServiceSpecIncidentUrgencyRule{}
			}
		} else {
			*(*ServiceSpecIncidentUrgencyRule)(ptr) = ServiceSpecIncidentUrgencyRule{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServiceSpecIncidentUrgencyRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServiceSpecIncidentUrgencyRule)(ptr) = obj
		} else {
			*(*ServiceSpecIncidentUrgencyRule)(ptr) = ServiceSpecIncidentUrgencyRule{}
		}
	default:
		iter.ReportError("decode ServiceSpecIncidentUrgencyRule", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecIncidentUrgencyRuleDuringSupportHoursCodec struct {
}

func (ServiceSpecIncidentUrgencyRuleDuringSupportHoursCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecIncidentUrgencyRuleDuringSupportHours)(ptr) == nil
}

func (ServiceSpecIncidentUrgencyRuleDuringSupportHoursCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecIncidentUrgencyRuleDuringSupportHours)(ptr)
	var objs []ServiceSpecIncidentUrgencyRuleDuringSupportHours
	if obj != nil {
		objs = []ServiceSpecIncidentUrgencyRuleDuringSupportHours{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleDuringSupportHours{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecIncidentUrgencyRuleDuringSupportHoursCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecIncidentUrgencyRuleDuringSupportHours)(ptr) = ServiceSpecIncidentUrgencyRuleDuringSupportHours{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecIncidentUrgencyRuleDuringSupportHours

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleDuringSupportHours{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecIncidentUrgencyRuleDuringSupportHours)(ptr) = objs[0]
			} else {
				*(*ServiceSpecIncidentUrgencyRuleDuringSupportHours)(ptr) = ServiceSpecIncidentUrgencyRuleDuringSupportHours{}
			}
		} else {
			*(*ServiceSpecIncidentUrgencyRuleDuringSupportHours)(ptr) = ServiceSpecIncidentUrgencyRuleDuringSupportHours{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServiceSpecIncidentUrgencyRuleDuringSupportHours

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleDuringSupportHours{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServiceSpecIncidentUrgencyRuleDuringSupportHours)(ptr) = obj
		} else {
			*(*ServiceSpecIncidentUrgencyRuleDuringSupportHours)(ptr) = ServiceSpecIncidentUrgencyRuleDuringSupportHours{}
		}
	default:
		iter.ReportError("decode ServiceSpecIncidentUrgencyRuleDuringSupportHours", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecIncidentUrgencyRuleOutsideSupportHoursCodec struct {
}

func (ServiceSpecIncidentUrgencyRuleOutsideSupportHoursCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecIncidentUrgencyRuleOutsideSupportHours)(ptr) == nil
}

func (ServiceSpecIncidentUrgencyRuleOutsideSupportHoursCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecIncidentUrgencyRuleOutsideSupportHours)(ptr)
	var objs []ServiceSpecIncidentUrgencyRuleOutsideSupportHours
	if obj != nil {
		objs = []ServiceSpecIncidentUrgencyRuleOutsideSupportHours{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecIncidentUrgencyRuleOutsideSupportHoursCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecIncidentUrgencyRuleOutsideSupportHours)(ptr) = ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecIncidentUrgencyRuleOutsideSupportHours

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecIncidentUrgencyRuleOutsideSupportHours)(ptr) = objs[0]
			} else {
				*(*ServiceSpecIncidentUrgencyRuleOutsideSupportHours)(ptr) = ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}
			}
		} else {
			*(*ServiceSpecIncidentUrgencyRuleOutsideSupportHours)(ptr) = ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServiceSpecIncidentUrgencyRuleOutsideSupportHours

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServiceSpecIncidentUrgencyRuleOutsideSupportHours)(ptr) = obj
		} else {
			*(*ServiceSpecIncidentUrgencyRuleOutsideSupportHours)(ptr) = ServiceSpecIncidentUrgencyRuleOutsideSupportHours{}
		}
	default:
		iter.ReportError("decode ServiceSpecIncidentUrgencyRuleOutsideSupportHours", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ServiceSpecSupportHoursCodec struct {
}

func (ServiceSpecSupportHoursCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ServiceSpecSupportHours)(ptr) == nil
}

func (ServiceSpecSupportHoursCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ServiceSpecSupportHours)(ptr)
	var objs []ServiceSpecSupportHours
	if obj != nil {
		objs = []ServiceSpecSupportHours{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecSupportHours{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ServiceSpecSupportHoursCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ServiceSpecSupportHours)(ptr) = ServiceSpecSupportHours{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ServiceSpecSupportHours

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecSupportHours{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ServiceSpecSupportHours)(ptr) = objs[0]
			} else {
				*(*ServiceSpecSupportHours)(ptr) = ServiceSpecSupportHours{}
			}
		} else {
			*(*ServiceSpecSupportHours)(ptr) = ServiceSpecSupportHours{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ServiceSpecSupportHours

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ServiceSpecSupportHours{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ServiceSpecSupportHours)(ptr) = obj
		} else {
			*(*ServiceSpecSupportHours)(ptr) = ServiceSpecSupportHours{}
		}
	default:
		iter.ReportError("decode ServiceSpecSupportHours", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type DependencySpecDependencyCodec struct {
}

func (DependencySpecDependencyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DependencySpecDependency)(ptr) == nil
}

func (DependencySpecDependencyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DependencySpecDependency)(ptr)
	var objs []DependencySpecDependency
	if obj != nil {
		objs = []DependencySpecDependency{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DependencySpecDependency{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DependencySpecDependencyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DependencySpecDependency)(ptr) = DependencySpecDependency{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DependencySpecDependency

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DependencySpecDependency{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DependencySpecDependency)(ptr) = objs[0]
			} else {
				*(*DependencySpecDependency)(ptr) = DependencySpecDependency{}
			}
		} else {
			*(*DependencySpecDependency)(ptr) = DependencySpecDependency{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DependencySpecDependency

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DependencySpecDependency{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DependencySpecDependency)(ptr) = obj
		} else {
			*(*DependencySpecDependency)(ptr) = DependencySpecDependency{}
		}
	default:
		iter.ReportError("decode DependencySpecDependency", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EventRuleSpecActionsCodec struct {
}

func (EventRuleSpecActionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EventRuleSpecActions)(ptr) == nil
}

func (EventRuleSpecActionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EventRuleSpecActions)(ptr)
	var objs []EventRuleSpecActions
	if obj != nil {
		objs = []EventRuleSpecActions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecActions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EventRuleSpecActionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EventRuleSpecActions)(ptr) = EventRuleSpecActions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EventRuleSpecActions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecActions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EventRuleSpecActions)(ptr) = objs[0]
			} else {
				*(*EventRuleSpecActions)(ptr) = EventRuleSpecActions{}
			}
		} else {
			*(*EventRuleSpecActions)(ptr) = EventRuleSpecActions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EventRuleSpecActions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecActions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EventRuleSpecActions)(ptr) = obj
		} else {
			*(*EventRuleSpecActions)(ptr) = EventRuleSpecActions{}
		}
	default:
		iter.ReportError("decode EventRuleSpecActions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EventRuleSpecConditionsCodec struct {
}

func (EventRuleSpecConditionsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EventRuleSpecConditions)(ptr) == nil
}

func (EventRuleSpecConditionsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EventRuleSpecConditions)(ptr)
	var objs []EventRuleSpecConditions
	if obj != nil {
		objs = []EventRuleSpecConditions{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecConditions{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EventRuleSpecConditionsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EventRuleSpecConditions)(ptr) = EventRuleSpecConditions{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EventRuleSpecConditions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecConditions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EventRuleSpecConditions)(ptr) = objs[0]
			} else {
				*(*EventRuleSpecConditions)(ptr) = EventRuleSpecConditions{}
			}
		} else {
			*(*EventRuleSpecConditions)(ptr) = EventRuleSpecConditions{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EventRuleSpecConditions

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecConditions{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EventRuleSpecConditions)(ptr) = obj
		} else {
			*(*EventRuleSpecConditions)(ptr) = EventRuleSpecConditions{}
		}
	default:
		iter.ReportError("decode EventRuleSpecConditions", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EventRuleSpecTimeFrameCodec struct {
}

func (EventRuleSpecTimeFrameCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EventRuleSpecTimeFrame)(ptr) == nil
}

func (EventRuleSpecTimeFrameCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EventRuleSpecTimeFrame)(ptr)
	var objs []EventRuleSpecTimeFrame
	if obj != nil {
		objs = []EventRuleSpecTimeFrame{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecTimeFrame{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EventRuleSpecTimeFrameCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EventRuleSpecTimeFrame)(ptr) = EventRuleSpecTimeFrame{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EventRuleSpecTimeFrame

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecTimeFrame{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EventRuleSpecTimeFrame)(ptr) = objs[0]
			} else {
				*(*EventRuleSpecTimeFrame)(ptr) = EventRuleSpecTimeFrame{}
			}
		} else {
			*(*EventRuleSpecTimeFrame)(ptr) = EventRuleSpecTimeFrame{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EventRuleSpecTimeFrame

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventRuleSpecTimeFrame{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EventRuleSpecTimeFrame)(ptr) = obj
		} else {
			*(*EventRuleSpecTimeFrame)(ptr) = EventRuleSpecTimeFrame{}
		}
	default:
		iter.ReportError("decode EventRuleSpecTimeFrame", "unexpected JSON type")
	}
}
