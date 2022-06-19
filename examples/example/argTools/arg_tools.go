// This code was generated by argtools.generator. DO NOT EDIT

package argTools

import (
	"github.com/terryhay/argtools/pkg/argParser"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/helpPrinter"
	"github.com/terryhay/argtools/pkg/parsedData"
)

const (
	// CommandIDNamelessCommand - checks arguments types
	CommandIDNamelessCommand argParserConfig.CommandID = iota + 1
	// CommandIDPrintHelpInfo - print help info
	CommandIDPrintHelpInfo
	// CommandIDPrint - print command line arguments with optional checking
	CommandIDPrint
)

const (
	// CommandH - print help info
	CommandH argParserConfig.Command = "-h"
	// CommandHelp - print help info
	CommandHelp = "help"
	// CommandPrint - print command line arguments with optional checking
	CommandPrint = "print"
)

const (
	// FlagCheck - check command arguments types
	FlagCheck argParserConfig.Flag = "-check"
	// FlagCheckargs - do arguments checking
	FlagCheckargs = "-checkargs"
	// FlagF - single float
	FlagF = "-f"
	// FlagFl - float list
	FlagFl = "-fl"
	// FlagI - int string
	FlagI = "-i"
	// FlagIl - int list
	FlagIl = "-il"
	// FlagS - single string
	FlagS = "-s"
	// FlagSl - string list
	FlagSl = "-sl"
)

// Parse - processes command line arguments
func Parse(args []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
	appArgConfig := argParserConfig.NewArgParserConfig(
		// appDescription
		argParserConfig.ApplicationDescription{
			AppName:      "example",
			NameHelpInfo: "shows how argtools generator works",
			DescriptionHelpInfo: []string{
				"you can write more detailed description here",
				"and use several paragraphs",
			},
		},
		// flagDescriptions
		map[argParserConfig.Flag]*argParserConfig.FlagDescription{
			FlagCheck: {
				DescriptionHelpInfo: "check command arguments types",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeSingle,
					SynopsisHelpDescription: "str",
				},
			},
			FlagS: {
				DescriptionHelpInfo: "single string",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeSingle,
					SynopsisHelpDescription: "str",
				},
			},
			FlagSl: {
				DescriptionHelpInfo: "string list",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeList,
					SynopsisHelpDescription: "str",
				},
			},
			FlagI: {
				DescriptionHelpInfo: "int string",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeSingle,
					SynopsisHelpDescription: "str",
				},
			},
			FlagIl: {
				DescriptionHelpInfo: "int list",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeList,
					SynopsisHelpDescription: "str",
				},
			},
			FlagF: {
				DescriptionHelpInfo: "single float",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeSingle,
					SynopsisHelpDescription: "str",
				},
			},
			FlagFl: {
				DescriptionHelpInfo: "float list",
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType:              argParserConfig.ArgAmountTypeList,
					SynopsisHelpDescription: "str",
				},
			},
			FlagCheckargs: {
				DescriptionHelpInfo: "do arguments checking",
			},
		},
		// commandDescriptions
		[]*argParserConfig.CommandDescription{
			{
				ID:                  CommandIDPrint,
				DescriptionHelpInfo: "print command line arguments with optional checking",
				Commands: map[argParserConfig.Command]bool{
					CommandPrint: true,
				},
				OptionalFlags: map[argParserConfig.Flag]bool{
					FlagS:         true,
					FlagSl:        true,
					FlagI:         true,
					FlagIl:        true,
					FlagF:         true,
					FlagFl:        true,
					FlagCheckargs: true,
				},
			},
		},
		// helpCommandDescription
		argParserConfig.NewHelpCommandDescription(
			CommandIDPrintHelpInfo,
			map[argParserConfig.Command]bool{
				CommandH:    true,
				CommandHelp: true,
			},
		),
		// namelessCommandDescription
		argParserConfig.NewNamelessCommandDescription(
			CommandIDNamelessCommand,
			"checks arguments types",
			&argParserConfig.ArgumentsDescription{
				AmountType:              argParserConfig.ArgAmountTypeList,
				SynopsisHelpDescription: "str list",
			},
			map[argParserConfig.Flag]bool{
				FlagCheck: true,
			},
			nil,
		))

	if res, err = argParser.Parse(appArgConfig, args); err != nil {
		return nil, err
	}

	if res.GetCommandID() == CommandIDPrintHelpInfo {
		helpPrinter.PrintHelpInfo(appArgConfig)
		return nil, nil
	}

	return res, nil
}
