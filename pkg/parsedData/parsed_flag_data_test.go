package parsedData

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"testing"
)

func TestParsedFlagDataGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *ParsedFlagData

		require.Equal(t, argParserConfig.Flag(""), nilPointer.GetFlag())
		require.Nil(t, nilPointer.GetArgData())
	})

	t.Run("simple", func(t *testing.T) {
		flag := argParserConfig.Flag(gofakeit.Name())
		argData := NewParsedArgData(nil)

		pointer := NewParsedFlagData(flag, argData)

		require.Equal(t, flag, pointer.GetFlag())
		require.Equal(t, argData, pointer.GetArgData())
	})
}
