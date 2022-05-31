package configChecker

import (
	"github.com/brianvoe/gofakeit"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigCheckerCorrectResponse(t *testing.T) {
	t.Parallel()

	require.Nil(t, Check(nil, nil, nil))
}

func TestConfigCheckerErrors(t *testing.T) {
	t.Parallel()

	randFlag := configYaml.Flag(gofakeit.Name())

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
				configYaml.Command(gofakeit.Name()): {
					RequiredFlags: []configYaml.Flag{
						randFlag,
						randFlag,
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "duplicate_flag_in_optional_list",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Name()): {
					OptionalFlags: []configYaml.Flag{
						randFlag,
						randFlag,
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "duplicate_flag_in_required_and_optional_lists",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Name()): {
					RequiredFlags: []configYaml.Flag{
						randFlag,
					},
					OptionalFlags: []configYaml.Flag{
						randFlag,
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},

		{
			caseName: "unused_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Name()): {},
			},
			flagDescriptionMap: map[configYaml.Flag]*configYaml.FlagDescription{
				randFlag: {
					Flag: randFlag,
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigFlagIsNotUsedInCommands,
		},

		{
			caseName: "undefined_required_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Name()): {
					RequiredFlags: []configYaml.Flag{
						randFlag,
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigUndefinedFlag,
		},
		{
			caseName: "undefined_optional_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Name()): {
					OptionalFlags: []configYaml.Flag{
						randFlag,
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigUndefinedFlag,
		},

		{
			caseName: "null_command_description_with_duplicate_required_flags",
			nullCommandDescription: &configYaml.NullCommandDescription{
				RequiredFlags: []configYaml.Flag{
					randFlag,
					randFlag,
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "null_command_description_with_duplicate_optional_flags",
			nullCommandDescription: &configYaml.NullCommandDescription{
				OptionalFlags: []configYaml.Flag{
					randFlag,
					randFlag,
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "null_command_description_with_duplicate_required_and_optional_flags",
			nullCommandDescription: &configYaml.NullCommandDescription{
				RequiredFlags: []configYaml.Flag{
					randFlag,
				},
				OptionalFlags: []configYaml.Flag{
					randFlag,
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
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
