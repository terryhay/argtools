package argParserConfig

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *CommandDescription

		require.Equal(t, CommandIDUndefined, nilPointer.GetID())
		require.Equal(t, "", nilPointer.GetDescriptionHelpInfo())
		require.Nil(t, nilPointer.GetCommands())
		require.Nil(t, nilPointer.GetArgDescription())
		require.Nil(t, nilPointer.GetRequiredFlags())
		require.Nil(t, nilPointer.GetOptionalFlags())
	})

	t.Run("simple", func(t *testing.T) {
		pointer := &CommandDescription{
			ID:                  CommandID(gofakeit.Uint32()),
			DescriptionHelpInfo: gofakeit.Name(),
			Commands: map[Command]bool{
				Command(gofakeit.Name()): true,
			},
			ArgDescription: &ArgumentsDescription{},
			RequiredFlags: map[Flag]bool{
				Flag(gofakeit.Name()): true,
			},
			OptionalFlags: map[Flag]bool{
				Flag(gofakeit.Name()): true,
			},
		}

		require.Equal(t, pointer.ID, pointer.GetID())
		require.Equal(t, pointer.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.Commands, pointer.GetCommands())
		require.Equal(t, pointer.ArgDescription, pointer.GetArgDescription())
		require.Equal(t, pointer.RequiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, pointer.OptionalFlags, pointer.GetOptionalFlags())
	})
}
