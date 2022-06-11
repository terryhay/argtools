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
	parseStateReadingFlag parseState = iota
	parseStateReadingSingleArgument
	parseStateReadingArgumentList
)

const countOfParseStates = int(parseStateReadingArgumentList) + 1

// ArgParserImpl - implementation of command line argument parser
type ArgParserImpl struct {
	namelessCommandDescription *argParserConfig.CommandDescription
	commandDescriptions        map[argParserConfig.Command]*argParserConfig.CommandDescription
	flagDescriptions           map[argParserConfig.Flag]*argParserConfig.FlagDescription
	stateProcessors            [countOfParseStates]func(arg string) *argtoolsError.Error

	// mutable data
	mArgDescription     *argParserConfig.ArgumentsDescription
	mErr                *argtoolsError.Error
	mFlag               argParserConfig.Flag
	mFlagDescription    *argParserConfig.FlagDescription
	mTmpFlagDescription *argParserConfig.FlagDescription
	mOK                 bool
	mParsedArgData      *parsedData.ParsedArgData
	mParseState         parseState

	// res data
	rParsedFlagDataMap map[argParserConfig.Flag]*parsedData.ParsedFlagData
}

// NewCmdArgParserImpl - ArgParserImpl object constructor
func NewCmdArgParserImpl(config argParserConfig.ArgParserConfig) *ArgParserImpl {
	commandDescriptions := make(
		map[argParserConfig.Command]*argParserConfig.CommandDescription,
		len(config.GetCommandDescriptions()))

	var command argParserConfig.Command
	for _, commandDescription := range config.GetCommandDescriptions() {
		for command = range commandDescription.GetCommands() {
			commandDescriptions[command] = commandDescription
		}
	}

	res := &ArgParserImpl{
		namelessCommandDescription: nameless2commandDescription(config.GetNamelessCommandDescription()),
		commandDescriptions:        commandDescriptions,
		flagDescriptions:           config.GetFlagDescriptions(),
	}
	res.stateProcessors = [countOfParseStates]func(arg string) *argtoolsError.Error{
		parseStateReadingFlag:           res.processReadingFlag,
		parseStateReadingSingleArgument: res.processReadingSingleArgument,
		parseStateReadingArgumentList:   res.processReadingArgumentList,
	}

	return res
}

// Parse - processes command line arguments
func (i *ArgParserImpl) Parse(args []string) (*parsedData.ParsedData, *argtoolsError.Error) {
	_ = i // check if pointer is nil here and check no further

	var (
		argIndexStartValue      = 1
		commandArgData          *parsedData.ParsedArgData
		res                     *parsedData.ParsedData
		usingCommandDescription *argParserConfig.CommandDescription
	)

	if len(args) == 0 {
		if i.namelessCommandDescription == nil {
			return nil, argtoolsError.NewError(
				argtoolsError.CodeArgParserNamelessCommandUndefined,
				fmt.Errorf(`ArgParserImpl.Parse: arguments are not set, but nameless command is not defined in config object`))
		}

		res = parsedData.NewParsedData(i.namelessCommandDescription.GetID(), "", nil, nil)
		if i.mErr = checkParsedData(i.namelessCommandDescription, res); i.mErr != nil {
			return nil, i.mErr
		}
		return res, nil
	}
	if len(i.commandDescriptions) == 0 && i.namelessCommandDescription == nil {
		return nil,
			argtoolsError.NewError(
				argtoolsError.CodeArgParserIsNotInitialized,
				fmt.Errorf(`ArgParserImpl.Parse: parser is not initialized`))
	}

	// Determinate command
	command := argParserConfig.Command(args[0])
	usingCommandDescription, i.mOK = i.commandDescriptions[command]
	if !i.mOK {
		if i.namelessCommandDescription == nil {
			return nil,
				argtoolsError.NewError(
					argtoolsError.CodeCantFindFlagNameInGroupSpec,
					fmt.Errorf(`ArgParserImpl.Parse: unexpected command: "%s"`, command))
		}
		usingCommandDescription = i.namelessCommandDescription
		command = ""
		argIndexStartValue = 0
	}

	if i.mArgDescription = usingCommandDescription.GetArgDescription(); i.mArgDescription != nil {
		commandArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
		i.mParsedArgData = commandArgData
	}
	i.mParseState = getParseState(i.mArgDescription)

	i.rParsedFlagDataMap = make(
		map[argParserConfig.Flag]*parsedData.ParsedFlagData,
		len(usingCommandDescription.GetRequiredFlags())+len(usingCommandDescription.GetOptionalFlags()))

	for argIndex := argIndexStartValue; argIndex < len(args); argIndex++ {
		if i.mErr = i.stateProcessors[i.mParseState](args[argIndex]); i.mErr != nil {
			return nil, i.mErr
		}
	}

	res = parsedData.NewParsedData(usingCommandDescription.GetID(), command, commandArgData, i.rParsedFlagDataMap)
	if i.mErr = checkParsedData(usingCommandDescription, res); i.mErr != nil {
		return nil, i.mErr
	}

	return res, nil
}

