package parsedData

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
)

// ParsedData - all parsed Command line data
type ParsedData struct {
	CommandID argParserConfig.CommandID
	Command   argParserConfig.Command
	ArgData   *ParsedArgData
	FlagData  map[argParserConfig.Flag]*ParsedFlagData
}

// NewParsedData - ParsedData object constructor
func NewParsedData(
	commandID argParserConfig.CommandID,
	command argParserConfig.Command,
	argData *ParsedArgData,
	flagData map[argParserConfig.Flag]*ParsedFlagData,
) *ParsedData {
	if len(flagData) == 0 {
		flagData = nil
	}
	return &ParsedData{
		CommandID: commandID,
		Command:   command,
		ArgData:   argData,
		FlagData:  flagData,
	}
}

// GetCommandID - CommandID field getter
func (i *ParsedData) GetCommandID() argParserConfig.CommandID {
	if i == nil {
		return argParserConfig.CommandIDUndefined
	}
	return i.CommandID
}

// GetCommand - Command field getter
func (i *ParsedData) GetCommand() argParserConfig.Command {
	if i == nil {
		return argParserConfig.CommandUndefined
	}
	return i.Command
}

// GetAgrData - AgrData field getter
func (i *ParsedData) GetAgrData() *ParsedArgData {
	if i == nil {
		return nil
	}
	return i.ArgData
}

// GetFlagData - FlagData field getter
func (i *ParsedData) GetFlagData() map[argParserConfig.Flag]*ParsedFlagData {
	if i == nil {
		return nil
	}
	return i.FlagData
}

// GetFlagArgValue - extract flag argument value
func (i *ParsedData) GetFlagArgValue(flag argParserConfig.Flag) (ArgValue, *argtoolsError.Error) {
	values, err := i.GetFlagArgValues(flag)
	if err != nil {
		return "", argtoolsError.NewError(
			err.Code(),
			fmt.Errorf(`ParsedData.GetFlagArgValue: %v`, err.Error()))
	}
	if len(values) == 0 {
		return "", argtoolsError.NewError(
			argtoolsError.CodeUndefinedError, // todo
			fmt.Errorf(`ParsedData.GetFlagArgValue: flag "%s" doesn't contain argument values'`, flag))
	}

	return values[0], nil
}

// GetFlagArgValues - extract flag argument value slice
func (i *ParsedData) GetFlagArgValues(flag argParserConfig.Flag) ([]ArgValue, *argtoolsError.Error) {
	if i == nil {
		return nil, argtoolsError.NewError(
			argtoolsError.CodeUndefinedError, // todo
			fmt.Errorf(`ParsedData.GetFlagArgValues: try to call method by nil pointer`))
	}
	parsedFlagData, ok := i.GetFlagData()[flag]
	if !ok {
		return nil, argtoolsError.NewError(
			argtoolsError.CodeUndefinedError, // todo
			fmt.Errorf(`ParsedData.GetFlagArgValues: flag "%s" is not found in flag data map`, flag))
	}

	return parsedFlagData.GetArgData().GetArgValues(), nil
}
