package parsedData

import "github.com/terryhay/argtools/pkg/argParserConfig"

// ParsedFlagData - parsed flag and arguments array
type ParsedFlagData struct {
	Flag    argParserConfig.Flag
	ArgData *ParsedArgData
}

// NewParsedFlagData - ParsedFlagData object constructor
func NewParsedFlagData(flag argParserConfig.Flag, argData *ParsedArgData) *ParsedFlagData {
	return &ParsedFlagData{
		Flag:    flag,
		ArgData: argData,
	}
}

// GetFlag - flag field getter
func (i *ParsedFlagData) GetFlag() argParserConfig.Flag {
	if i == nil {
		return argParserConfig.FlagUndefined
	}
	return i.Flag
}

// GetArgData - ArgData field getter
func (i *ParsedFlagData) GetArgData() *ParsedArgData {
	if i == nil {
		return nil
	}
	return i.ArgData
}
