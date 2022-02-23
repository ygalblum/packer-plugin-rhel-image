package main

import (
	"fmt"
	"os"
	rhelData "packer-plugin-scaffolding/datasource/rhel"
	rhelVersion "packer-plugin-scaffolding/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("my-datasource", new(rhelData.Datasource))
	pps.SetVersion(rhelVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
