package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

type NamelessCommandComponent string

func GenNamelessCommandComponent(
	namelessCommandDescription *configYaml.NamelessCommandDescription,
	namelessCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
) NamelessCommandComponent {

	if namelessCommandDescription == nil {
		return "\t\tnil"
	}

	builder := new(strings.Builder)

	builder.WriteString(fmt.Sprintf(`		&argParserConfig.NamelessCommandDescription{
			ID: %s,
			DescriptionHelpInfo: "%s",`,
		namelessCommandIDTemplateData.GetID(),
		namelessCommandDescription.GetDescriptionHelpInfo()))

	if namelessCommandDescription.GetArgumentsDescription() != nil {
		builder.WriteString(fmt.Sprintf("\n%s", GenArgDescriptionElement(namelessCommandDescription.GetArgumentsDescription(), "\t\t\t")))
	} else {
		builder.WriteString("\n")
	}

	builder.WriteString(`		},
`)
	return NamelessCommandComponent(builder.String())
}
