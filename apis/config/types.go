package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NodeConfiguration struct {
	metav1.TypeMeta

	// ContainerRuntimeConfiguration holds settings for the container runtime configuration
	ContainerRuntimeConfiguration ContainerRuntimeConfiguration
}

type ContainerRuntimeConfiguration struct {
	// +optional
	Version string

	// +optional
	Endpoint string
}
