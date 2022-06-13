package argParserImpl

import (
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
)

type parseState uint8

const (
	parseStateReadingFlag parseState = iota
	parseStateReadingSingleArgument
	parseStateReadingArgumentList
)

const countOfParseStates = int(parseStateReadingArgumentList) + 1

// ArgParserImpl - implementation of command line argument parser
type ArgParserImpl struct {
	commandDescriptions        map[argParserConfig.Command]*argParserConfig.CommandDescription
	namelessCommandDescription *argParserConfig.CommandDescription
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
	res := &ArgParserImpl{
		commandDescriptions: createCommandDescriptionMap(
			config.GetCommandDescriptions(),
			config.GetHelpCommandDescription()),
		namelessCommandDescription: castNamelessCommandDescriptionToPointer(config.GetNamelessCommandDescription()),
		flagDescriptions:           config.GetFlagDescriptions(),
	}
	res.stateProcessors = [countOfParseStates]func(arg string) *argtoolsError.Error{
		parseStateReadingFlag:           res.processReadingFlag,
		parseStateReadingSingleArgument: res.processReadingSingleArgument,
		parseStateReadingArgumentList:   res.processReadingArgumentList,
	}

	return res
}

func createCommandDescriptionMap(
	commandsDescriptionSlice []*argParserConfig.CommandDescription,
	helpCommandDescription argParserConfig.HelpCommandDescription,
) map[argParserConfig.Command]*argParserConfig.CommandDescription {

	commandsCount := 0
	commandDescription := castHelpCommandDescriptionToPointer(helpCommandDescription)
	if commandDescription != nil {
		commandsCount++
	}
	commandsCount += len(commandsDescriptionSlice)

	res := make(map[argParserConfig.Command]*argParserConfig.CommandDescription, commandsCount)

	var command argParserConfig.Command
	for command = range commandDescription.GetCommands() {
		res[command] = commandDescription
	}

	for i := range commandsDescriptionSlice {
		for command = range commandsDescriptionSlice[i].GetCommands() {
			res[command] = commandDescription
		}
	}

	return res
}

func castHelpCommandDescriptionToPointer(helpCommandDescription argParserConfig.HelpCommandDescription) *argParserConfig.CommandDescription {
	if helpCommandDescription == nil {
		return nil
	}
	return helpCommandDescription.(*argParserConfig.CommandDescription)
}

func castNamelessCommandDescriptionToPointer(namelessCommandDescription argParserConfig.NamelessCommandDescription) *argParserConfig.CommandDescription {
	if namelessCommandDescription == nil {
		return nil
	}
	return namelessCommandDescription.(*argParserConfig.CommandDescription)
}
