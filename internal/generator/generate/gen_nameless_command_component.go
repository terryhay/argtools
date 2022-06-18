package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

type NamelessCommandDescriptionSection string

func GenNamelessCommandComponent(
	namelessCommandDescription *configYaml.NamelessCommandDescription,
	namelessCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
	flagsIDTemplateData map[configYaml.Flag]*idTemplateDataCreator.IDTemplateData,
) NamelessCommandDescriptionSection {

	if namelessCommandDescription == nil {
		return "\t\tnil"
	}

	var (
		builder strings.Builder
		flag    configYaml.Flag
	)

	builder.WriteString(fmt.Sprintf(`		argParserConfig.NewNamelessCommandDescription(
			%s,
			"%s",
`,
		namelessCommandIDTemplateData.GetID(),
		namelessCommandDescription.GetDescriptionHelpInfo()))

	if namelessCommandDescription.GetArgumentsDescription() == nil {
		builder.WriteString("nil,\n")
	} else {
		builder.WriteString(string(GenArgDescriptionElement(namelessCommandDescription.GetArgumentsDescription(), "\t\t\t", false)))
	}

	if len(namelessCommandDescription.GetRequiredFlags()) == 0 {
		builder.WriteString("nil,\n")
	} else {
		builder.WriteString("\t\t\tmap[argParserConfig.Flag]bool{\n")
		for _, flag = range namelessCommandDescription.GetRequiredFlags() {
			builder.WriteString(fmt.Sprintf("\t\t\t\t%s: true,\n", flagsIDTemplateData[flag].GetNameID()))
		}
		builder.WriteString("\t\t\t},\n")
	}

	if len(namelessCommandDescription.GetOptionalFlags()) == 0 {
		builder.WriteString("nil,\n")
	} else {
		builder.WriteString("\t\t\tmap[argParserConfig.Flag]bool{\n")
		for _, flag = range namelessCommandDescription.GetOptionalFlags() {
			builder.WriteString(fmt.Sprintf("\t\t\t\t%s: true,\n", flagsIDTemplateData[flag].GetNameID()))
		}
		builder.WriteString("\t\t\t},\n")
	}

	builder.WriteString(`		)`)
	return NamelessCommandDescriptionSection(builder.String())
}
