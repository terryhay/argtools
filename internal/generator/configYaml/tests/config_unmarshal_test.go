package tests

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_version.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: config unmarshal error: no required field \"version\"",
		},
		{
			yamlFileName:      "no_help_description.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: config unmarshal error: no required field \"help_command_description\"",
		},
		{
			yamlFileName:      "no_help_command_description.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: config unmarshal error: no required field \"help_command_description\"",
		},
		{
			yamlFileName:      "no_command_description_and_null_command.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: config unmarshal error: one or more of felds \"null_command_description\" or \"command_descriptions\" must be set",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := configYaml.GetConfig(fmt.Sprintf("./config_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}
}

func TestConfigUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_flag_descriptions.yaml",
		},
		{
			yamlFileName: "no_command_descriptions_but_has_null_command_description.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := configYaml.GetConfig(fmt.Sprintf("./config_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
