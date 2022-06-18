package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
)

type HelpCommandDescriptionSection string

const (
	helpCommandComponentPattern = `		argParserConfig.NewHelpCommandDescription(
			CommandIDPrintHelpInfo,
			map[argParserConfig.Command]bool{
				%s
			},
		),`
)

func GenHelpCommandComponent(helpCommandDescription *configYaml.HelpCommandDescription) HelpCommandDescriptionSection {
	if helpCommandDescription == nil {
		return "nil"
	}

	commandList := fmt.Sprintf("\"%s\": true,", helpCommandDescription.GetCommand())
	for _, command := range helpCommandDescription.GetAdditionalCommands() {
		commandList += fmt.Sprintf("\n\t\t\t\t\"%s\": true,", command)
	}

	return HelpCommandDescriptionSection(fmt.Sprintf(helpCommandComponentPattern,
		commandList,
	))
}
