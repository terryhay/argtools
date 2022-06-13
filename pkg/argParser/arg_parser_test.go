package argParser

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"testing"
)

func TestArgParser(t *testing.T) {
	res, err := Parse(
		argParserConfig.NewArgParserConfig(
			argParserConfig.ApplicationDescription{},
			nil,
			nil,
			nil,
			nil),
		nil)

	require.Nil(t, res)
	require.NotNil(t, err)
}
