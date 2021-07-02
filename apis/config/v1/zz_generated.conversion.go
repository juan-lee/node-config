// +build !ignore_autogenerated

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	config "github.com/juan-lee/node-config/apis/config"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ContainerRuntimeConfiguration)(nil), (*config.ContainerRuntimeConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ContainerRuntimeConfiguration_To_config_ContainerRuntimeConfiguration(a.(*ContainerRuntimeConfiguration), b.(*config.ContainerRuntimeConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ContainerRuntimeConfiguration)(nil), (*ContainerRuntimeConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ContainerRuntimeConfiguration_To_v1_ContainerRuntimeConfiguration(a.(*config.ContainerRuntimeConfiguration), b.(*ContainerRuntimeConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeConfiguration)(nil), (*config.NodeConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeConfiguration_To_config_NodeConfiguration(a.(*NodeConfiguration), b.(*config.NodeConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeConfiguration)(nil), (*NodeConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeConfiguration_To_v1_NodeConfiguration(a.(*config.NodeConfiguration), b.(*NodeConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_ContainerRuntimeConfiguration_To_config_ContainerRuntimeConfiguration(in *ContainerRuntimeConfiguration, out *config.ContainerRuntimeConfiguration, s conversion.Scope) error {
	out.Version = in.Version
	out.Endpoint = in.Endpoint
	return nil
}

// Convert_v1_ContainerRuntimeConfiguration_To_config_ContainerRuntimeConfiguration is an autogenerated conversion function.
func Convert_v1_ContainerRuntimeConfiguration_To_config_ContainerRuntimeConfiguration(in *ContainerRuntimeConfiguration, out *config.ContainerRuntimeConfiguration, s conversion.Scope) error {
	return autoConvert_v1_ContainerRuntimeConfiguration_To_config_ContainerRuntimeConfiguration(in, out, s)
}

func autoConvert_config_ContainerRuntimeConfiguration_To_v1_ContainerRuntimeConfiguration(in *config.ContainerRuntimeConfiguration, out *ContainerRuntimeConfiguration, s conversion.Scope) error {
	out.Version = in.Version
	out.Endpoint = in.Endpoint
	return nil
}

// Convert_config_ContainerRuntimeConfiguration_To_v1_ContainerRuntimeConfiguration is an autogenerated conversion function.
func Convert_config_ContainerRuntimeConfiguration_To_v1_ContainerRuntimeConfiguration(in *config.ContainerRuntimeConfiguration, out *ContainerRuntimeConfiguration, s conversion.Scope) error {
	return autoConvert_config_ContainerRuntimeConfiguration_To_v1_ContainerRuntimeConfiguration(in, out, s)
}

func autoConvert_v1_NodeConfiguration_To_config_NodeConfiguration(in *NodeConfiguration, out *config.NodeConfiguration, s conversion.Scope) error {
	if err := Convert_v1_ContainerRuntimeConfiguration_To_config_ContainerRuntimeConfiguration(&in.ContainerRuntimeConfiguration, &out.ContainerRuntimeConfiguration, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_NodeConfiguration_To_config_NodeConfiguration is an autogenerated conversion function.
func Convert_v1_NodeConfiguration_To_config_NodeConfiguration(in *NodeConfiguration, out *config.NodeConfiguration, s conversion.Scope) error {
	return autoConvert_v1_NodeConfiguration_To_config_NodeConfiguration(in, out, s)
}

func autoConvert_config_NodeConfiguration_To_v1_NodeConfiguration(in *config.NodeConfiguration, out *NodeConfiguration, s conversion.Scope) error {
	if err := Convert_config_ContainerRuntimeConfiguration_To_v1_ContainerRuntimeConfiguration(&in.ContainerRuntimeConfiguration, &out.ContainerRuntimeConfiguration, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_NodeConfiguration_To_v1_NodeConfiguration is an autogenerated conversion function.
func Convert_config_NodeConfiguration_To_v1_NodeConfiguration(in *config.NodeConfiguration, out *NodeConfiguration, s conversion.Scope) error {
	return autoConvert_config_NodeConfiguration_To_v1_NodeConfiguration(in, out, s)
}