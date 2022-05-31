package parsedData

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParsedArgDataGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *ParsedArgData

		require.Nil(t, nilPointer.GetArgValues())
	})

	t.Run("simple", func(t *testing.T) {
		argValues := []ArgValue{ArgValue(gofakeit.Name())}

		pointer := NewParsedArgData(argValues)

		require.Equal(t, argValues, pointer.GetArgValues())
	})
}
