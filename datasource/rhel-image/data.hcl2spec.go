// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package rhelimage

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	OfflineToken    *string `mapstructure:"offline_token" cty:"offline_token" hcl:"offline_token"`
	ImageChecksum   *string `mapstructure:"image_checksum" cty:"image_checksum" hcl:"image_checksum"`
	TargetDirectory *string `mapstructure:"target_directory" cty:"target_directory" hcl:"target_directory"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"offline_token":    &hcldec.AttrSpec{Name: "offline_token", Type: cty.String, Required: false},
		"image_checksum":   &hcldec.AttrSpec{Name: "image_checksum", Type: cty.String, Required: false},
		"target_directory": &hcldec.AttrSpec{Name: "target_directory", Type: cty.String, Required: false},
	}
	return s
}

// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDatasourceOutput struct {
	ImagePath *string `mapstructure:"image_path" cty:"image_path" hcl:"image_path"`
}

// FlatMapstructure returns a new FlatDatasourceOutput.
// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DatasourceOutput) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDatasourceOutput)
}

// HCL2Spec returns the hcl spec of a DatasourceOutput.
// This spec is used by HCL to read the fields of DatasourceOutput.
// The decoded values from this spec will then be applied to a FlatDatasourceOutput.
func (*FlatDatasourceOutput) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"image_path": &hcldec.AttrSpec{Name: "image_path", Type: cty.String, Required: false},
	}
	return s
}
