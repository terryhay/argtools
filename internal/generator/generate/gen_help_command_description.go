package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
)

type HelpCommandComponent string

const (
	helpCommandComponentPattern = `		argParserConfig.NewHelpCommandDescription(
			CommandIDPrintHelpInfo,
			map[argParserConfig.Command]bool{
				%s
			},
		),`
)

func GenHelpCommandComponent(helpCommandDescription *configYaml.HelpCommandDescription) HelpCommandComponent {
	if helpCommandDescription == nil {
		return "nil"
	}

	commandList := fmt.Sprintf("\"%s\": true,", helpCommandDescription.GetCommand())
	for _, command := range helpCommandDescription.GetAdditionalCommands() {
		commandList += fmt.Sprintf("\n\t\t\t\t\"%s\": true,", command)
	}

	return HelpCommandComponent(fmt.Sprintf(helpCommandComponentPattern,
		commandList,
	))
}
