package argTools

import (
	"github.com/terryhay/argtools/internal/argParserImpl"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/helpPrinter"
	"github.com/terryhay/argtools/pkg/parsedData"
)

const (
	// CommandIDNullCommand - checks arguments types
	CommandIDNullCommand argParserConfig.CommandID = iota + 1
	// CommandIDHelp - print help info
	CommandIDHelp
)

const (
	// CommandH - print help info
	CommandH argParserConfig.Command = "-h"
	// CommandHelp - print help info
	CommandHelp = "help"
)

const (
	// FlagC - yaml file config path
	FlagC argParserConfig.Flag = "-c"
	// FlagO - generate package path
	FlagO = "-o"
)

// Parse - processes command line arguments
func Parse(args []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
	appArgConfig := argParserConfig.NewArgParserConfig(
		// appDescription
		argParserConfig.ApplicationDescription{
			AppName:      "gen_argtools",
			NameHelpInfo: "code generator",
			DescriptionHelpInfo: []string{
				"generate argTools package which contains a command line data parser",
			},
		},
		// flagDescriptions
		map[argParserConfig.Flag]*argParserConfig.FlagDescription{
			FlagC: {
				DescriptionHelpInfo: "yaml file config path",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeSingle,
					SynopsisHelpDescription: "file",
				},
			},
			FlagO: {
				DescriptionHelpInfo: "generate package path",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeSingle,
					SynopsisHelpDescription: "dir",
				},
			},
		},
		// commandDescriptions
		nil,
		// nullCommandDescription
		&argParserConfig.NullCommandDescription{
			ID:                  CommandIDNullCommand,
			DescriptionHelpInfo: "generate argTools package",
			RequiredFlags: map[argParserConfig.Flag]bool{
				FlagC: true,
				FlagO: true,
			},
		},
	)

	if res, err = argParserImpl.NewCmdArgParserImpl(appArgConfig).Parse(args); err != nil {
		return nil, err
	}

	if res.GetCommandID() == CommandIDHelp {
		helpPrinter.PrintHelpInfo(appArgConfig)
		return nil, nil
	}

	return res, nil
}
