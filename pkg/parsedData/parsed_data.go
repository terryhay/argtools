package parsedData

import (
	"github.com/terryhay/argtools/pkg/argParserConfig"
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
