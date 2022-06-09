package argParserImpl

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()

	var (
		nullCommandID = argParserConfig.CommandID(gofakeit.Uint32())
		requiredFlag  = argParserConfig.Flag("-" + gofakeit.Color())
		optionalFlag  = argParserConfig.Flag("-" + gofakeit.Color())
		arg           = gofakeit.Color()
	)

	testData := []struct {
		caseName string

		config argParserConfig.ArgParserConfig
		args   []string

		expectedParsedData *parsedData.ParsedData
		expectedErr        *argtoolsError.Error
	}{
		{
			caseName:    "empty_config",
			expectedErr: fakeError(argtoolsError.CodeArgParserNullCommandUndefined),
		},
		{
			caseName:    "no_command_descriptions",
			args:        []string{arg},
			expectedErr: fakeError(argtoolsError.CodeArgParserIsNotInitialized),
		},
		{
			caseName: "unexpected_arg_error",
			args: []string{
				gofakeit.Color(),
			},
			config: argParserConfig.ArgParserConfig{
				CommandDescriptions: []*argParserConfig.CommandDescription{
					{
						Commands: map[argParserConfig.Command]bool{
							argParserConfig.Command(gofakeit.Color()): true,
						},
					},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeCantFindFlagNameInGroupSpec),
		},
		{
			caseName: "no_args_for_null_command",
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
				},
			},
			expectedParsedData: &parsedData.ParsedData{
				CommandID: nullCommandID,
			},
		},
		{
			caseName: "no_args_for_null_command_with_required_flag",
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					RequiredFlags: map[argParserConfig.Flag]bool{
						requiredFlag: true,
					},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "no_args_for_null_command_with_required_argument",
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					ArgDescription: &argParserConfig.ArgumentsDescription{
						AmountType: argParserConfig.ArgAmountTypeSingle,
					},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserCommandDoesNotContainArgs),
		},
		{
			caseName: "waste_arg_for_null_command",
			args:     []string{arg},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserUnexpectedArg),
		},
		{
			caseName: "arg_for_null_command_with_required_arg",
			args:     []string{arg},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					ArgDescription: &argParserConfig.ArgumentsDescription{
						AmountType: argParserConfig.ArgAmountTypeSingle,
					},
				},
			},
			expectedParsedData: &parsedData.ParsedData{
				CommandID: nullCommandID,
				ArgData: &parsedData.ParsedArgData{
					ArgValues: []parsedData.ArgValue{parsedData.ArgValue(arg)},
				},
			},
		},
		{
			caseName: "duplicate_flags",
			args: []string{
				string(requiredFlag),
				string(requiredFlag),
			},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					RequiredFlags: map[argParserConfig.Flag]bool{
						argParserConfig.Flag(arg): true,
					},
				},
				FlagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
					requiredFlag: {},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserDuplicateFlags),
		},
		{
			caseName: "arg_with_dash_in_front_in_single_argument_case",
			args: []string{
				string(requiredFlag),
				"-" + gofakeit.Color(),
			},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					RequiredFlags: map[argParserConfig.Flag]bool{
						argParserConfig.Flag(arg): true,
					},
				},
				FlagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
					requiredFlag: {
						ArgDescription: &argParserConfig.ArgumentsDescription{
							AmountType: argParserConfig.ArgAmountTypeSingle,
						},
					},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserDashInFrontOfArg),
		},
		{
			caseName: "arg_with_dash_in_front_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				"-" + gofakeit.Color(),
			},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					RequiredFlags: map[argParserConfig.Flag]bool{
						argParserConfig.Flag(arg): true,
					},
				},
				FlagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
					requiredFlag: {
						ArgDescription: &argParserConfig.ArgumentsDescription{
							AmountType: argParserConfig.ArgAmountTypeList,
						},
					},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserDashInFrontOfArg),
		},
		{
			caseName: "unexpected_flag_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				gofakeit.Color(),
				"-" + gofakeit.Color(),
			},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					RequiredFlags: map[argParserConfig.Flag]bool{
						argParserConfig.Flag(arg): true,
					},
				},
				FlagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
					requiredFlag: {
						ArgDescription: &argParserConfig.ArgumentsDescription{
							AmountType: argParserConfig.ArgAmountTypeList,
						},
					},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserUnexpectedFlag),
		},
		{
			caseName: "duplicated_flag_in_list_argument_case",
			args: []string{
				string(requiredFlag),
				arg,
				string(requiredFlag),
			},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					RequiredFlags: map[argParserConfig.Flag]bool{
						requiredFlag: true,
					},
				},
				FlagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
					requiredFlag: {
						ArgDescription: &argParserConfig.ArgumentsDescription{
							AmountType: argParserConfig.ArgAmountTypeList,
						},
					},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserDuplicateFlags),
		},
		{
			caseName: "correct_list_argument_case",
			args: []string{
				string(requiredFlag),
				arg,
				string(optionalFlag),
			},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					RequiredFlags: map[argParserConfig.Flag]bool{
						requiredFlag: true,
					},
					OptionalFlags: map[argParserConfig.Flag]bool{
						optionalFlag: true,
					},
				},
				FlagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
					requiredFlag: {
						ArgDescription: &argParserConfig.ArgumentsDescription{
							AmountType: argParserConfig.ArgAmountTypeList,
						},
					},
					optionalFlag: {},
				},
			},
			expectedParsedData: &parsedData.ParsedData{
				CommandID: nullCommandID,
				FlagData: map[argParserConfig.Flag]*parsedData.ParsedFlagData{
					requiredFlag: {
						Flag: requiredFlag,
						ArgData: &parsedData.ParsedArgData{
							ArgValues: []parsedData.ArgValue{
								parsedData.ArgValue(arg),
							},
						},
					},
					optionalFlag: {
						Flag: optionalFlag,
					},
				},
			},
		},
		{
			caseName: "failed_final_parsed_data_checking",
			args:     []string{arg},
			config: argParserConfig.ArgParserConfig{
				NullCommandDescription: &argParserConfig.NullCommandDescription{
					ID: nullCommandID,
					ArgDescription: &argParserConfig.ArgumentsDescription{
						AmountType: argParserConfig.ArgAmountTypeSingle,
					},
					RequiredFlags: map[argParserConfig.Flag]bool{
						requiredFlag: true,
					},
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserRequiredFlagIsNotSet),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			impl := NewCmdArgParserImpl(td.config)
			data, err := impl.Parse(td.args)

			if td.expectedErr != nil {
				require.Nil(t, data)
				require.NotNil(t, err)

				require.Equal(t, td.expectedErr.Code(), err.Code())
				return
			}

			require.Nil(t, err)
			require.Equal(t, td.expectedParsedData, data)
		})
	}
}

func fakeError(code argtoolsError.Code) *argtoolsError.Error {
	return argtoolsError.NewError(code, fmt.Errorf(""))
}
