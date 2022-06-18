package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

const (
	namelessCommandDescriptionNilPart = `
		// namelessCommandDescription
		nil`

	namelessCommandDescriptionSectionPattern = `
		// namelessCommandDescription
		argParserConfig.NewNamelessCommandDescription(
			%s,
			"%s",
			%s,
			%s,
			%s,
		)`
)

// NamelessCommandDescriptionSection - string with nameless command description section
type NamelessCommandDescriptionSection string

// GenNamelessCommandComponent creates a paste section with nameless command description
func GenNamelessCommandComponent(
	namelessCommandDescription *configYaml.NamelessCommandDescription,
	namelessCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
	flagsIDTemplateData map[string]*idTemplateDataCreator.IDTemplateData,
) NamelessCommandDescriptionSection {

	if namelessCommandDescription == nil {
		return namelessCommandDescriptionNilPart
	}

	var (
		builder strings.Builder
		flag    string
	)

	argumentsDescriptionPart := "nil"
	if namelessCommandDescription.GetArgumentsDescription() != nil {
		argumentsDescriptionPart = string(GenArgDescriptionPart(namelessCommandDescription.GetArgumentsDescription(), "\t\t\t", false))
	}

	requiredFlagsPart := "nil"
	if len(namelessCommandDescription.GetRequiredFlags()) != 0 {
		builder.WriteString("\t\t\tmap[argParserConfig.Flag]bool{\n")
		for _, flag = range namelessCommandDescription.GetRequiredFlags() {
			builder.WriteString(fmt.Sprintf("\t\t\t\t%s: true,\n", flagsIDTemplateData[flag].GetNameID()))
		}
		builder.WriteString("\t\t\t}")
		requiredFlagsPart = builder.String()
	}

	optionalFlagsPart := "nil"
	if len(namelessCommandDescription.GetOptionalFlags()) != 0 {
		builder.Reset()

		builder.WriteString("\t\t\tmap[argParserConfig.Flag]bool{\n")
		for _, flag = range namelessCommandDescription.GetOptionalFlags() {
			builder.WriteString(fmt.Sprintf("\t\t\t\t%s: true,\n", flagsIDTemplateData[flag].GetNameID()))
		}
		builder.WriteString("\t\t\t}")
		optionalFlagsPart = builder.String()
	}

	return NamelessCommandDescriptionSection(fmt.Sprintf(namelessCommandDescriptionSectionPattern,
		namelessCommandIDTemplateData.GetID(),
		namelessCommandDescription.GetDescriptionHelpInfo(),
		argumentsDescriptionPart,
		requiredFlagsPart,
		optionalFlagsPart,
	))
}
