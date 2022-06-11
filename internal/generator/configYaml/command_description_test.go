package configYaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *CommandDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, Command(""), pointer.GetCommand())
		require.Equal(t, "", pointer.GetDescriptionHelpInfo())
		require.Nil(t, pointer.GetRequiredFlags())
		require.Nil(t, pointer.GetOptionalFlags())
		require.Nil(t, pointer.GetAdditionalCommands())
		require.Nil(t, pointer.GetArgumentsDescription())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &CommandDescription{
			Command:              Command(gofakeit.Name()),
			DescriptionHelpInfo:  gofakeit.Name(),
			RequiredFlags:        []Flag{Flag(gofakeit.Name())},
			OptionalFlags:        []Flag{Flag(gofakeit.Name())},
			AdditionalCommands:   []Command{Command(gofakeit.Name())},
			ArgumentsDescription: &ArgumentsDescription{},
		}

		require.Equal(t, pointer.Command, pointer.GetCommand())
		require.Equal(t, pointer.DescriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, pointer.RequiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, pointer.OptionalFlags, pointer.GetOptionalFlags())
		require.Equal(t, pointer.AdditionalCommands, pointer.GetAdditionalCommands())
		require.Equal(t, pointer.ArgumentsDescription, pointer.GetArgumentsDescription())
	})
}

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
			config, err := GetConfig(fmt.Sprintf("./testCases/command_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &CommandDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
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
			config, err := GetConfig(fmt.Sprintf("./testCases/command_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
