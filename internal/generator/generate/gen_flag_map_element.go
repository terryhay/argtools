package generate

import (
	"argtools/internal/generator/configYaml"
	"argtools/internal/generator/idTemplateDataCreator"
	"fmt"
	"strings"
)

type FlagMapElements string

const (
	flagMapElementPrefix = `			%s: {
`
	flagMapElementDescriptionHelpInfo = `				DescriptionHelpInfo:  "%s",
`
	flagMapElementPostfix = `			},
`
)

func GenFlagMapElements(
	flagDescriptions []*configYaml.FlagDescription,
	flagsIDTemplateData map[configYaml.Flag]*idTemplateDataCreator.IDTemplateData) FlagMapElements {

	if len(flagDescriptions) == 0 {
		return "nil,"
	}

	builder := new(strings.Builder)
	builder.WriteString(`		map[argParserConfig.Flag]*argParserConfig.FlagDescription{`)

	var (
		flagDescription      *configYaml.FlagDescription
		argumentsDescription *configYaml.ArgumentsDescription
	)

	builder.WriteString("\n")
	for i := range flagDescriptions {
		flagDescription = flagDescriptions[i]

		builder.WriteString(fmt.Sprintf(flagMapElementPrefix, flagsIDTemplateData[flagDescription.GetFlag()].GetStringID()))
		builder.WriteString(fmt.Sprintf(flagMapElementDescriptionHelpInfo, flagDescription.GetDescriptionHelpInfo()))

		if argumentsDescription = flagDescription.GetArgumentsDescription(); argumentsDescription != nil {
			builder.WriteString(string(GenArgDescriptionElement(argumentsDescription)))
		}

		builder.WriteString(flagMapElementPostfix)
	}
	builder.WriteString(`		},`)

	return FlagMapElements(builder.String())
}
