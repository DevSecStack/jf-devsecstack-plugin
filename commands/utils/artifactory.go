package utils

import (
	"fmt"

	buildinfo "github.com/jfrog/jfrog-cli-core/v2/artifactory/commands/buildinfo"
	"github.com/jfrog/jfrog-cli-core/v2/common/build"
	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-core/v2/common/spec"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	serviceUtils "github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

type ArtifactoryDetails struct {
	BuildConfiguration *build.BuildConfiguration
	ServerDetails      *config.ServerDetails
	CommonConf         serviceUtils.CommonConf
}

func GetArtifactoryDetails(c *components.Context) (details *ArtifactoryDetails, err error) {
	details = new(ArtifactoryDetails)

	details.ServerDetails, err = commands.GetConfig(c.GetStringFlagValue("server"), true)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to get Artifactory details: %s", err))
		return nil, err
	}

	authConfig, err := details.ServerDetails.CreateArtAuthConfig()
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to create auth config: %s", err))
		return nil, err
	}

	details.CommonConf, err = serviceUtils.NewCommonConfImpl(authConfig)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to create common config: %s", err))
		return nil, err
	}

	details.BuildConfiguration, err = CreateBuildConfigurationWithModule(c)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to create build configuration: %s", err))
		return nil, err
	}
	return details, nil
}

func CreateBuildConfigurationWithModule(c *components.Context) (buildConfig *build.BuildConfiguration, err error) {
	log.Debug(fmt.Sprintf("Creating build configuration: %s %s %s %s", c.GetStringFlagValue(BuildName), c.GetStringFlagValue(BuildNumber), c.GetStringFlagValue(Project), c.GetStringFlagValue(Module)))
	buildConfig = new(build.BuildConfiguration)
	err = buildConfig.SetBuildName(c.GetStringFlagValue(BuildName)).
		SetBuildNumber(c.GetStringFlagValue(BuildNumber)).
		SetProject(c.GetStringFlagValue(Project)).
		SetModule(c.GetStringFlagValue(Module)).
		ValidateBuildAndModuleParams()
	return
}

func CreateDefaultBuildAddDependenciesSpec(packages []Package) *spec.SpecFiles {
	dependenciesSpec := spec.SpecFiles{
		Files: []spec.File{},
	}
	for _, pkg := range packages {
		dependenciesSpec.Files = append(dependenciesSpec.Files, spec.File{
			Aql: serviceUtils.Aql{
				ItemsFind: fmt.Sprintf(`{"sha256":"%s"}`, pkg.Checksum),
			},
		})
	}
	return &dependenciesSpec
}

func BuildAddDependencies(c *components.Context, packages []Package) error {
	log.Debug(fmt.Sprintf("Adding %d dependencies to build-info", len(packages)))
	details, err := GetArtifactoryDetails(c)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to get Artifactory details: %s", err))
		return err
	}

	dependenciesSpec := CreateDefaultBuildAddDependenciesSpec(packages)

	buildAddDependenciesCmd := buildinfo.NewBuildAddDependenciesCommand().
		SetDryRun(c.GetBoolFlagValue(DryRun)).
		SetBuildConfiguration(details.BuildConfiguration).
		SetDependenciesSpec(dependenciesSpec).
		SetServerDetails(details.ServerDetails)
	buildAddDependenciesCmd.Run()

	return nil
}