// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package strconf

import (
	v1 "k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSource) DeepCopyInto(out *ConfigSource) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSource.
func (in *ConfigSource) DeepCopy() *ConfigSource {
	if in == nil {
		return nil
	}
	out := new(ConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringOrConfig) DeepCopyInto(out *StringOrConfig) {
	*out = *in
	in.ConfigVal.DeepCopyInto(&out.ConfigVal)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringOrConfig.
func (in *StringOrConfig) DeepCopy() *StringOrConfig {
	if in == nil {
		return nil
	}
	out := new(StringOrConfig)
	in.DeepCopyInto(out)
	return out
}
