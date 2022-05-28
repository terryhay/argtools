package tests

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_command.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: commandDescription unmarshal error: no required field \"command\"",
		},
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: commandDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := configYaml.GetConfig(fmt.Sprintf("./command_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}
}

func TestCommandDescriptionUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_additional_names.yaml",
		},
		{
			yamlFileName: "no_arguments_description.yaml",
		},
		{
			yamlFileName: "no_required_flags.yaml",
		},
		{
			yamlFileName: "no_optional_flags.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := configYaml.GetConfig(fmt.Sprintf("./command_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
