package argParserConfig

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNamelessCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *NamelessCommandDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, CommandIDUndefined, pointer.GetID())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Nil(t, pointer.GetArgDescription())
		require.Nil(t, pointer.GetRequiredFlags())
		require.Nil(t, pointer.GetOptionalFlags())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &NamelessCommandDescription{
			ID:                  CommandID(gofakeit.Uint32()),
			DescriptionHelpInfo: gofakeit.Name(),
			ArgDescription:      &ArgumentsDescription{},
			RequiredFlags:       map[Flag]bool{Flag(gofakeit.Name()): true},
			OptionalFlags:       map[Flag]bool{Flag(gofakeit.Name()): true},
		}

		require.Equal(t, pointer.ID, pointer.GetID())
		require.Equal(t, pointer.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.ArgDescription, pointer.GetArgDescription())
		require.Equal(t, pointer.RequiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, pointer.OptionalFlags, pointer.GetOptionalFlags())
	})
}
