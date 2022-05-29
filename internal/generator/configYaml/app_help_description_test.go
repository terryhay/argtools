package configYaml

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppHelpDescriptionGetters(t *testing.T) {
	t.Parallel()

	t.Run("null_pointer", func(t *testing.T) {
		var nilPointer *AppHelpDescription

		require.Equal(t, "", nilPointer.GetApplicationName())
		require.Equal(t, "", nilPointer.GetNameHelpInfo())
		require.Equal(t, "", nilPointer.GetDescriptionHelpInfo())
	})

	t.Run("simple", func(t *testing.T) {
		applicationName := gofakeit.Name()
		nameHelpInfo := gofakeit.Name()
		descriptionHelpInfo := gofakeit.Name()

		pointer := &AppHelpDescription{
			ApplicationName:     applicationName,
			NameHelpInfo:        nameHelpInfo,
			DescriptionHelpInfo: descriptionHelpInfo,
		}

		require.Equal(t, applicationName, pointer.GetApplicationName())
		require.Equal(t, nameHelpInfo, pointer.GetNameHelpInfo())
		require.Equal(t, descriptionHelpInfo, pointer.GetDescriptionHelpInfo())
	})
}

func TestAppHelpDescriptionUnmarshalErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		yamlFileName      string
		expectedErrorText string
	}{
		{
			yamlFileName:      "no_app_name.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: appHelpDescription unmarshal error: no required field \"app_name\"",
		},
		{
			yamlFileName:      "no_name_help_info.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: appHelpDescription unmarshal error: no required field \"name_help_info\"",
		},
		{
			yamlFileName:      "no_description_help_info.yaml",
			expectedErrorText: "configYaml.GetConfig: unmarshal error: appHelpDescription unmarshal error: no required field \"description_help_info\"",
		},
	}

	for _, td := range testData {
		t.Run(td.yamlFileName, func(t *testing.T) {
			config, err := GetConfig(fmt.Sprintf("./testCases/app_help_description_cases/%s", td.yamlFileName))
			require.Nil(t, config)
			require.NotNil(t, err)
			require.Equal(t, argtoolsError.CodeGetConfigUnmarshalError, err.Code())
			require.Equal(t, td.expectedErrorText, err.Error())
		})
	}

	t.Run("fake_unmarshal_error", func(t *testing.T) {
		pointer := &AppHelpDescription{}
		err := pointer.UnmarshalYAML(func(interface{}) error {
			return fmt.Errorf("error")
		})

		require.NotNil(t, err)
	})
}
