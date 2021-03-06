// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricHPATrait) DeepCopyInto(out *MetricHPATrait) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricHPATrait.
func (in *MetricHPATrait) DeepCopy() *MetricHPATrait {
	if in == nil {
		return nil
	}
	out := new(MetricHPATrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricHPATrait) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricHPATraitList) DeepCopyInto(out *MetricHPATraitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricHPATrait, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricHPATraitList.
func (in *MetricHPATraitList) DeepCopy() *MetricHPATraitList {
	if in == nil {
		return nil
	}
	out := new(MetricHPATraitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricHPATraitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricHPATraitSpec) DeepCopyInto(out *MetricHPATraitSpec) {
	*out = *in
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.PromThreshold != nil {
		in, out := &in.PromThreshold, &out.PromThreshold
		*out = new(int32)
		**out = **in
	}
	out.WorkloadReference = in.WorkloadReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricHPATraitSpec.
func (in *MetricHPATraitSpec) DeepCopy() *MetricHPATraitSpec {
	if in == nil {
		return nil
	}
	out := new(MetricHPATraitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricHPATraitStatus) DeepCopyInto(out *MetricHPATraitStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1alpha1.TypedReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricHPATraitStatus.
func (in *MetricHPATraitStatus) DeepCopy() *MetricHPATraitStatus {
	if in == nil {
		return nil
	}
	out := new(MetricHPATraitStatus)
	in.DeepCopyInto(out)
	return out
}
