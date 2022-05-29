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

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *CommandDescription

		require.Equal(t, Command(""), nilPointer.GetCommand())
		require.Equal(t, "", nilPointer.GetDescriptionHelpInfo())
		require.Nil(t, nilPointer.GetRequiredFlags())
		require.Nil(t, nilPointer.GetOptionalFlags())
		require.Nil(t, nilPointer.GetAdditionalCommands())
		require.Nil(t, nilPointer.GetArgumentsDescription())
	})

	t.Run("simple", func(t *testing.T) {
		command := Command(gofakeit.Name())
		descriptionHelpInfo := gofakeit.Name()
		requiredFlags := []Flag{Flag(gofakeit.Name())}
		optionalFlags := []Flag{Flag(gofakeit.Name())}
		additionalCommands := []Command{Command(gofakeit.Name())}
		argumentsDescription := &ArgumentsDescription{}

		pointer := &CommandDescription{
			Command:              command,
			DescriptionHelpInfo:  descriptionHelpInfo,
			RequiredFlags:        requiredFlags,
			OptionalFlags:        optionalFlags,
			AdditionalCommands:   additionalCommands,
			ArgumentsDescription: argumentsDescription,
		}

		require.Equal(t, command, pointer.GetCommand())
		require.Equal(t, descriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, requiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, optionalFlags, pointer.GetOptionalFlags())
		require.Equal(t, additionalCommands, pointer.GetAdditionalCommands())
		require.Equal(t, argumentsDescription, pointer.GetArgumentsDescription())
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
