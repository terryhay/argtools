package argtoolsError

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgtoolsError(t *testing.T) {
	t.Parallel()

	t.Run("nil_pointer_getters", func(t *testing.T) {
		var err *Error
		require.Equal(t, uint(CodeNone), err.Code().ToUint())
		require.Equal(t, "", err.Error())
	})

	t.Run("getters", func(t *testing.T) {
		err := fmt.Errorf(gofakeit.Name())
		argToolsErr := NewError(CodeUndefinedError, err)

		require.NotNil(t, argToolsErr)
		require.Equal(t, uint(CodeUndefinedError), argToolsErr.Code().ToUint())
		require.Equal(t, err.Error(), argToolsErr.Error())
	})
}
