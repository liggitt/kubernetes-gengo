//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by defaulter-gen. DO NOT EDIT.

package marker

import (
	"encoding/json"

	runtime "k8s.io/apimachinery/pkg/runtime"
	external "k8s.io/gengo/v2/examples/defaulter-gen/output_tests/marker/external"
	externalexternal "k8s.io/gengo/v2/examples/defaulter-gen/output_tests/marker/external/external"
	external2 "k8s.io/gengo/v2/examples/defaulter-gen/output_tests/marker/external2"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Defaulted{}, func(obj interface{}) { SetObjectDefaults_Defaulted(obj.(*Defaulted)) })
	scheme.AddTypeDefaultingFunc(&DefaultedOmitempty{}, func(obj interface{}) { SetObjectDefaults_DefaultedOmitempty(obj.(*DefaultedOmitempty)) })
	scheme.AddTypeDefaultingFunc(&DefaultedWithFunction{}, func(obj interface{}) { SetObjectDefaults_DefaultedWithFunction(obj.(*DefaultedWithFunction)) })
	scheme.AddTypeDefaultingFunc(&DefaultedWithReference{}, func(obj interface{}) { SetObjectDefaults_DefaultedWithReference(obj.(*DefaultedWithReference)) })
	return nil
}

func SetObjectDefaults_Defaulted(in *Defaulted) {
	if in.StringDefault == "" {
		in.StringDefault = "bar"
	}
	if in.StringPointer == nil {
		var ptrVar1 string = "default"
		in.StringPointer = &ptrVar1
	}
	if in.Int64 == nil {
		var ptrVar1 int64 = 64
		in.Int64 = &ptrVar1
	}
	if in.Int32 == nil {
		var ptrVar1 int32 = 32
		in.Int32 = &ptrVar1
	}
	if in.IntDefault == 0 {
		in.IntDefault = 1
	}
	if in.FloatDefault == 0 {
		in.FloatDefault = 0.5
	}
	if in.List == nil {
		if err := json.Unmarshal([]byte(`["foo", "bar"]`), &in.List); err != nil {
			panic(err)
		}
	}
	for i := range in.List {
		if in.List[i] == nil {
			var ptrVar1 string = "apple"
			in.List[i] = &ptrVar1
		}
	}
	if in.Sub == nil {
		if err := json.Unmarshal([]byte(`{"s": "foo", "i": 5}`), &in.Sub); err != nil {
			panic(err)
		}
	}
	if in.Sub != nil {
		if in.Sub.I == 0 {
			in.Sub.I = 1
		}
	}
	if in.StructList == nil {
		if err := json.Unmarshal([]byte(`[{"s": "foo1", "i": 1}, {"s": "foo2"}]`), &in.StructList); err != nil {
			panic(err)
		}
	}
	for i := range in.StructList {
		a := &in.StructList[i]
		if a.I == 0 {
			a.I = 1
		}
	}
	if in.PtrStructList == nil {
		if err := json.Unmarshal([]byte(`[{"s": "foo1", "i": 1}, {"s": "foo2"}]`), &in.PtrStructList); err != nil {
			panic(err)
		}
	}
	for i := range in.PtrStructList {
		a := in.PtrStructList[i]
		if a != nil {
			if a.I == 0 {
				a.I = 1
			}
		}
	}
	if in.StringList == nil {
		if err := json.Unmarshal([]byte(`["foo"]`), &in.StringList); err != nil {
			panic(err)
		}
	}
	if in.OtherSub.I == 0 {
		in.OtherSub.I = 1
	}
	if in.Map == nil {
		if err := json.Unmarshal([]byte(`{"foo": "bar"}`), &in.Map); err != nil {
			panic(err)
		}
	}
	for i_Map := range in.Map {
		if in.Map[i_Map] == nil {
			var ptrVar1 string = "apple"
			in.Map[i_Map] = &ptrVar1
		}
	}
	if in.StructMap == nil {
		if err := json.Unmarshal([]byte(`{"foo": {"S": "string", "I": 1}}`), &in.StructMap); err != nil {
			panic(err)
		}
	}
	if in.PtrStructMap == nil {
		if err := json.Unmarshal([]byte(`{"foo": {"S": "string", "I": 1}}`), &in.PtrStructMap); err != nil {
			panic(err)
		}
	}
	if in.AliasPtr == nil {
		var ptrVar1 string = "banana"
		in.AliasPtr = &ptrVar1
	}
}

