package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"strings"
)

const (
	helpCommandDescriptionSectionNilPattern = `
		// helpCommandDescription
		nil`
	helpCommandDescriptionSectionPattern = `
		// helpCommandDescription
		argParserConfig.NewHelpCommandDescription(
			CommandIDPrintHelpInfo,
			map[argParserConfig.Command]bool{
				%s
			},
		)`
)

// HelpCommandDescriptionSection - string with help command description section
type HelpCommandDescriptionSection string

// GenHelpCommandDescriptionSection creates a paste section with help command description
func GenHelpCommandDescriptionSection(
	helpCommandDescription *configYaml.HelpCommandDescription,
) HelpCommandDescriptionSection {

	if helpCommandDescription == nil {
		return helpCommandDescriptionSectionNilPattern
	}

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("\"%s\": true,", helpCommandDescription.GetCommand()))
	for _, command := range helpCommandDescription.GetAdditionalCommands() {
		builder.WriteString(fmt.Sprintf("\n\t\t\t\t\"%s\": true,", command))
	}

	return HelpCommandDescriptionSection(fmt.Sprintf(helpCommandDescriptionSectionPattern,
		builder.String(),
	))
}
