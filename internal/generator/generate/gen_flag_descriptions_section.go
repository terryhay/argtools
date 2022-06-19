package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

const (
	flagDescriptionsNilPart = `
		// flagDescriptions
		nil`

	flagDescriptionMapPrefix = `
		// flagDescriptions
		map[argParserConfig.Flag]*argParserConfig.FlagDescription{`
	flagDescriptionMapElementKeyPart = `
			%s: {`
	flagDescriptionMapElementDescriptionHelpInfo = `
				DescriptionHelpInfo:  "%s",`
	flagDescriptionMapElementPostfix = `
			},`
	flagDescriptionMapPostfix = `
		}`
)

// FlagDescriptionsSection - string with flag constant definitions list
type FlagDescriptionsSection string

// GenFlagDescriptionsSection creates a paste section with flag descriptions
func GenFlagDescriptionsSection(
	flagDescriptions []*configYaml.FlagDescription,
	flagsIDTemplateData map[string]*idTemplateDataCreator.IDTemplateData) FlagDescriptionsSection {

	if len(flagDescriptions) == 0 {
		return flagDescriptionsNilPart
	}

	builder := new(strings.Builder)
	builder.WriteString(flagDescriptionMapPrefix)

	var (
		flagDescription      *configYaml.FlagDescription
		argumentsDescription *configYaml.ArgumentsDescription
	)

	for i := 0; i < len(flagDescriptions); i++ {
		flagDescription = flagDescriptions[i]

		builder.WriteString(fmt.Sprintf(flagDescriptionMapElementKeyPart,
			flagsIDTemplateData[flagDescription.GetFlag()].GetNameID()))
		builder.WriteString(fmt.Sprintf(flagDescriptionMapElementDescriptionHelpInfo,
			flagDescription.GetDescriptionHelpInfo()))

		if argumentsDescription = flagDescription.GetArgumentsDescription(); argumentsDescription != nil {
			builder.WriteString(string(GenArgDescriptionPart(argumentsDescription, "\t\t\t\t", true)))
			builder.WriteString(",")
		}

		builder.WriteString(flagDescriptionMapElementPostfix)
	}
	builder.WriteString(flagDescriptionMapPostfix)

	return FlagDescriptionsSection(builder.String())
}
