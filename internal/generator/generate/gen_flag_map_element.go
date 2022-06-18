package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

type FlagDescriptionsSection string

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
	flagsIDTemplateData map[configYaml.Flag]*idTemplateDataCreator.IDTemplateData) FlagDescriptionsSection {

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

		builder.WriteString(fmt.Sprintf(flagMapElementPrefix, flagsIDTemplateData[flagDescription.GetFlag()].GetNameID()))
		builder.WriteString(fmt.Sprintf(flagMapElementDescriptionHelpInfo, flagDescription.GetDescriptionHelpInfo()))

		if argumentsDescription = flagDescription.GetArgumentsDescription(); argumentsDescription != nil {
			builder.WriteString(string(GenArgDescriptionElement(argumentsDescription, "\t\t\t\t", true)))
		}

		builder.WriteString(flagMapElementPostfix)
	}
	builder.WriteString(`		},`)

	return FlagDescriptionsSection(builder.String())
}
