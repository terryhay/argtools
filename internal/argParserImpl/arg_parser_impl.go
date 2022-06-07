package argParserImpl

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
)

type parseState uint8

const (
	// nolint
	parseStateNone parseState = iota
	parseStateReadingFlag
	parseStateReadingSingleArgument
	parseStateReadingArgumentArray
)

// ArgParserImpl - implementation of command line argument parser
type ArgParserImpl struct {
	nullCommandDescription *argParserConfig.CommandDescription
	commandDescriptions    map[argParserConfig.Command]*argParserConfig.CommandDescription
	flagDescriptions       map[argParserConfig.Flag]*argParserConfig.FlagDescription
}

// NewCmdArgParserImpl - ArgParserImpl object constructor
func NewCmdArgParserImpl(config argParserConfig.ArgParserConfig) (impl *ArgParserImpl) {
	commandDescriptions := make(
		map[argParserConfig.Command]*argParserConfig.CommandDescription,
		len(config.GetCommandDescriptions()))

	var command argParserConfig.Command
	for _, commandDescription := range config.GetCommandDescriptions() {
		for command = range commandDescription.GetCommands() {
			commandDescriptions[command] = commandDescription
		}
	}

	impl = &ArgParserImpl{
		commandDescriptions: commandDescriptions,
		flagDescriptions:    config.GetFlagDescriptions(),
	}
	if config.GetNullCommandDescription() != nil {
		impl.nullCommandDescription = &argParserConfig.CommandDescription{
			ID:                  config.GetNullCommandDescription().GetID(),
			DescriptionHelpInfo: config.GetNullCommandDescription().GetDescriptionHelpInfo(),
			ArgDescription:      config.GetNullCommandDescription().GetArgDescription(),
			RequiredFlags:       config.GetNullCommandDescription().GetRequiredFlags(),
			OptionalFlags:       config.GetNullCommandDescription().GetOptionalFlags(),
		}
	}

	return impl
}