func (i *ArgParserImpl) processReadingFlag(arg string) *argtoolsError.Error {
	_ = i // check if pointer is nil here and check no further

	i.mFlag = argParserConfig.Flag(arg)
	if i.mFlagDescription, i.mOK = i.flagDescriptions[i.mFlag]; !i.mOK {
		return argtoolsError.NewError(
			argtoolsError.CodeArgParserUnexpectedArg,
			fmt.Errorf(`ArgParserImpl.Parse: unexpected argument: "%s"`, arg))
	}

	i.mParsedArgData = nil
	if i.mArgDescription = i.mFlagDescription.GetArgDescription(); i.mArgDescription != nil {
		i.mParsedArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
	}

	if _, i.mOK = i.rParsedFlagDataMap[i.mFlag]; i.mOK {
		return argtoolsError.NewError(
			argtoolsError.CodeArgParserDuplicateFlags,
			fmt.Errorf(`ArgParserImpl.Parse: duplicate flag: "%s"`, arg))
	}
	i.rParsedFlagDataMap[i.mFlag] = parsedData.NewParsedFlagData(i.mFlag, i.mParsedArgData)

	i.mParseState = getParseState(i.mArgDescription)
	return nil
}

func (i *ArgParserImpl) processReadingSingleArgument(arg string) *argtoolsError.Error {
	_ = i // check if pointer is nil here and check no further

	if !checkNoDashInFront(arg) {
		return i.notSetArgValueCase(arg)
	}

	if i.mErr = isValueAllowed(i.mArgDescription, arg); i.mErr != nil {
		return argtoolsError.NewError(
			i.mErr.Code(),
			fmt.Errorf(`ArgParserImpl.Parse: flag "%s": %s`, i.mFlag, i.mErr.Error()))
	}

	i.mParsedArgData.ArgValues = []parsedData.ArgValue{parsedData.ArgValue(arg)}

	i.mParseState = parseStateReadingFlag
	return nil
}

func (i *ArgParserImpl) processReadingArgumentList(arg string) *argtoolsError.Error {
	_ = i // check if pointer is nil here and check no further

	if !checkNoDashInFront(arg) {
		if len(i.mParsedArgData.ArgValues) == 0 {
			return i.notSetArgValueCase(arg)
		}

		if i.mFlagDescription, i.mOK = i.flagDescriptions[argParserConfig.Flag(arg)]; !i.mOK {
			return argtoolsError.NewError(
				argtoolsError.CodeArgParserUnexpectedFlag,
				fmt.Errorf(`ArgParserImpl.Parse: unexpected flag: "%s"`, arg))
		}

		i.mParsedArgData = nil
		if i.mArgDescription = i.mFlagDescription.GetArgDescription(); i.mArgDescription != nil {
			i.mParsedArgData = parsedData.NewParsedArgData(make([]parsedData.ArgValue, 0, 8))
		}

		i.mFlag = argParserConfig.Flag(arg)
		if _, i.mOK = i.rParsedFlagDataMap[i.mFlag]; i.mOK {
			return argtoolsError.NewError(
				argtoolsError.CodeArgParserDuplicateFlags,
				fmt.Errorf(`ArgParserImpl.Parse: duplicate flag: "%s"`, arg))
		}
		i.rParsedFlagDataMap[i.mFlag] = parsedData.NewParsedFlagData(i.mFlag, i.mParsedArgData)

		i.mParseState = getParseState(i.mArgDescription)
		return nil
	}

	i.mParsedArgData.ArgValues = append(i.mParsedArgData.ArgValues, parsedData.ArgValue(arg))

	return nil
}

func (i *ArgParserImpl) notSetArgValueCase(arg string) *argtoolsError.Error {
	// current command line argument looks like a flag
	// let's check if it is a flag
	i.mTmpFlagDescription, i.mOK = i.flagDescriptions[argParserConfig.Flag(arg)]
	if !i.mOK {
		return argtoolsError.NewError(
			argtoolsError.CodeArgParserDashInFrontOfArg,
			fmt.Errorf(`ArgParserImpl.Parse: argument "%s" contains a dash in front`, arg))
	}

	// arg is a flag, ok. but we are expecting flag/command argument value now,
	// so let's try to get it from default values slice
	if len(i.mArgDescription.GetDefaultValues()) == 0 {
		return argtoolsError.NewError(
			argtoolsError.CodeArgParserFlagMustHaveArg,
			fmt.Errorf(`ArgParserImpl.Parse: flag "%s" must have an arg`, arg))
	}
	i.mParsedArgData.ArgValues = copyDefaultValues2ArgValues(i.mArgDescription.GetDefaultValues(), i.mParsedArgData.ArgValues)

	// default value is set, good
	// time to switch logic to process a flag
	i.mParseState = parseStateReadingFlag
	return i.processReadingFlag(arg)
}

func getParseState(argumentsDescription *argParserConfig.ArgumentsDescription) parseState {
	switch argumentsDescription.GetAmountType() {
	case argParserConfig.ArgAmountTypeSingle:
		return parseStateReadingSingleArgument
	case argParserConfig.ArgAmountTypeList:
		return parseStateReadingArgumentList
	default:
		return parseStateReadingFlag
	}
}

func copyDefaultValues2ArgValues(defaultValueSlice []string, argValueSlice []parsedData.ArgValue) []parsedData.ArgValue {
	for i := 0; i < len(defaultValueSlice); i++ {
		argValueSlice = append(argValueSlice, parsedData.ArgValue(defaultValueSlice[i]))
	}

	return argValueSlice
}
