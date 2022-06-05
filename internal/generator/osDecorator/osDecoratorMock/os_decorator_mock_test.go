package osDecoratorMock

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/internal/generator/osDecorator"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"os"
	"testing"
)

func TestNewMockOSDecorator(t *testing.T) {
	t.Parallel()

	mockArgs := []string{gofakeit.Color()}
	mockCreateFuncErrRes := fmt.Errorf(gofakeit.Name())
	mockIsNotExistFuncRes := gofakeit.Bool()
	mockMkdirAllErrRes := fmt.Errorf(gofakeit.Name())
	mockStatFuncErrRes := fmt.Errorf(gofakeit.Name())

	mockOSDecorator := NewMockOSDecorator(
		MockOSDecoratorInit{
			Args: mockArgs,
			CreateFunc: func(path string) (osDecorator.FileDecorator, error) {
				return nil, mockCreateFuncErrRes
			},
			ExitFunc: func(err *argtoolsError.Error) {

			},
			IsNotExistFunc: func(err error) bool {
				return mockIsNotExistFuncRes
			},
			MkdirAllFunc: func(path string, perm os.FileMode) error {
				return mockMkdirAllErrRes
			},
			StatFunc: func(name string) (os.FileInfo, error) {
				return nil, mockStatFuncErrRes
			},
		},
	)

	require.Equal(t, mockArgs, mockOSDecorator.GetArgs())

	_, err := mockOSDecorator.Create("")
	require.Equal(t, mockCreateFuncErrRes, err)

	mockOSDecorator.Exit(nil)

	res := mockOSDecorator.IsNotExist(nil)
	require.Equal(t, res, mockIsNotExistFuncRes)

	err = mockOSDecorator.MkdirAll("", 0)
	require.Equal(t, err, mockMkdirAllErrRes)

	_, err = mockOSDecorator.Stat("")
	require.Equal(t, err, mockStatFuncErrRes)
}

func TestMockFileDecorator(t *testing.T) {
	t.Parallel()

	mockCloseErrRes := fmt.Errorf(gofakeit.Name())
	mockWriteStringErrRes := fmt.Errorf(gofakeit.Name())

	mockFileDecorator := NewMockFileDecorator(
		func() error {
			return mockCloseErrRes
		},
		func(s string) error {
			return mockWriteStringErrRes
		})

	err := mockFileDecorator.Close()
	require.Equal(t, mockCloseErrRes, err)

	err = mockFileDecorator.WriteString("")
	require.Equal(t, mockWriteStringErrRes, err)
}
