//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The Kubernetes Authors.

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

package internalversion

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]Env, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]Volume, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Env) DeepCopyInto(out *Env) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Env.
func (in *Env) DeepCopy() *Env {
	if in == nil {
		return nil
	}
	out := new(Env)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressionFromSource) DeepCopyInto(out *ExpressionFromSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressionFromSource.
func (in *ExpressionFromSource) DeepCopy() *ExpressionFromSource {
	if in == nil {
		return nil
	}
	out := new(ExpressionFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FinalizerItem) DeepCopyInto(out *FinalizerItem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FinalizerItem.
func (in *FinalizerItem) DeepCopy() *FinalizerItem {
	if in == nil {
		return nil
	}
	out := new(FinalizerItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KwokConfiguration) DeepCopyInto(out *KwokConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Options = in.Options
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KwokConfiguration.
func (in *KwokConfiguration) DeepCopy() *KwokConfiguration {
	if in == nil {
		return nil
	}
	out := new(KwokConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KwokConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KwokConfigurationOptions) DeepCopyInto(out *KwokConfigurationOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KwokConfigurationOptions.
func (in *KwokConfigurationOptions) DeepCopy() *KwokConfigurationOptions {
	if in == nil {
		return nil
	}
	out := new(KwokConfigurationOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KwokctlConfiguration) DeepCopyInto(out *KwokctlConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Options.DeepCopyInto(&out.Options)
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]Component, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KwokctlConfiguration.
func (in *KwokctlConfiguration) DeepCopy() *KwokctlConfiguration {
	if in == nil {
		return nil
	}
	out := new(KwokctlConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KwokctlConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KwokctlConfigurationOptions) DeepCopyInto(out *KwokctlConfigurationOptions) {
	*out = *in
	if in.Runtimes != nil {
		in, out := &in.Runtimes, &out.Runtimes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KwokctlConfigurationOptions.
func (in *KwokctlConfigurationOptions) DeepCopy() *KwokctlConfigurationOptions {
	if in == nil {
		return nil
	}
	out := new(KwokctlConfigurationOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorRequirement) DeepCopyInto(out *SelectorRequirement) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorRequirement.
func (in *SelectorRequirement) DeepCopy() *SelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(SelectorRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stage) DeepCopyInto(out *Stage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stage.
func (in *Stage) DeepCopy() *Stage {
	if in == nil {
		return nil
	}
	out := new(Stage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageDelay) DeepCopyInto(out *StageDelay) {
	*out = *in
	if in.DurationMilliseconds != nil {
		in, out := &in.DurationMilliseconds, &out.DurationMilliseconds
		*out = new(int64)
		**out = **in
	}
	if in.DurationFrom != nil {
		in, out := &in.DurationFrom, &out.DurationFrom
		*out = new(ExpressionFromSource)
		**out = **in
	}
	if in.JitterDurationMilliseconds != nil {
		in, out := &in.JitterDurationMilliseconds, &out.JitterDurationMilliseconds
		*out = new(int64)
		**out = **in
	}
	if in.JitterDurationFrom != nil {
		in, out := &in.JitterDurationFrom, &out.JitterDurationFrom
		*out = new(ExpressionFromSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageDelay.
func (in *StageDelay) DeepCopy() *StageDelay {
	if in == nil {
		return nil
	}
	out := new(StageDelay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageEvent) DeepCopyInto(out *StageEvent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageEvent.
func (in *StageEvent) DeepCopy() *StageEvent {
	if in == nil {
		return nil
	}
	out := new(StageEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageFinalizers) DeepCopyInto(out *StageFinalizers) {
	*out = *in
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make([]FinalizerItem, len(*in))
		copy(*out, *in)
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]FinalizerItem, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageFinalizers.
func (in *StageFinalizers) DeepCopy() *StageFinalizers {
	if in == nil {
		return nil
	}
	out := new(StageFinalizers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageNext) DeepCopyInto(out *StageNext) {
	*out = *in
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		*out = new(StageEvent)
		**out = **in
	}
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = new(StageFinalizers)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageNext.
func (in *StageNext) DeepCopy() *StageNext {
	if in == nil {
		return nil
	}
	out := new(StageNext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageResourceRef) DeepCopyInto(out *StageResourceRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageResourceRef.
func (in *StageResourceRef) DeepCopy() *StageResourceRef {
	if in == nil {
		return nil
	}
	out := new(StageResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageSelector) DeepCopyInto(out *StageSelector) {
	*out = *in
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchAnnotations != nil {
		in, out := &in.MatchAnnotations, &out.MatchAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]SelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageSelector.
func (in *StageSelector) DeepCopy() *StageSelector {
	if in == nil {
		return nil
	}
	out := new(StageSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageSpec) DeepCopyInto(out *StageSpec) {
	*out = *in
	out.ResourceRef = in.ResourceRef
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(StageSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Delay != nil {
		in, out := &in.Delay, &out.Delay
		*out = new(StageDelay)
		(*in).DeepCopyInto(*out)
	}
	in.Next.DeepCopyInto(&out.Next)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageSpec.
func (in *StageSpec) DeepCopy() *StageSpec {
	if in == nil {
		return nil
	}
	out := new(StageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}
