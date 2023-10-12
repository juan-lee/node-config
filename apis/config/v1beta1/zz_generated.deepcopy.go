//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRuntimeConfiguration) DeepCopyInto(out *ContainerRuntimeConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRuntimeConfiguration.
func (in *ContainerRuntimeConfiguration) DeepCopy() *ContainerRuntimeConfiguration {
	if in == nil {
		return nil
	}
	out := new(ContainerRuntimeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfiguration) DeepCopyInto(out *NodeConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ContainerRuntimeConfiguration = in.ContainerRuntimeConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfiguration.
func (in *NodeConfiguration) DeepCopy() *NodeConfiguration {
	if in == nil {
		return nil
	}
	out := new(NodeConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
