package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_NodeConfiguration defaults the object
func SetDefaults_NodeConfiguration(obj *NodeConfiguration) {
	if len(obj.ContainerRuntimeConfiguration.Version) == 0 {
		obj.ContainerRuntimeConfiguration.Version = "1.4.4"
	}
}
