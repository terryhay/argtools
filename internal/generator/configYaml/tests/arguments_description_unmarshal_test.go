package tests

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

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
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := configYaml.GetConfig(fmt.Sprintf("./arguments_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}
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
			config, err := configYaml.GetConfig(fmt.Sprintf("./arguments_description_cases/%s", td.yamlFileName))
			require.NotNil(t, config)
			require.Nil(t, err)
		})
	}
}
