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
	nullCommandDescription *argParserConfig.NullCommandDescription
	commandDescriptions    map[argParserConfig.Command]*argParserConfig.CommandDescription
	flagDescriptions       map[argParserConfig.Flag]*argParserConfig.FlagDescription
}

// NewCmdArgParserImpl - ArgParserImpl object constructor
func NewCmdArgParserImpl(config *argParserConfig.ArgParserConfig) (impl *ArgParserImpl) {
	commandDescriptions := make(
		map[argParserConfig.Command]*argParserConfig.CommandDescription,
		len(config.GetCommandDescriptions()))

	var command argParserConfig.Command
	for _, commandDescription := range config.GetCommandDescriptions() {
		for command = range commandDescription.GetCommands() {
			commandDescriptions[command] = commandDescription
		}
	}

	return &ArgParserImpl{
		nullCommandDescription: config.NullCommandDescription,
		commandDescriptions:    commandDescriptions,
		flagDescriptions:       config.GetFlagDescriptions(),
	}
}

// Parse - processes command line arguments
func (i *ArgParserImpl) Parse(args []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
	_ = i // check if pointer is nil

	if len(args) == 0 {
		return nil, nil
	}
	if len(i.commandDescriptions) == 0 && i.nullCommandDescription == nil {
		return nil,
			argtoolsError.NewError(argtoolsError.CodeParserIsNotInitialized, fmt.Errorf(`CmdArgParser: parser is not initialized`))
	}

	var (
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
		return nil,
			argtoolsError.NewError(argtoolsError.CodeCantFindFlagNameInGroupSpec,
				fmt.Errorf(`CmdArgParser: unexpected command: "%s"`, command))
	}

	if usingArgDescription = usingCommandDescription.GetArgDescription(); usingArgDescription != nil {
		commandArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
		parsedArgData = commandArgData
	}
	state := getParseState(usingArgDescription)

	parsedFlagDataMap := make(
		map[argParserConfig.Flag]*parsedData.ParsedFlagData,
		len(usingCommandDescription.GetRequiredFlags())+len(usingCommandDescription.GetOptionalFlags()))

	for argIndex := 1; argIndex < len(args); argIndex++ {
		arg = args[argIndex]

		switch state {
		case parseStateReadingSingleArgument:
			parsedArgData.ArgValues = []parsedData.ArgValue{parsedData.ArgValue(arg)}

			state = parseStateReadingFlag

		case parseStateReadingFlag:
			flag = argParserConfig.Flag(arg)
			if flagDescription, contain = i.flagDescriptions[flag]; !contain {
				return nil, nil //todo: error
			}

			parsedArgData = nil
			if usingArgDescription = flagDescription.GetArgDescription(); usingArgDescription != nil {
				parsedArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
			}

			if _, contain = parsedFlagDataMap[flag]; contain {
				// todo: return error
				panic("duplicate")
			}
			parsedFlagDataMap[flag] = parsedData.NewParsedFlagData(flag, parsedArgData)

			state = getParseState(usingArgDescription)

		case parseStateReadingArgumentArray:
			if flagDescription, contain = i.flagDescriptions[argParserConfig.Flag(arg)]; contain {
				parsedArgData = nil
				if usingArgDescription = flagDescription.GetArgDescription(); usingArgDescription != nil {
					parsedArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
				}

				if _, contain = parsedFlagDataMap[flag]; contain {
					// todo: return error
					panic("duplicate")
				}
				parsedFlagDataMap[flag] = parsedData.NewParsedFlagData(flag, parsedArgData)

				state = getParseState(usingArgDescription)
				continue
			}

			parsedArgData.ArgValues = append(parsedArgData.ArgValues, parsedData.ArgValue(arg))

		default:
			return nil,
				argtoolsError.NewError(argtoolsError.CodeUndefinedError,
					fmt.Errorf("CmdArgParser: internal error: unexpected state: %d", state))
		}
	}

	return parsedData.NewParsedData(
			usingCommandDescription.GetID(),
			command,
			commandArgData,
			parsedFlagDataMap),
		i.checkFlags(usingCommandDescription, parsedFlagDataMap)
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

func (i *ArgParserImpl) checkFlags(
	usingCommandDescription *argParserConfig.CommandDescription,
	parsedFlagDataMap map[argParserConfig.Flag]*parsedData.ParsedFlagData,
) *argtoolsError.Error {
	_ = i // check if pointer is nil

	var (
		contain bool
		flag    argParserConfig.Flag
	)

	// check if all required flags is set
	for flag = range usingCommandDescription.GetRequiredFlags() {
		if _, contain = parsedFlagDataMap[flag]; !contain {
			return argtoolsError.NewError(argtoolsError.CodeRequiredFlagIsNotSet,
				fmt.Errorf("CmdArgParser: required flag is not set: %s", flag))
		}
	}

	return nil
}
