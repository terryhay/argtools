package parsedData

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"testing"
)

func TestParsedDataGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *ParsedData

		require.Equal(t, argParserConfig.CommandIDUndefined, nilPointer.GetCommandID())
		require.Equal(t, argParserConfig.CommandUndefined, nilPointer.GetCommand())
		require.Nil(t, nilPointer.GetAgrData())
		require.Nil(t, nilPointer.GetFlagData())
	})

	t.Run("simple", func(t *testing.T) {
		commandID := argParserConfig.CommandID(gofakeit.Uint32())
		command := argParserConfig.Command(gofakeit.Name())
		argData := NewParsedArgData(nil)
		flagData := map[argParserConfig.Flag]*ParsedFlagData{
			argParserConfig.Flag(gofakeit.Name()): NewParsedFlagData(
				argParserConfig.Flag(gofakeit.Name()),
				NewParsedArgData(nil)),
		}

		pointer := NewParsedData(commandID, command, argData, flagData)

		require.Equal(t, commandID, pointer.GetCommandID())
		require.Equal(t, command, pointer.GetCommand())
		require.Equal(t, argData, pointer.GetAgrData())
		require.Equal(t, flagData, pointer.GetFlagData())
	})
}
