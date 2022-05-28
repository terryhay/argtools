package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/idTemplateDataCreator"
	"strings"
)

type NullCommandComponent string

func GenNullCommandComponent(
	nullCommandDescription *configYaml.NullCommandDescription,
	nullCommandIDTemplateData *idTemplateDataCreator.IDTemplateData,
) NullCommandComponent {

	if nullCommandDescription == nil {
		return "\t\tnil"
	}

	builder := new(strings.Builder)

	builder.WriteString(fmt.Sprintf(`		&argParserConfig.NullCommandDescription{
			ID: %s,
			DescriptionHelpInfo: "%s",`,
		nullCommandIDTemplateData.GetID(),
		nullCommandDescription.GetDescriptionHelpInfo()))

	if nullCommandDescription.GetArgumentsDescription() != nil {
		builder.WriteString(fmt.Sprintf("\n%s", GenArgDescriptionElement(nullCommandDescription.GetArgumentsDescription())))
	}

	builder.WriteString(`		},
`)
	return NullCommandComponent(builder.String())
}
