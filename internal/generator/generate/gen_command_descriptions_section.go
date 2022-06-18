package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

const (
	commandDescriptionSliceNilPart = `
		// commandDescriptions
		nil`

	commandDescriptionSliceElementPrefixPart = `
		// commandDescriptions
		[]*argParserConfig.CommandDescription{`
	commandDescriptionSliceElementRequiredPart = `
			{
				ID:                  %s,
				DescriptionHelpInfo: "%s",`
	commandDescriptionSliceElementCommandsPart = `
				Commands: map[argParserConfig.Command]bool{%s
				},`
	commandDescriptionSliceElementRequiredFlagsPart = `
				RequiredFlags: map[argParserConfig.Flag]bool{%s
				},`
	commandDescriptionSliceElementOptionalFlagsPart = `
				OptionalFlags: map[argParserConfig.Flag]bool{%s
				},`
	commandDescriptionSliceElementPostfix = `
			},`
	commandDescriptionSlicePostfix = `
		}`
)

// CommandDescriptionsSection - string with command constant definitions list
type CommandDescriptionsSection string

// GenCommandDescriptionsSection creates a paste section with command descriptions
func GenCommandDescriptionsSection(
	commandDescriptions []*configYaml.CommandDescription,
	commandsIDTemplateData map[string]*idTemplateDataCreator.IDTemplateData,
	flagsIDTemplateData map[string]*idTemplateDataCreator.IDTemplateData,
) CommandDescriptionsSection {

	if len(commandDescriptions) == 0 {
		return commandDescriptionSliceNilPart
	}

	builder := strings.Builder{}
	builder.WriteString(commandDescriptionSliceElementPrefixPart)

	var (
		commandDescription *configYaml.CommandDescription
		i, j               int
	)

	idTemplateDataSlice := make([]*idTemplateDataCreator.IDTemplateData, 0, 8)
	for i = range commandDescriptions {
		commandDescription = commandDescriptions[i]

		idTemplateDataSlice = append(idTemplateDataSlice, commandsIDTemplateData[commandDescription.GetCommand()])
		for j = range commandDescription.GetAdditionalCommands() {
			idTemplateDataSlice = append(idTemplateDataSlice, commandsIDTemplateData[commandDescription.GetAdditionalCommands()[j]])
		}

		builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementRequiredPart,
			commandsIDTemplateData[commandDescription.GetCommand()].GetID(),
			commandDescription.GetDescriptionHelpInfo()))
		builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementCommandsPart, joinCallNames(idTemplateDataSlice)))
		idTemplateDataSlice = []*idTemplateDataCreator.IDTemplateData{}

		if len(commandDescription.GetRequiredFlags()) > 0 {
			for j = range commandDescription.GetRequiredFlags() {
				idTemplateDataSlice = append(idTemplateDataSlice, flagsIDTemplateData[commandDescription.GetRequiredFlags()[j]])
			}
			builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementRequiredFlagsPart, joinCallNames(idTemplateDataSlice)))
			idTemplateDataSlice = []*idTemplateDataCreator.IDTemplateData{}
		}

		if len(commandDescription.GetOptionalFlags()) > 0 {
			for j = range commandDescription.GetOptionalFlags() {
				idTemplateDataSlice = append(idTemplateDataSlice, flagsIDTemplateData[commandDescription.GetOptionalFlags()[j]])
			}
			builder.WriteString(fmt.Sprintf(commandDescriptionSliceElementOptionalFlagsPart, joinCallNames(idTemplateDataSlice)))
			idTemplateDataSlice = []*idTemplateDataCreator.IDTemplateData{}
		}

		builder.WriteString(commandDescriptionSliceElementPostfix)
	}

	builder.WriteString(commandDescriptionSlicePostfix)

	return CommandDescriptionsSection(builder.String())
}

func joinCallNames(nameAndIDSlice []*idTemplateDataCreator.IDTemplateData) (res string) {
	for i := range nameAndIDSlice {
		res += fmt.Sprintf("\n\t\t\t\t\t%s: true,", nameAndIDSlice[i].GetNameID())
	}
	return res
}
