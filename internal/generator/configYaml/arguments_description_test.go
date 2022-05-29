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

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *ArgumentsDescription

		require.Equal(t, argParserConfig.ArgAmountTypeNoArgs, nilPointer.GetAmountType())
		require.Equal(t, "", nilPointer.GetSynopsisHelpDescription())
		require.Nil(t, nilPointer.GetDefaultValues())
		require.Nil(t, nilPointer.GetAllowedValues())
	})

	t.Run("simple", func(t *testing.T) {
		amountType := argParserConfig.ArgAmountTypeSingle
		synopsisHelpDescription := gofakeit.Name()
		defaultValues := []string{gofakeit.Name()}
		allowedValues := []string{gofakeit.Name()}

		pointer := &ArgumentsDescription{
			AmountType:              amountType,
			SynopsisHelpDescription: synopsisHelpDescription,
			DefaultValues:           defaultValues,
			AllowedValues:           allowedValues,
		}

		require.Equal(t, amountType, pointer.GetAmountType())
		require.Equal(t, synopsisHelpDescription, pointer.GetSynopsisHelpDescription())
		require.Equal(t, defaultValues, pointer.GetDefaultValues())
		require.Equal(t, allowedValues, pointer.GetAllowedValues())
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
			expectedErrorText: "configYaml.GetConfig: unmarshal error: argumentsDescriptions unmarshal error: no required field \"synopsis_description\"",
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
