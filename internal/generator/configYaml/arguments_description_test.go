package configYaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArgumentsDescriptionGetters(t *testing.T) {
	t.Parallel()

	var pointer *ArgumentsDescription

	t.Run("nil_pointer", func(t *testing.T) {
		require.Equal(t, argParserConfig.ArgAmountTypeNoArgs, pointer.GetAmountType())
		require.Equal(t, "", pointer.GetSynopsisHelpDescription())
		require.Nil(t, pointer.GetDefaultValues())
		require.Nil(t, pointer.GetAllowedValues())
	})

	t.Run("initialized_pointer", func(t *testing.T) {
		pointer = &ArgumentsDescription{
			AmountType:              argParserConfig.ArgAmountTypeSingle,
			SynopsisHelpDescription: gofakeit.Name(),
			DefaultValues:           []string{gofakeit.Name()},
			AllowedValues:           []string{gofakeit.Name()},
		}

		require.Equal(t, pointer.AmountType, pointer.GetAmountType())
		require.Equal(t, pointer.SynopsisHelpDescription, pointer.GetSynopsisHelpDescription())
		require.Equal(t, pointer.DefaultValues, pointer.GetDefaultValues())
		require.Equal(t, pointer.AllowedValues, pointer.GetAllowedValues())
	})
}

func TestArgumentsDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_amount_type.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: argumentsDescriptions unmarshal error: no required field \"amount_type\"",
		},
		{
			yamlFileName:      "no_synopsis_description.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: argumentsDescriptions unmarshal error: no required field \"synopsis_description\"",
		},
		{
			yamlFileName:      "unexpected_amount_type.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: argumentsDescriptions unmarshal error: can't convert string value \"amount_type\": unexpected \"amount_type\" value: trash\\nallowed values: \"single\", \"array\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./testCases/arguments_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &ArgumentsDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}

func TestArgumentsDescriptionUnmarshalNoErrorWhenNoOptionalFields(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName string
	}{
		{
			yamlFileName: "no_default_values.yaml",
		},
		{
			yamlFileName: "no_allowed_values.yaml",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./testCases/arguments_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
