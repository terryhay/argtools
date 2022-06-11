package configDataExtractor

import (
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractFlagDescriptionMapErrors(t *testing.T) {
	t.Parallel()

	testData := []*struct {
		caseName          string
		flagDescriptions  []*configYaml.FlagDescription
		expectedErrorCode argtoolsError.Code
	}{
		{
			caseName: "single_empty_flag_description",
			flagDescriptions: []*configYaml.FlagDescription{
				nil,
			},
			expectedErrorCode: argtoolsError.CodeUndefinedError,
		},
		{
			caseName: "empty_flag_description_in_front",
			flagDescriptions: []*configYaml.FlagDescription{
				nil,
				{},
			},
			expectedErrorCode: argtoolsError.CodeUndefinedError,
		},
		{
			caseName: "empty_flag_description_in_back",
			flagDescriptions: []*configYaml.FlagDescription{
				{},
				nil,
			},
			expectedErrorCode: argtoolsError.CodeUndefinedError,
		},

		{
			caseName: "duplicate_flag_descriptions",
			flagDescriptions: []*configYaml.FlagDescription{
				{
					Flag: "flag",
				},
				{
					Flag: "flag",
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			flagDescriptionMap, err := ExtractFlagDescriptionMap(td.flagDescriptions)
			require.Nil(t, flagDescriptionMap)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}

func TestExtractFlagDescriptionMap(t *testing.T) {
	t.Parallel()

	testData := []*struct {
		caseName         string
		flagDescriptions []*configYaml.FlagDescription
		expectedMap      map[configYaml.Flag]*configYaml.FlagDescription
	}{
		{
			caseName:         "no_flag_description",
			flagDescriptions: nil,
		},

		{
			caseName: "single_flag_description",
			flagDescriptions: []*configYaml.FlagDescription{
				{
					Flag: "flag",
				},
			},
			expectedMap: map[configYaml.Flag]*configYaml.FlagDescription{
				"flag": {
					Flag: "flag",
				},
			},
		},
		{
			caseName: "two_flag_descriptions",
			flagDescriptions: []*configYaml.FlagDescription{
				{
					Flag: "flag1",
				},
				{
					Flag: "flag2",
				},
			},
			expectedMap: map[configYaml.Flag]*configYaml.FlagDescription{
				"flag1": {
					Flag: "flag1",
				},
				"flag2": {
					Flag: "flag2",
				},
			},
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			flagDescriptionMap, err := ExtractFlagDescriptionMap(td.flagDescriptions)
			require.Nil(t, err)
			require.Equal(t, len(td.expectedMap), len(flagDescriptionMap))

			for flag, expectedFlagDescription := range td.expectedMap {
				flagDescription, contain := flagDescriptionMap[flag]
				require.True(t, contain)
				require.Equal(t, expectedFlagDescription.Flag, flagDescription.Flag)
			}
		})
	}
}
