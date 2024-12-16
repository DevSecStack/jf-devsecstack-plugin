package main

import (
	"github.com/devsecstack/jf-devsecstack-plugin/commands/cargo"
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = Name
	app.Description = Description
	app.Version = Version
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		cargo.CargoAddDependencies()}
}
