package tests_test

import (
	"argtools/internal/generator/configYaml"
	"argtools/pkg/argtoolsError"
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
			configYamlPath:    "./unexised/path",
			expectedErrorCode: argtoolsError.CodeGetConfigReadFileError,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			config, err := configYaml.GetConfig(td.configYamlPath)
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}
