package cargo

import (
	"errors"
	"fmt"

	"github.com/devsecstack/jf-devsecstack-plugin/commands/utils"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

func CargoAddDependencies() components.Command {
	return components.Command{
		Name:        "cargo-add-dependencies",
		Aliases:     []string{"cad"},
		Description: "Add dependencies to build-info from Cargo.lock file",
		Flags: []components.Flag{
			components.NewStringFlag(utils.BuildName,"Build name"),
			components.NewStringFlag(utils.BuildNumber,"Build number"),
			components.NewStringFlag(utils.Project, "[Optional]\nJFrog project key."),
			components.NewStringFlag(utils.Module,"[Optional]\nOptional module name in the build-info for adding the dependency."),
			components.NewStringFlag(utils.Server,"Artifactory server ID"),
			components.NewBoolFlag(utils.DryRun,"[Default: false]\nSet to true to disable communication with Artifactory."),
		},
		Action: func(c *components.Context) error {
			return cargoAddDependenciesCmd(c)
		},
	}
}

func cargoAddDependenciesCmd(c *components.Context) error {
	fmt.Printf("%s Adding dependencies to build-info from Cargo.lock file...", utils.UnquoteCodePoint("\\U0001F980"))
	packages, err := utils.GetPackages()
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to get packages: %s", err))
		return errors.New("failed to get packages")
	}

	err = utils.BuildAddDependencies(c, packages)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to add dependencies: %s", err))
		return errors.New("failed to add dependencies")
	}

	return nil
}
