package main

import (
	"fmt"
	"os"
	rhelImageData "packer-plugin-rhel-image/datasource/rhel-image"
	rhelImageVersion "packer-plugin-rhel-image/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("my-datasource", new(rhelImageData.Datasource))
	pps.SetVersion(rhelImageVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
