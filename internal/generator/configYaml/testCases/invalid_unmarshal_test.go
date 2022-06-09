package testCases_test

import (
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalErrorsInvalidPath(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName          string
		configYamlPath    string
		expectedErrorCode argtoolsError.Code
	}{
		{
			caseName:          "empty_string_path",
			configYamlPath:    "",
			expectedErrorCode: argtoolsError.CodeGetConfigReadFileError,
		},
		{
			caseName:          "not_existed_path",
			configYamlPath:    "./unexist/path",
			expectedErrorCode: argtoolsError.CodeGetConfigReadFileError,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			_, err := configYaml.GetConfig(td.configYamlPath)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}
