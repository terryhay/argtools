package tests

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelpCommandDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_command.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: helpCommandDescription unmarshal error: no required field \"command\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := configYaml.GetConfig(fmt.Sprintf("./help_command_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}
}

func TestHelpCommandDescriptionUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_additional_commands.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := configYaml.GetConfig(fmt.Sprintf("./help_command_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
