package argParserConfig

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFlagDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *FlagDescription

		require.Equal(t, "", nilPointer.GetDescriptionHelpInfo())
		require.Nil(t, nilPointer.GetArgDescription())
	})

	t.Run("simple", func(t *testing.T) {
		pointer := &FlagDescription{
			DescriptionHelpInfo: gofakeit.Name(),
			ArgDescription:      &ArgumentsDescription{},
		}

		require.Equal(t, pointer.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.ArgDescription, pointer.GetArgDescription())
	})
}
