package generate

import (
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
)

// Generate creates argParser.go file text body
func Generate(
	config *configYaml.Config,
	flagDescriptionMap map[string]*configYaml.FlagDescription,
) string {

	creator := idTemplateDataCreator.NewIDTemplateCreator()
	commandsIDTemplateData, namelessCommandIDTemplateData, flagsIDTemplateData := creator.CreateIDTemplateData(
		config.GetCommandDescriptions(),
		config.GetHelpCommandDescription(),
		config.GetNamelessCommandDescription(),
		flagDescriptionMap)

	return GenArgParserFileBody(
		GenCommandIDListSection(commandsIDTemplateData, namelessCommandIDTemplateData),
		GenCommandListSection(commandsIDTemplateData),
		GenFlagIDConstantsSection(flagsIDTemplateData),
		GenAppDescriptionSection(config.GetAppHelpDescription()),
		GenFlagDescriptionsSection(config.GetFlagDescriptions(), flagsIDTemplateData),
		GenCommandDescriptionsSection(config.GetCommandDescriptions(), commandsIDTemplateData, flagsIDTemplateData),
		GenHelpCommandDescriptionSection(config.GetHelpCommandDescription()),
		GenNamelessCommandComponent(config.GetNamelessCommandDescription(), namelessCommandIDTemplateData, flagsIDTemplateData),
		commandsIDTemplateData[config.GetHelpCommandDescription().GetCommand()].GetID())
}
