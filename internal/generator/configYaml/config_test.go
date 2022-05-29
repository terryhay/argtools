package configYaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *Config

		require.Equal(t, Version(""), nilPointer.GetVersion())
		require.Nil(t, nilPointer.GetAppHelpDescription())
		require.Nil(t, nilPointer.GetHelpCommandDescription())
		require.Nil(t, nilPointer.GetNullCommandDescription())
		require.Nil(t, nilPointer.GetCommandDescriptions())
		require.Nil(t, nilPointer.GetFlagDescriptions())
	})

	t.Run("simple", func(t *testing.T) {
		version := Version(gofakeit.Name())
		appHelpDescription := &AppHelpDescription{}
		helpCommandDescription := &HelpCommandDescription{}
		nullCommandDescription := &NullCommandDescription{}
		commandDescriptions := []*CommandDescription{{}}
		flagDescriptions := []*FlagDescription{{}}

		pointer := &Config{
			Version:                version,
			AppHelpDescription:     appHelpDescription,
			HelpCommandDescription: helpCommandDescription,
			NullCommandDescription: nullCommandDescription,
			CommandDescriptions:    commandDescriptions,
			FlagDescriptions:       flagDescriptions,
		}

		require.Equal(t, version, pointer.GetVersion())
		require.Equal(t, appHelpDescription, pointer.GetAppHelpDescription())
		require.Equal(t, helpCommandDescription, pointer.GetHelpCommandDescription())
		require.Equal(t, nullCommandDescription, pointer.GetNullCommandDescription())
		require.Equal(t, commandDescriptions, pointer.GetCommandDescriptions())
		require.Equal(t, flagDescriptions, pointer.GetFlagDescriptions())
	})
}

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
			yamlFileName:      "no_app_help_description.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: config unmarshal error: no required field \"app_help_description\"",
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
			config, err := GetConfig(fmt.Sprintf("./testCases/config_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &Config{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
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
			config, err := GetConfig(fmt.Sprintf("./testCases/config_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
