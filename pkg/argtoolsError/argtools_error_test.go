package argtoolsError

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArgtoolsError(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf(gofakeit.Name())
	argToolsErr := NewError(CodeUndefinedError, err)

	require.NotNil(t, argToolsErr)
	require.Equal(t, CodeUndefinedError, argToolsErr.Code())
	require.Equal(t, err.Error(), argToolsErr.Error())
}
