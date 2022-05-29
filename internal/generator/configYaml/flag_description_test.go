package configYaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlagDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *FlagDescription

		require.Equal(t, Flag(""), nilPointer.GetFlag())
		require.Equal(t, "", nilPointer.GetDescriptionHelpInfo())
		require.Equal(t, "", nilPointer.GetSynopsisDescription())
		require.Nil(t, nilPointer.GetArgumentsDescription())
	})

	t.Run("simple", func(t *testing.T) {
		flag := Flag(gofakeit.Name())
		descriptionHelpInfo := gofakeit.Name()
		synopsisDescription := gofakeit.Name()
		argumentsDescription := &ArgumentsDescription{}

		pointer := &FlagDescription{
			Flag:                 flag,
			DescriptionHelpInfo:  descriptionHelpInfo,
			SynopsisDescription:  synopsisDescription,
			ArgumentsDescription: argumentsDescription,
		}

		require.Equal(t, flag, pointer.GetFlag())
		require.Equal(t, descriptionHelpInfo, pointer.GetDescriptionHelpInfo())
		require.Equal(t, synopsisDescription, pointer.GetSynopsisDescription())
		require.Equal(t, argumentsDescription, pointer.GetArgumentsDescription())
	})
}

func TestFlagDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_flag.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: flagDescription unmarshal error: no required field \"flag\"",
		},
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: flagDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./testCases/flag_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &FlagDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestFlagDescriptionUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_arguments_description.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./testCases/flag_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
