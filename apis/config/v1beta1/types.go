package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// ContainerRuntimeConfiguration holds settings for the container runtime configuration
	ContainerRuntimeConfiguration ContainerRuntimeConfiguration `json:"containerRuntime"`
}

type ContainerRuntimeConfiguration struct {
	// +optional
	Version string `json:"version"`
}
