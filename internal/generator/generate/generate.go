package generate

import (
	"argtools/internal/generator/configYaml"
	"argtools/internal/generator/idTemplateDataCreator"
)

// Generate - creates argParser.go file text body
func Generate(
	config *configYaml.Config,
	flagDescriptionMap map[configYaml.Flag]*configYaml.FlagDescription) string {

	creator := idTemplateDataCreator.NewIDTemplateCreator()
	commandsIDTemplateData, nullCommandIDTemplateData, flagsIDTemplateData := creator.CreateIDTemplateData(
		config.GetCommandDescriptions(),
		config.GetHelpCommandDescription(),
		config.GetNullCommandDescription(),
		flagDescriptionMap)

	return GenArgParserFileBody(
		GenCommandIDConstants(commandsIDTemplateData, nullCommandIDTemplateData),
		GenCommandStringIDConstants(commandsIDTemplateData),
		GenFlagStringIDConstants(flagsIDTemplateData),
		GenAppDescription(config.GetAppHelpDescription()),
		GenFlagMapElements(config.GetFlagDescriptions(), flagsIDTemplateData),
		GenCommandSliceElements(config.GetCommandDescriptions(), config.GetHelpCommandDescription(), commandsIDTemplateData, flagsIDTemplateData),
		GenNullCommandComponent(config.GetNullCommandDescription(), nullCommandIDTemplateData),
		commandsIDTemplateData[config.GetHelpCommandDescription().GetCommand()].GetID())
}
