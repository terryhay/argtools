package generate

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"strings"
)

const (
	appDescriptionPattern = `
		// appDescription
		argParserConfig.ApplicationDescription{
			AppName: "%s",
			NameHelpInfo: "%s",
			DescriptionHelpInfo: %s,
		}`
)

// AppDescriptionSection - string with application description paste section
type AppDescriptionSection string

// GenAppDescriptionSection creates a paste section with application description
func GenAppDescriptionSection(
	appDescription *configYaml.AppHelpDescription,
) AppDescriptionSection {

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

	return AppDescriptionSection(fmt.Sprintf(appDescriptionPattern,
		appDescription.GetApplicationName(),
		appDescription.GetNameHelpInfo(),
		descriptionHelpInfo))
}
