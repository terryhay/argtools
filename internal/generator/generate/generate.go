package generate

import (
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
)

// Generate - creates argParser.go file text body
func Generate(
	config *configYaml.Config,
	flagDescriptionMap map[configYaml.Flag]*configYaml.FlagDescription) string {

	creator := idTemplateDataCreator.NewIDTemplateCreator()
	commandsIDTemplateData, namelessCommandIDTemplateData, flagsIDTemplateData := creator.CreateIDTemplateData(
		config.GetCommandDescriptions(),
		config.GetHelpCommandDescription(),
		config.GetNamelessCommandDescription(),
		flagDescriptionMap)

	return GenArgParserFileBody(
		GenCommandIDConstants(commandsIDTemplateData, namelessCommandIDTemplateData),
		GenCommandStringIDConstants(commandsIDTemplateData),
		GenFlagStringIDConstants(flagsIDTemplateData),
		GenAppDescription(config.GetAppHelpDescription()),
		GenFlagMapElements(config.GetFlagDescriptions(), flagsIDTemplateData),
		GenCommandSliceElements(config.GetCommandDescriptions(), config.GetHelpCommandDescription(), commandsIDTemplateData, flagsIDTemplateData),
		GenNamelessCommandComponent(config.GetNamelessCommandDescription(), namelessCommandIDTemplateData),
		commandsIDTemplateData[config.GetHelpCommandDescription().GetCommand()].GetID())
}