// Parse - processes command line arguments
func (i *ArgParserImpl) Parse(args []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
	_ = i // check if pointer is nil

	if len(args) == 0 {
		if i.nullCommandDescription == nil {
			return nil,
				argtoolsError.NewError(argtoolsError.CodeArgParserNullCommandUndefined, fmt.Errorf(`ArgParserImpl: arguments are not set, but null command is not defined in config object`))
		}

		res = parsedData.NewParsedData(i.nullCommandDescription.GetID(), "", nil, nil)
		err = i.checkParsedData(i.nullCommandDescription, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	if len(i.commandDescriptions) == 0 && i.nullCommandDescription == nil {
		return nil,
			argtoolsError.NewError(argtoolsError.CodeArgParserIsNotInitialized, fmt.Errorf(`CmdArgParser: parser is not initialized`))
	}

	var (
		argIndexStartValue      = 1
		flag                    argParserConfig.Flag
		arg                     string
		contain                 bool
		commandArgData          *parsedData.ParsedArgData
		usingCommandDescription *argParserConfig.CommandDescription
		usingArgDescription     *argParserConfig.ArgumentsDescription
		flagDescription         *argParserConfig.FlagDescription
		parsedArgData           *parsedData.ParsedArgData
	)

	// Determinate command
	command := argParserConfig.Command(args[0])
	usingCommandDescription, contain = i.commandDescriptions[command]
	if !contain {
		if i.nullCommandDescription == nil {

			return nil,
				argtoolsError.NewError(argtoolsError.CodeCantFindFlagNameInGroupSpec,
					fmt.Errorf(`CmdArgParser: unexpected command: "%s"`, command))
		}
		usingCommandDescription = i.nullCommandDescription
		usingArgDescription = usingCommandDescription.GetArgDescription()
		command = ""
		argIndexStartValue = 0
	}

	if usingArgDescription = usingCommandDescription.GetArgDescription(); usingArgDescription != nil {
		commandArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
		parsedArgData = commandArgData
	}
	state := getParseState(usingArgDescription)

	parsedFlagDataMap := make(
		map[argParserConfig.Flag]*parsedData.ParsedFlagData,
		len(usingCommandDescription.GetRequiredFlags())+len(usingCommandDescription.GetOptionalFlags()))

	for argIndex := argIndexStartValue; argIndex < len(args); argIndex++ {
		arg = args[argIndex]

		switch state {
		case parseStateReadingSingleArgument:
			if !checkNoDashInFront(arg) {
				return nil, argtoolsError.NewError(argtoolsError.CodeArgParserDashInFrontOfArg,
					fmt.Errorf(`CmdArgParser: duplicate flag: "%s"`, arg))
			}

			parsedArgData.ArgValues = []parsedData.ArgValue{parsedData.ArgValue(arg)}

			state = parseStateReadingFlag

		case parseStateReadingFlag:
			flag = argParserConfig.Flag(arg)
			if flagDescription, contain = i.flagDescriptions[flag]; !contain {
				return nil, argtoolsError.NewError(argtoolsError.CodeArgParserUnexpectedArg,
					fmt.Errorf(`CmdArgParser: unexpected argument: "%s"`, arg))
			}

			parsedArgData = nil
			if usingArgDescription = flagDescription.GetArgDescription(); usingArgDescription != nil {
				parsedArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
			}

			if _, contain = parsedFlagDataMap[flag]; contain {
				return nil, argtoolsError.NewError(argtoolsError.CodeArgParserDuplicateFlags,
					fmt.Errorf(`CmdArgParser: duplicate flag: "%s"`, arg))
			}
			parsedFlagDataMap[flag] = parsedData.NewParsedFlagData(flag, parsedArgData)

			state = getParseState(usingArgDescription)

		case parseStateReadingArgumentArray:
			if !checkNoDashInFront(arg) {
				if len(parsedArgData.ArgValues) == 0 {
					return nil, argtoolsError.NewError(argtoolsError.CodeArgParserDashInFrontOfArg,
						fmt.Errorf(`CmdArgParser: duplicate flag: "%s"`, arg))
				}

				if flagDescription, contain = i.flagDescriptions[argParserConfig.Flag(arg)]; !contain {
					return nil, argtoolsError.NewError(argtoolsError.CodeArgParserUnexpectedFlag,
						fmt.Errorf(`CmdArgParser: unexpected flag: "%s"`, arg))
				}

				parsedArgData = nil
				if usingArgDescription = flagDescription.GetArgDescription(); usingArgDescription != nil {
					parsedArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
				}

				flag = argParserConfig.Flag(arg)
				if _, contain = parsedFlagDataMap[flag]; contain {
					return nil, argtoolsError.NewError(argtoolsError.CodeArgParserDuplicateFlags,
						fmt.Errorf(`CmdArgParser: duplicate flag: "%s"`, arg))
				}
				parsedFlagDataMap[flag] = parsedData.NewParsedFlagData(flag, parsedArgData)

				state = getParseState(usingArgDescription)
				continue
			}

			parsedArgData.ArgValues = append(parsedArgData.ArgValues, parsedData.ArgValue(arg))
		}
	}

	res = parsedData.NewParsedData(usingCommandDescription.GetID(), command, commandArgData, parsedFlagDataMap)
	err = i.checkParsedData(usingCommandDescription, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getParseState(argumentsDescription *argParserConfig.ArgumentsDescription) parseState {
	switch argumentsDescription.GetAmountType() {
	case argParserConfig.ArgAmountTypeSingle:
		return parseStateReadingSingleArgument
	case argParserConfig.ArgAmountTypeList:
		return parseStateReadingArgumentArray
	default:
		return parseStateReadingFlag
	}
}

func checkNoDashInFront(arg string) bool {
	return arg[:1] != "-"
}

func (i *ArgParserImpl) checkParsedData(
	usingCommandDescription *argParserConfig.CommandDescription,
	data *parsedData.ParsedData,
) *argtoolsError.Error {
	var (
		argDescription *argParserConfig.ArgumentsDescription
		contain        bool
		flag           argParserConfig.Flag
		parsedFlagData = data.GetFlagData()
	)

	// check if all required flags is set
	for flag = range usingCommandDescription.GetRequiredFlags() {
		if _, contain = parsedFlagData[flag]; !contain {
			return argtoolsError.NewError(
				argtoolsError.CodeArgParserRequiredFlagIsNotSet,
				fmt.Errorf("CmdArgParser.checkParsedData: required flag is not set: %s", flag))
		}
	}

	// check command arguments
	argDescription = usingCommandDescription.GetArgDescription()
	if argDescription.GetAmountType() != argParserConfig.ArgAmountTypeNoArgs {
		if data.GetAgrData() == nil {
			return argtoolsError.NewError(
				argtoolsError.CodeArgParserCommandDoesNotContainArgs,
				fmt.Errorf("CmdArgParser.checkParsedData: required flag is not set: %s", flag))
		}
	}

	return nil
}
