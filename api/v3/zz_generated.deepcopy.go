//go:build !ignore_autogenerated
// +build !ignore_autogenerated

//
// Copyright 2022 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by controller-gen. DO NOT EDIT.

package v3

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APICatalog) DeepCopyInto(out *APICatalog) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APICatalog.
func (in *APICatalog) DeepCopy() *APICatalog {
	if in == nil {
		return nil
	}
	out := new(APICatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BedrockOperator) DeepCopyInto(out *BedrockOperator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BedrockOperator.
func (in *BedrockOperator) DeepCopy() *BedrockOperator {
	if in == nil {
		return nil
	}
	out := new(BedrockOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bedrockshim) DeepCopyInto(out *Bedrockshim) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bedrockshim.
func (in *Bedrockshim) DeepCopy() *Bedrockshim {
	if in == nil {
		return nil
	}
	out := new(Bedrockshim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSData) DeepCopyInto(out *CSData) {
	*out = *in
	if in.ServicesConfig != nil {
		in, out := &in.ServicesConfig, &out.ServicesConfig
		*out = make(map[string]ServiceConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSData.
func (in *CSData) DeepCopy() *CSData {
	if in == nil {
		return nil
	}
	out := new(CSData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonService) DeepCopyInto(out *CommonService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonService.
func (in *CommonService) DeepCopy() *CommonService {
	if in == nil {
		return nil
	}
	out := new(CommonService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CommonService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonServiceList) DeepCopyInto(out *CommonServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CommonService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonServiceList.
func (in *CommonServiceList) DeepCopy() *CommonServiceList {
	if in == nil {
		return nil
	}
	out := new(CommonServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CommonServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonServiceSpec) DeepCopyInto(out *CommonServiceSpec) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(Features)
		(*in).DeepCopyInto(*out)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServiceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.License = in.License
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonServiceSpec.
func (in *CommonServiceSpec) DeepCopy() *CommonServiceSpec {
	if in == nil {
		return nil
	}
	out := new(CommonServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonServiceStatus) DeepCopyInto(out *CommonServiceStatus) {
	*out = *in
	if in.BedrockOperators != nil {
		in, out := &in.BedrockOperators, &out.BedrockOperators
		*out = make([]BedrockOperator, len(*in))
		copy(*out, *in)
	}
	in.ConfigStatus.DeepCopyInto(&out.ConfigStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonServiceStatus.
func (in *CommonServiceStatus) DeepCopy() *CommonServiceStatus {
	if in == nil {
		return nil
	}
	out := new(CommonServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigStatus) DeepCopyInto(out *ConfigStatus) {
	*out = *in
	if in.TopologyConfigurableCRs != nil {
		in, out := &in.TopologyConfigurableCRs, &out.TopologyConfigurableCRs
		*out = make([]ConfigurableCR, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigStatus.
func (in *ConfigStatus) DeepCopy() *ConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurableCR) DeepCopyInto(out *ConfigurableCR) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurableCR.
func (in *ConfigurableCR) DeepCopy() *ConfigurableCR {
	if in == nil {
		return nil
	}
	out := new(ConfigurableCR)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	if in.Bedrockshim != nil {
		in, out := &in.Bedrockshim, &out.Bedrockshim
		*out = new(Bedrockshim)
		**out = **in
	}
	if in.APICatalog != nil {
		in, out := &in.APICatalog, &out.APICatalog
		*out = new(APICatalog)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseList) DeepCopyInto(out *LicenseList) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseList.
func (in *LicenseList) DeepCopy() *LicenseList {
	if in == nil {
		return nil
	}
	out := new(LicenseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConfig) DeepCopyInto(out *ServiceConfig) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make(map[string]runtime.RawExtension, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConfig.
func (in *ServiceConfig) DeepCopy() *ServiceConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceConfig)
	in.DeepCopyInto(out)
	return out
}
