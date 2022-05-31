package argParserConfig

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgumentsDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *ArgumentsDescription

		require.Equal(t, ArgAmountTypeNoArgs, nilPointer.GetAmountType())
		require.Equal(t, "", nilPointer.GetSynopsisHelpDescription())
		require.Nil(t, nilPointer.GetDefaultValues())
		require.Nil(t, nilPointer.GetAllowedValues())
	})

	t.Run("simple", func(t *testing.T) {
		pointer := &ArgumentsDescription{
			AmountType:              ArgAmountTypeSingle,
			SynopsisHelpDescription: gofakeit.Name(),
			DefaultValues:           []string{gofakeit.Name()},
			AllowedValues: map[string]bool{
				gofakeit.Name(): true,
			},
		}

		require.Equal(t, pointer.AmountType, pointer.GetAmountType())
		require.Equal(t, pointer.SynopsisHelpDescription, pointer.GetSynopsisHelpDescription())
		require.Equal(t, pointer.DefaultValues, pointer.GetDefaultValues())
		require.Equal(t, pointer.AllowedValues, pointer.GetAllowedValues())
	})
}
