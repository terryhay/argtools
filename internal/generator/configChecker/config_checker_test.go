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

	flag := "-" + configYaml.Flag(gofakeit.Color())
	if len(flag) >= maxFlagLen {
		flag = flag[:maxFlagLen]
	}

	testData := []*struct {
		caseName                   string
		namelessCommandDescription *configYaml.NamelessCommandDescription
		commandDescriptionMap      map[configYaml.Command]*configYaml.CommandDescription
		flagDescriptionMap         map[configYaml.Flag]*configYaml.FlagDescription
		expectedErrorCode          argtoolsError.Code
	}{
		{
			caseName: "duplicate_flag_in_required_list",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Name()): {
					RequiredFlags: []configYaml.Flag{
						flag,
						flag,
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
						flag,
						flag,
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
						flag,
					},
					OptionalFlags: []configYaml.Flag{
						flag,
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
				flag: {
					Flag: flag,
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigFlagIsNotUsedInCommands,
		},

		{
			caseName: "undefined_required_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Name()): {
					RequiredFlags: []configYaml.Flag{
						flag,
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
						flag,
					},
				},
			},
			expectedErrorCode: argtoolsError.CodeConfigUndefinedFlag,
		},

		{
			caseName: "nameless_command_description_with_duplicate_required_flags",
			namelessCommandDescription: &configYaml.NamelessCommandDescription{
				RequiredFlags: []configYaml.Flag{
					flag,
					flag,
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_description_with_duplicate_optional_flags",
			namelessCommandDescription: &configYaml.NamelessCommandDescription{
				OptionalFlags: []configYaml.Flag{
					flag,
					flag,
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_description_with_duplicate_required_and_optional_flags",
			namelessCommandDescription: &configYaml.NamelessCommandDescription{
				RequiredFlags: []configYaml.Flag{
					flag,
				},
				OptionalFlags: []configYaml.Flag{
					flag,
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigContainsDuplicateFlags,
		},
		{
			caseName: "nameless_command_required_flag_does_not_have_dash_in_front",
			namelessCommandDescription: &configYaml.NamelessCommandDescription{
				RequiredFlags: []configYaml.Flag{
					flag[1:],
				},
				OptionalFlags: []configYaml.Flag{
					flag,
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigFlagMustHaveDashInFront,
		},
		{
			caseName: "nameless_command_optional_flag_has_russian_char",
			namelessCommandDescription: &configYaml.NamelessCommandDescription{
				RequiredFlags: []configYaml.Flag{
					flag,
				},
				OptionalFlags: []configYaml.Flag{
					"-йцукен",
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigIncorrectCharacterInFlagName,
		},
		{
			caseName: "command_with_too_long_required_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Color()): {
					RequiredFlags: []configYaml.Flag{
						flag + "d",
					},
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigIncorrectFlagLen,
		},
		{
			caseName: "command_with_empty_optional_flag",
			commandDescriptionMap: map[configYaml.Command]*configYaml.CommandDescription{
				configYaml.Command(gofakeit.Color()): {
					OptionalFlags: []configYaml.Flag{
						"",
					},
				},
			},

			expectedErrorCode: argtoolsError.CodeConfigIncorrectCharacterInFlagName,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := Check(td.namelessCommandDescription, td.commandDescriptionMap, td.flagDescriptionMap)
			require.NotNil(t, err)
			require.Equal(t, td.expectedErrorCode, err.Code())
		})
	}
}
