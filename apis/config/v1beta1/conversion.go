package v1beta1

import (
	"github.com/juan-lee/node-config/apis/config"

	"k8s.io/apimachinery/pkg/conversion"
)

func Convert_v1beta1_ContainerRuntimeConfiguration_To_config_ContainerRuntimeConfiguration(in *ContainerRuntimeConfiguration, out *config.ContainerRuntimeConfiguration, s conversion.Scope) error {
	if err := autoConvert_v1beta1_ContainerRuntimeConfiguration_To_config_ContainerRuntimeConfiguration(in, out, s); err != nil {
		return err
	}
	out.Version = in.Version
	return nil
}

func Convert_config_ContainerRuntimeConfiguration_To_v1beta1_ContainerRuntimeConfiguration(in *config.ContainerRuntimeConfiguration, out *ContainerRuntimeConfiguration, s conversion.Scope) error {
	if err := autoConvert_config_ContainerRuntimeConfiguration_To_v1beta1_ContainerRuntimeConfiguration(in, out, s); err != nil {
		return err
	}
	out.Version = in.Version
	return nil
}
