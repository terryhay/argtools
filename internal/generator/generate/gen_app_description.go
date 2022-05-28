package generate

import (
	"argtools/internal/generator/configYaml"
	"fmt"
)

const appDescriptionTemplate = `			AppName: "%s",
			NameHelpInfo: "%s",
			DescriptionHelpInfo: "%s",
`

type AppDescriptionComponent string

func GenAppDescription(appDescription *configYaml.AppHelpDescription) AppDescriptionComponent {
	return AppDescriptionComponent(fmt.Sprintf(appDescriptionTemplate,
		appDescription.GetApplicationName(),
		appDescription.GetNameHelpInfo(),
		appDescription.GetDescriptionHelpInfo()))
}
