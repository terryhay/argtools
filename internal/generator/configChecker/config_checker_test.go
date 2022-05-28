package configChecker

import (
	"argtools/internal/generator/configYaml"
	"argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigCheckerErrors(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName               string
		nullCommandDescription *configYaml.NullCommandDescription
		commandDescriptionMap  map[configYaml.Command]*configYaml.CommandDescription
		flagDescriptionMap     map[configYaml.Flag]*configYaml.FlagDescription
		expectedErrorCode      argtoolsError.Code
	}{
		{
			caseName: "duplicate_flag_in_required_list",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command": {
					RequiredFlags: []configYaml.Flag{
						"flag",
						"flag",
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "duplicate_flag_in_optional_list",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command": {
					OptionalFlags: []configYaml.Flag{
						"flag",
						"flag",
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "duplicate_flag_in_required_and_optional_lists",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command": {
					RequiredFlags: []configYaml.Flag{
						"flag",
					},
					OptionalFlags: []configYaml.Flag{
						"flag",
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},

		{
			caseName: "unused_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command": {},
			},
			flagDescriptionMap: map[configYaml.Flag]*configYaml.FlagDescription{
				"flag": {
					Flag: "flag",
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigFlagIsNotUsedInCommands,
		},

		{
			caseName: "undefined_required_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command": {
					RequiredFlags: []configYaml.Flag{
						"flag",
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigUndefinedFlag,
		},
		{
			caseName: "undefined_optional_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				"command": {
					OptionalFlags: []configYaml.Flag{
						"flag",
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigUndefinedFlag,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := Check(td.nullCommandDescription, td.commandDescriptionMap, td.flagDescriptionMap)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}
