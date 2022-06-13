package argParser

import (
	"github.com/terryhay/argtools/internal/argParserImpl"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
)

func Parse(config argParserConfig.ArgParserConfig, args []string) (*parsedData.ParsedData, *argtoolsError.Error) {
	return argParserImpl.NewCmdArgParserImpl(config).Parse(args)
}
