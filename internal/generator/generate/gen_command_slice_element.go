package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

type commandSliceElement string

const (
	commandSliceElementPrefix = `			{
`
	commandSliceElementIDTemplate = `				ID: %s,
`
	commandSliceElementCommandsTemplate = `				Commands: map[argParserConfig.Command]bool{%s
				},
`
	commandSliceElementRequiredFlagsTemplate = `				RequiredFlags: map[argParserConfig.Flag]bool{%s
				},
`
	commandSliceElementOptionalFlagsTemplate = `				OptionalFlags: map[argParserConfig.Flag]bool{%s
				},
`
	commandSliceElementPostfix = `			},
`
)

func GenCommandSliceElements(
	commandDescriptions []*configYaml.CommandDescription,
	helpCommandDescription *configYaml.HelpCommandDescription,
	commandsIDTemplateData map[configYaml.Command]*idTemplateDataCreator.IDTemplateData,
	flagsIDTemplateData map[configYaml.Flag]*idTemplateDataCreator.IDTemplateData) commandSliceElement {

	builder := new(strings.Builder)
	builder.WriteString(`		[]*argParserConfig.CommandDescription{
`)

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

		builder.WriteString(commandSliceElementPrefix)
		builder.WriteString(fmt.Sprintf(commandSliceElementIDTemplate, commandsIDTemplateData[commandDescription.GetCommand()].GetID()))
		builder.WriteString(fmt.Sprintf(commandSliceElementCommandsTemplate, joinCallNames(idTemplateDataSlice)))
		idTemplateDataSlice = []*idTemplateDataCreator.IDTemplateData{}

		if len(commandDescription.GetRequiredFlags()) > 0 {
			for j = range commandDescription.GetRequiredFlags() {
				idTemplateDataSlice = append(idTemplateDataSlice, flagsIDTemplateData[commandDescription.GetRequiredFlags()[j]])
			}
			builder.WriteString(fmt.Sprintf(commandSliceElementRequiredFlagsTemplate, joinCallNames(idTemplateDataSlice)))
			idTemplateDataSlice = []*idTemplateDataCreator.IDTemplateData{}
		}

		if len(commandDescription.GetOptionalFlags()) > 0 {
			for j = range commandDescription.GetOptionalFlags() {
				idTemplateDataSlice = append(idTemplateDataSlice, flagsIDTemplateData[commandDescription.GetOptionalFlags()[j]])
			}
			builder.WriteString(fmt.Sprintf(commandSliceElementOptionalFlagsTemplate, joinCallNames(idTemplateDataSlice)))
			idTemplateDataSlice = []*idTemplateDataCreator.IDTemplateData{}
		}

		builder.WriteString(commandSliceElementPostfix)
	}

	idTemplateDataSlice = append(idTemplateDataSlice, commandsIDTemplateData[helpCommandDescription.GetCommand()])
	for j = range helpCommandDescription.GetAdditionalCommands() {
		idTemplateDataSlice = append(idTemplateDataSlice, commandsIDTemplateData[helpCommandDescription.GetAdditionalCommands()[j]])
	}
	builder.WriteString(commandSliceElementPrefix)
	builder.WriteString(fmt.Sprintf(commandSliceElementIDTemplate, commandsIDTemplateData[helpCommandDescription.GetCommand()].GetID()))
	builder.WriteString(fmt.Sprintf(commandSliceElementCommandsTemplate, joinCallNames(idTemplateDataSlice)))
	builder.WriteString(commandSliceElementPostfix)

	builder.WriteString(`		},`)

	return commandSliceElement(builder.String())
}

func joinCallNames(nameAndIDSlice []*idTemplateDataCreator.IDTemplateData) (res string) {
	for i := range nameAndIDSlice {
		res += fmt.Sprintf("\n\t\t\t\t\t%s: true,", nameAndIDSlice[i].GetStringID())
	}

	return res
}
