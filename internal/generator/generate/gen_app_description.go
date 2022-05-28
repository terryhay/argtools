package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
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
