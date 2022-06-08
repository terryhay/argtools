package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"strings"
)

const appDescriptionTemplate = `			AppName: "%s",
			NameHelpInfo: "%s",
			DescriptionHelpInfo: %s,
`

type AppDescriptionComponent string

func GenAppDescription(appDescription *configYaml.AppHelpDescription) AppDescriptionComponent {
	descriptionHelpInfo := "nil"
	if len(appDescription.GetDescriptionHelpInfo()) > 0 {
		builder := strings.Builder{}
		builder.WriteString("[]string{")
		for _, paragraph := range appDescription.GetDescriptionHelpInfo() {
			builder.WriteString(fmt.Sprintf("\n\t\t\t\t\"%s\",", paragraph))
		}
		builder.WriteString("\n\t\t\t}")
		descriptionHelpInfo = builder.String()
	}
	return AppDescriptionComponent(fmt.Sprintf(appDescriptionTemplate,
		appDescription.GetApplicationName(),
		appDescription.GetNameHelpInfo(),
		descriptionHelpInfo))
}