func SetObjectDefaults_DefaultedOmitempty(in *DefaultedOmitempty) {
	if in.StringDefault == "" {
		in.StringDefault = "bar"
	}
	if in.StringPointer == nil {
		var ptrVar1 string = "default"
		in.StringPointer = &ptrVar1
	}
	if in.Int64 == nil {
		var ptrVar1 int64 = 64
		in.Int64 = &ptrVar1
	}
	if in.Int32 == nil {
		var ptrVar1 int32 = 32
		in.Int32 = &ptrVar1
	}
	if in.IntDefault == 0 {
		in.IntDefault = 1
	}
	if in.FloatDefault == 0 {
		in.FloatDefault = 0.5
	}
	if in.List == nil {
		if err := json.Unmarshal([]byte(`["foo", "bar"]`), &in.List); err != nil {
			panic(err)
		}
	}
	for i := range in.List {
		if in.List[i] == nil {
			var ptrVar1 string = "apple"
			in.List[i] = &ptrVar1
		}
	}
	if in.Sub == nil {
		if err := json.Unmarshal([]byte(`{"s": "foo", "i": 5}`), &in.Sub); err != nil {
			panic(err)
		}
	}
	if in.Sub != nil {
		if in.Sub.I == 0 {
			in.Sub.I = 1
		}
	}
	if in.StructList == nil {
		if err := json.Unmarshal([]byte(`[{"s": "foo1", "i": 1}, {"s": "foo2"}]`), &in.StructList); err != nil {
			panic(err)
		}
	}
	for i := range in.StructList {
		a := &in.StructList[i]
		if a.I == 0 {
			a.I = 1
		}
	}
	if in.PtrStructList == nil {
		if err := json.Unmarshal([]byte(`[{"s": "foo1", "i": 1}, {"s": "foo2"}]`), &in.PtrStructList); err != nil {
			panic(err)
		}
	}
	for i := range in.PtrStructList {
		a := in.PtrStructList[i]
		if a != nil {
			if a.I == 0 {
				a.I = 1
			}
		}
	}
	if in.StringList == nil {
		if err := json.Unmarshal([]byte(`["foo"]`), &in.StringList); err != nil {
			panic(err)
		}
	}
	if in.OtherSub.I == 0 {
		in.OtherSub.I = 1
	}
	if in.Map == nil {
		if err := json.Unmarshal([]byte(`{"foo": "bar"}`), &in.Map); err != nil {
			panic(err)
		}
	}
	for i_Map := range in.Map {
		if in.Map[i_Map] == nil {
			var ptrVar1 string = "apple"
			in.Map[i_Map] = &ptrVar1
		}
	}
	if in.StructMap == nil {
		if err := json.Unmarshal([]byte(`{"foo": {"S": "string", "I": 1}}`), &in.StructMap); err != nil {
			panic(err)
		}
	}
	if in.PtrStructMap == nil {
		if err := json.Unmarshal([]byte(`{"foo": {"S": "string", "I": 1}}`), &in.PtrStructMap); err != nil {
			panic(err)
		}
	}
	if in.AliasPtr == nil {
		var ptrVar1 string = "banana"
		in.AliasPtr = &ptrVar1
	}
}

func SetObjectDefaults_DefaultedWithFunction(in *DefaultedWithFunction) {
	SetDefaults_DefaultedWithFunction(in)
	if in.S1 == "" {
		in.S1 = "default_marker"
	}
	if in.S2 == "" {
		in.S2 = "default_marker"
	}
}

func SetObjectDefaults_DefaultedWithReference(in *DefaultedWithReference) {
	if in.AliasConvertDefaultPointer == nil {
		ptrVar1 := DefaultedValueItem(SomeValue)
		in.AliasConvertDefaultPointer = &ptrVar1
	}
	if in.PointerAliasDefault == nil {
		var ptrVar1 string = "apple"
		in.PointerAliasDefault = &ptrVar1
	}
	if in.AliasPointerInside == nil {
		ptrVar1 := string(SomeDefault)
		in.AliasPointerInside = &ptrVar1
	}
	if in.AliasOverride == nil {
		ptrVar1 := string(SomeDefault)
		in.AliasOverride = &ptrVar1
	}
	if in.AliasPointerDefault == nil {
		ptrVar1 := DefaultedValueItem(SomeValue)
		in.AliasPointerDefault = &ptrVar1
	}
	if in.AliasNonPointer == "" {
		in.AliasNonPointer = ValueItem(SomeValue)
	}
	if in.AliasPointer == nil {
		ptrVar1 := ValueItem(SomeValue)
		in.AliasPointer = &ptrVar1
	}
	if in.SymbolReference == "" {
		in.SymbolReference = string(SomeDefault)
	}
	if in.SameNamePackageSymbolReference1 == "" {
		in.SameNamePackageSymbolReference1 = string(external.AConstant)
	}
	if in.SameNamePackageSymbolReference2 == "" {
		in.SameNamePackageSymbolReference2 = string(externalexternal.AnotherConstant)
	}
	if in.PointerConversion == nil {
		ptrVar9 := string(SomeValue)
		ptrVar8 := &ptrVar9
		ptrVar7 := (*B1)(&ptrVar8)
		ptrVar6 := (*B2)(&ptrVar7)
		ptrVar5 := &ptrVar6
		ptrVar4 := &ptrVar5
		ptrVar3 := &ptrVar4
		ptrVar2 := (*B3)(&ptrVar3)
		ptrVar1 := &ptrVar2
		in.PointerConversion = (*B4)(&ptrVar1)
	}
	if in.PointerConversionValue == nil {
		ptrVar8 := string(SomeValue)
		ptrVar7 := &ptrVar8
		ptrVar6 := (*B1)(&ptrVar7)
		ptrVar5 := (*B2)(&ptrVar6)
		ptrVar4 := &ptrVar5
		ptrVar3 := &ptrVar4
		ptrVar2 := &ptrVar3
		ptrVar1 := (*B3)(&ptrVar2)
		in.PointerConversionValue = &ptrVar1
	}
	if in.FullyQualifiedLocalSymbol == "" {
		in.FullyQualifiedLocalSymbol = string(SomeValue)
	}
	if in.ImportFromAliasCast == nil {
		ptrVar1 := external2.String(SomeValue)
		in.ImportFromAliasCast = &ptrVar1
	}
}
