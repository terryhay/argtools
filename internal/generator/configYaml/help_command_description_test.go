package configYaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelpCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *HelpCommandDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, Command(""), pointer.GetCommand())
		require.Nil(t, pointer.GetAdditionalCommands())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &HelpCommandDescription{
			Command:            Command(gofakeit.Name()),
			AdditionalCommands: []Command{Command(gofakeit.Name())},
		}

		require.Equal(t, pointer.Command, pointer.GetCommand())
		require.Equal(t, pointer.AdditionalCommands, pointer.GetAdditionalCommands())
	})
}

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
			config, err := GetConfig(fmt.Sprintf("./testCases/help_command_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &HelpCommandDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
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
			config, err := GetConfig(fmt.Sprintf("./testCases/help_command_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
