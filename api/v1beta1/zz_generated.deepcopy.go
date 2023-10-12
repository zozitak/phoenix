//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright (C) 2023 R6 Security, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the Server Side Public License, version 1,
 * as published by MongoDB, Inc.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * Server Side Public License for more details.
 *
 * You should have received a copy of the Server Side Public License
 * along with this program. If not, see
 * <http://www.mongodb.com/licensing/server-side-public-license>.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdaptiveMovingTargetDefense) DeepCopyInto(out *AdaptiveMovingTargetDefense) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdaptiveMovingTargetDefense.
func (in *AdaptiveMovingTargetDefense) DeepCopy() *AdaptiveMovingTargetDefense {
	if in == nil {
		return nil
	}
	out := new(AdaptiveMovingTargetDefense)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdaptiveMovingTargetDefense) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdaptiveMovingTargetDefenseList) DeepCopyInto(out *AdaptiveMovingTargetDefenseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AdaptiveMovingTargetDefense, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdaptiveMovingTargetDefenseList.
func (in *AdaptiveMovingTargetDefenseList) DeepCopy() *AdaptiveMovingTargetDefenseList {
	if in == nil {
		return nil
	}
	out := new(AdaptiveMovingTargetDefenseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AdaptiveMovingTargetDefenseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdaptiveMovingTargetDefenseSpec) DeepCopyInto(out *AdaptiveMovingTargetDefenseSpec) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = make([]ResponseStrategy, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdaptiveMovingTargetDefenseSpec.
func (in *AdaptiveMovingTargetDefenseSpec) DeepCopy() *AdaptiveMovingTargetDefenseSpec {
	if in == nil {
		return nil
	}
	out := new(AdaptiveMovingTargetDefenseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdaptiveMovingTargetDefenseStatus) DeepCopyInto(out *AdaptiveMovingTargetDefenseStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdaptiveMovingTargetDefenseStatus.
func (in *AdaptiveMovingTargetDefenseStatus) DeepCopy() *AdaptiveMovingTargetDefenseStatus {
	if in == nil {
		return nil
	}
	out := new(AdaptiveMovingTargetDefenseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResponseStrategy) DeepCopyInto(out *ResponseStrategy) {
	*out = *in
	out.Rule = in.Rule
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResponseStrategy.
func (in *ResponseStrategy) DeepCopy() *ResponseStrategy {
	if in == nil {
		return nil
	}
	out := new(ResponseStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityEvent) DeepCopyInto(out *SecurityEvent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityEvent.
func (in *SecurityEvent) DeepCopy() *SecurityEvent {
	if in == nil {
		return nil
	}
	out := new(SecurityEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityEvent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityEventList) DeepCopyInto(out *SecurityEventList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityEvent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityEventList.
func (in *SecurityEventList) DeepCopy() *SecurityEventList {
	if in == nil {
		return nil
	}
	out := new(SecurityEventList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityEventList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityEventSpec) DeepCopyInto(out *SecurityEventSpec) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Rule = in.Rule
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityEventSpec.
func (in *SecurityEventSpec) DeepCopy() *SecurityEventSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityEventSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityEventStatus) DeepCopyInto(out *SecurityEventStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityEventStatus.
func (in *SecurityEventStatus) DeepCopy() *SecurityEventStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityEventStatus)
	in.DeepCopyInto(out)
	return out
}