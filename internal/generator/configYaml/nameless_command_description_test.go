package configYaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"
)

func TestNamelessCommandDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *NamelessCommandDescription

		require.Equal(t, "", nilPointer.GetDescriptionHelpInfo())
		require.Nil(t, nilPointer.GetRequiredFlags())
		require.Nil(t, nilPointer.GetOptionalFlags())
		require.Nil(t, nilPointer.GetArgumentsDescription())
	})

	t.Run("simple", func(t *testing.T) {
		descriptionHelpInfo := gofakeit.Name()
		requiredFlags := []Flag{Flag(gofakeit.Name())}
		optionalFlags := []Flag{Flag(gofakeit.Name())}
		argumentsDescription := &ArgumentsDescription{}

		pointer := &NamelessCommandDescription{
			DescriptionHelpInfo:  descriptionHelpInfo,
			RequiredFlags:        requiredFlags,
			OptionalFlags:        optionalFlags,
			ArgumentsDescription: argumentsDescription,
		}

		require.Equal(t, descriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, requiredFlags, pointer.GetRequiredFlags())
		require.Equal(t, optionalFlags, pointer.GetOptionalFlags())
		require.Equal(t, argumentsDescription, pointer.GetArgumentsDescription())
	})
}

func TestNamelessCommandDescriptionErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: commandDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./testCases/nameless_command_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &NamelessCommandDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestNamelessCommandDescriptionNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_arguments_description.yaml",
		},
		{
			yamlFileName: "no_optional_flags.yaml",
		},
		{
			yamlFileName: "no_required_flags.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./testCases/nameless_command_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
