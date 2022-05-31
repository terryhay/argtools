package argParserConfig

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgParserConfigGetters(t *testing.T) {
	t.Parallel()

	t.Run("simple", func(t *testing.T) {
		pointer := NewArgParserConfig(
			ApplicationDescription{},
			map[Flag]*FlagDescription{
				Flag(gofakeit.Name()): {},
			},
			[]*CommandDescription{
				{},
			},
			&NullCommandDescription{},
		)

		require.Equal(t, pointer.AppDescription, pointer.GetAppDescription())
		require.Equal(t, pointer.FlagDescriptions, pointer.GetFlagDescriptions())
		require.Equal(t, pointer.CommandDescriptions, pointer.GetCommandDescriptions())
		require.Equal(t, pointer.NullCommandDescription, pointer.GetNullCommandDescription())
	})
}
