package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	DefaultContainerRuntimeVersion  = "1.4.4"
	DefaultContainerRuntimeEndpoint = "unix:///var/run/containerd/containerd.sock"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_NodeConfiguration defaults the object
func SetDefaults_NodeConfiguration(obj *NodeConfiguration) {
	if len(obj.ContainerRuntimeConfiguration.Version) == 0 {
		obj.ContainerRuntimeConfiguration.Version = DefaultContainerRuntimeVersion
	}
	if len(obj.ContainerRuntimeConfiguration.Endpoint) == 0 {
		obj.ContainerRuntimeConfiguration.Endpoint = DefaultContainerRuntimeEndpoint
	}
}
