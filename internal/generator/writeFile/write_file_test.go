package writeFile

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/internal/generator/osDecorator"
	"github.com/terryhay/argtools/internal/generator/osDecorator/osDecoratorMock"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	t.Parallel()

	mockCreateFuncErrRes := fmt.Errorf(gofakeit.Name())
	mockMkdirAllErrRes := fmt.Errorf(gofakeit.Name())
	mockStatFuncErrRes := fmt.Errorf(gofakeit.Name())

	dirPath := gofakeit.Color()

	testData := []struct {
		caseName string

		osd      osDecorator.OSDecorator
		dirPath  string
		fileBody string

		expectedErrCode argtoolsError.Code
	}{
		{
			caseName: "check_dir_path_error",

			osd: osDecoratorMock.NewMockOSDecorator(osDecoratorMock.MockOSDecoratorInit{
				IsNotExistFunc: func(err error) bool {
					return err != nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, mockStatFuncErrRes
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: argtoolsError.CodeGeneratorInvalidGeneratePath,
		},
		{
			caseName: "check_dir_path_error",

			osd: osDecoratorMock.NewMockOSDecorator(osDecoratorMock.MockOSDecoratorInit{
				IsNotExistFunc: func(err error) bool {
					return err != nil
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return mockMkdirAllErrRes
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					if path == dirPath {
						return nil, nil
					}
					return nil, mockStatFuncErrRes
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: argtoolsError.CodeGeneratorCreateDirError,
		},
		{
			caseName: "file_create_error",

			osd: osDecoratorMock.NewMockOSDecorator(osDecoratorMock.MockOSDecoratorInit{
				CreateFunc: func(path string) (osDecorator.FileDecorator, error) {
					return nil, mockCreateFuncErrRes
				},
				IsNotExistFunc: func(err error) bool {
					return false
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, nil
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: argtoolsError.CodeGeneratorCreateFileError,
		},
		{
			caseName: "file_create_error",

			osd: osDecoratorMock.NewMockOSDecorator(osDecoratorMock.MockOSDecoratorInit{
				CreateFunc: func(path string) (osDecorator.FileDecorator, error) {
					return osDecoratorMock.NewMockFileDecorator(
							func() error {
								return nil
							},
							func(s string) error {
								return fmt.Errorf(gofakeit.Color())
							}),
						nil
				},
				IsNotExistFunc: func(err error) bool {
					return false
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, nil
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: argtoolsError.CodeGeneratorWriteFileError,
		},
		{
			caseName: "file_close_error",

			osd: osDecoratorMock.NewMockOSDecorator(osDecoratorMock.MockOSDecoratorInit{
				CreateFunc: func(path string) (osDecorator.FileDecorator, error) {
					return osDecoratorMock.NewMockFileDecorator(
							func() error {
								return fmt.Errorf(gofakeit.Color())
							},
							func(s string) error {
								return nil
							}),
						nil
				},
				IsNotExistFunc: func(err error) bool {
					return false
				},
				MkdirAllFunc: func(path string, perm os.FileMode) error {
					return nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, nil
				},
			}),
			dirPath:         dirPath,
			expectedErrCode: argtoolsError.CodeGeneratorFileCloseError,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := Write(td.osd, td.dirPath, td.fileBody)
			if td.expectedErrCode == argtoolsError.CodeNone {
				require.Nil(t, err)
				return
			}
			require.Equal(t, td.expectedErrCode, err.Code())
		})
	}
}
