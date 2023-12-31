//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentDetail) DeepCopyInto(out *DeploymentDetail) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentDetail.
func (in *DeploymentDetail) DeepCopy() *DeploymentDetail {
	if in == nil {
		return nil
	}
	out := new(DeploymentDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderedDeployment) DeepCopyInto(out *OrderedDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderedDeployment.
func (in *OrderedDeployment) DeepCopy() *OrderedDeployment {
	if in == nil {
		return nil
	}
	out := new(OrderedDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderedDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderedDeploymentList) DeepCopyInto(out *OrderedDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OrderedDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderedDeploymentList.
func (in *OrderedDeploymentList) DeepCopy() *OrderedDeploymentList {
	if in == nil {
		return nil
	}
	out := new(OrderedDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderedDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderedDeploymentSpec) DeepCopyInto(out *OrderedDeploymentSpec) {
	*out = *in
	if in.DeploymentOrder != nil {
		in, out := &in.DeploymentOrder, &out.DeploymentOrder
		*out = make([]DeploymentDetail, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderedDeploymentSpec.
func (in *OrderedDeploymentSpec) DeepCopy() *OrderedDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(OrderedDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderedDeploymentStatus) DeepCopyInto(out *OrderedDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderedDeploymentStatus.
func (in *OrderedDeploymentStatus) DeepCopy() *OrderedDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(OrderedDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}
