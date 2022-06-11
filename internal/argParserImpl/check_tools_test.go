package argParserImpl

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
	"testing"
)

func TestCheckNoDashInFront(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		arg           string
		expectedValue bool
	}{
		{
			caseName:      "empty",
			expectedValue: true,
		},
		{
			caseName:      "dash_in_front",
			arg:           "-" + gofakeit.Color(),
			expectedValue: false,
		},
		{
			caseName:      "no_dash_in_front",
			arg:           gofakeit.Color(),
			expectedValue: true,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			require.Equal(t, td.expectedValue, checkNoDashInFront(td.arg))
		})
	}
}

func TestCheckParsedData(t *testing.T) {
	t.Parallel()

	testData := []struct {
		caseName string

		commandDescription *argParserConfig.CommandDescription
		data               *parsedData.ParsedData

		expectedErr *argtoolsError.Error
	}{
		{
			caseName: "nil_arguments",
		},
		{
			caseName: "required flag is not set",
			commandDescription: &argParserConfig.CommandDescription{
				RequiredFlags: map[argParserConfig.Flag]bool{
					argParserConfig.Flag(gofakeit.Color()): true,
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserRequiredFlagIsNotSet),
		},
		{
			caseName: "command arg is not set",
			commandDescription: &argParserConfig.CommandDescription{
				ArgDescription: &argParserConfig.ArgumentsDescription{
					AmountType: argParserConfig.ArgAmountTypeSingle,
				},
			},
			expectedErr: fakeError(argtoolsError.CodeArgParserCommandDoesNotContainArgs),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := checkParsedData(td.commandDescription, td.data)

			if td.expectedErr == nil {
				require.Nil(t, err)
				return
			}

			require.Equal(t, td.expectedErr.Code(), err.Code())
		})
	}
}

func TestIsValueAllowed(t *testing.T) {
	t.Parallel()

	value := gofakeit.Color()

	testData := []struct {
		caseName string

		argDescription *argParserConfig.ArgumentsDescription
		value          string

		expectedErr *argtoolsError.Error
	}{
		{
			caseName:    "nil_arguments",
			expectedErr: fakeError(argtoolsError.CodeArgParserCheckValueAllowabilityError),
		},
		{
			caseName: "no_allowed_values",
			argDescription: &argParserConfig.ArgumentsDescription{
				AmountType: argParserConfig.ArgAmountTypeNoArgs,
			},
		},
		{
			caseName: "value_is_not_allowed",
			argDescription: &argParserConfig.ArgumentsDescription{
				AmountType: argParserConfig.ArgAmountTypeNoArgs,
				AllowedValues: map[string]bool{
					value: true,
				},
			},
			value:       gofakeit.Color(),
			expectedErr: fakeError(argtoolsError.CodeArgParserArgValueIsNotAllowed),
		},
		{
			caseName: "value_is_allowed",
			argDescription: &argParserConfig.ArgumentsDescription{
				AmountType: argParserConfig.ArgAmountTypeNoArgs,
				AllowedValues: map[string]bool{
					value: true,
				},
			},
			value: value,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := isValueAllowed(td.argDescription, td.value)

			if td.expectedErr == nil {
				require.Nil(t, err)
				return
			}

			require.Equal(t, td.expectedErr.Code(), err.Code())
		})
	}
}
