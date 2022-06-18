package generate

import (
	"fmt"
)

const argParserFileTemplate = `// This code was generated by argtools.generator. DO NOT EDIT

package argTools

import (
	"github.com/terryhay/argtools/pkg/argParser"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/helpPrinter"
	"github.com/terryhay/argtools/pkg/parsedData"
)

%s

%s

%s// Parse - processes command line arguments
func Parse(args []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
	appArgConfig := argParserConfig.NewArgParserConfig(
		// appDescription
		argParserConfig.ApplicationDescription{
%s		},
		// flagDescriptions
%s
		// commandDescriptions
%s
		// helpCommandDescription
%s
		// namelessCommandDescription
%s)

	if res, err = argParser.Parse(appArgConfig, args); err != nil {
		return nil, err
	}

	if res.GetCommandID() == %s {
		helpPrinter.PrintHelpInfo(appArgConfig)
		return nil, nil
	}

	return res, nil
}
`

// GenArgParserFileBody applies data to argParserFileTemplate
func GenArgParserFileBody(
	commandIDListSection CommandIDListSection,
	commandNameIDListSection CommandListSection,
	flagStringIDListSection FlagStringIDListSection,
	appDescriptionSection AppDescriptionSection,
	flagDescriptionsSection FlagDescriptionsSection,
	commandDescriptionsSection CommandDescriptionsSection,
	helpCommandDescriptionSection HelpCommandDescriptionSection,
	namelessCommandDescriptionSection NamelessCommandDescriptionSection,
	helpCommandID string) string {

	return fmt.Sprintf(
		argParserFileTemplate,

		commandIDListSection,
		commandNameIDListSection,
		flagStringIDListSection,

		appDescriptionSection,
		flagDescriptionsSection,
		commandDescriptionsSection,
		helpCommandDescriptionSection,
		namelessCommandDescriptionSection,

		helpCommandID,
	)
}
