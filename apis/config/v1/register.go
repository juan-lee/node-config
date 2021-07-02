package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeBuilder the schema builder
	SchemeBuilder = runtime.NewSchemeBuilder(
		addKnownTypes,
		addDefaultingFuncs,
	)
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

// GroupName is the group name use in this package
const GroupName = "config.juan-lee.dev"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{
	Group:   GroupName,
	Version: "v1",
}

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&NodeConfiguration{},
	)
	return nil
}
