package parsedData

import (
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"testing"
)

func TestParsedDataGetters(t *testing.T) {
	t.Parallel()

	var pointer *ParsedData

	{
		require.Equal(t, argParserConfig.CommandIDUndefined, pointer.GetCommandID())
		require.Equal(t, argParserConfig.CommandUndefined, pointer.GetCommand())
		require.Nil(t, pointer.GetAgrData())
		require.Nil(t, pointer.GetFlagData())
	}
	{
		pointer = NewParsedData(
			argParserConfig.CommandID(gofakeit.Uint32()),
			argParserConfig.Command(gofakeit.Name()),
			NewParsedArgData(nil),
			map[argParserConfig.Flag]*ParsedFlagData{
				argParserConfig.Flag(gofakeit.Name()): NewParsedFlagData(
					argParserConfig.Flag(gofakeit.Name()),
					NewParsedArgData(nil)),
			},
		)

		require.Equal(t, pointer.CommandID, pointer.GetCommandID())
		require.Equal(t, pointer.Command, pointer.GetCommand())
		require.Equal(t, pointer.ArgData, pointer.GetAgrData())
		require.Equal(t, pointer.FlagData, pointer.GetFlagData())
	}
	{
		pointer = NewParsedData(
			argParserConfig.CommandID(gofakeit.Uint32()),
			argParserConfig.Command(gofakeit.Name()),
			NewParsedArgData(nil),
			nil,
		)

		require.Nil(t, pointer.GetFlagData())
	}
}

func TestGetFlagArgValuesErrors(t *testing.T) {
	t.Parallel()

	flag := argParserConfig.Flag(gofakeit.Color())

	testData := []struct {
		caseName string

		parsedData      *ParsedData
		flag            argParserConfig.Flag
		expectedErrCode argtoolsError.Code
	}{
		{
			caseName:        "nil_pointer",
			expectedErrCode: argtoolsError.CodeParsedDataNilPointer,
		},
		{
			caseName:        "no_flag_data",
			parsedData:      &ParsedData{},
			expectedErrCode: argtoolsError.CodeParsedDataFlagDoesNotContainArgs,
		},
		{
			caseName: "no_flag_data",
			flag:     flag,
			parsedData: &ParsedData{
				FlagData: map[argParserConfig.Flag]*ParsedFlagData{
					flag: {
						Flag:    flag,
						ArgData: &ParsedArgData{},
					},
				},
			},
			expectedErrCode: argtoolsError.CodeNone,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName+"_GetFlagArgValues", func(t *testing.T) {
			v, err := td.parsedData.GetFlagArgValues(td.flag)
			require.Equal(t, 0, len(v))
			if td.expectedErrCode == argtoolsError.CodeNone {
				require.Nil(t, err)
			} else {
				require.Equal(t, td.expectedErrCode, err.Code())
			}

		})
	}
}

func TestGetFlagArgValueErrors(t *testing.T) {
	t.Parallel()

	flag := argParserConfig.Flag(gofakeit.Color())

	testData := []struct {
		caseName string

		parsedData      *ParsedData
		flag            argParserConfig.Flag
		expectedErrCode argtoolsError.Code
	}{
		{
			caseName:        "nil_pointer",
			expectedErrCode: argtoolsError.CodeParsedDataNilPointer,
		},
		{
			caseName:        "no_flag_data",
			parsedData:      &ParsedData{},
			expectedErrCode: argtoolsError.CodeParsedDataFlagDoesNotContainArgs,
		},
		{
			caseName: "no_flag_data",
			flag:     flag,
			parsedData: &ParsedData{
				FlagData: map[argParserConfig.Flag]*ParsedFlagData{
					flag: {
						Flag:    flag,
						ArgData: &ParsedArgData{},
					},
				},
			},
			expectedErrCode: argtoolsError.CodeParsedDataFlagDoesNotContainArgs,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName+"_GetFlagArgValues", func(t *testing.T) {
			v, err := td.parsedData.GetFlagArgValue(td.flag)
			require.Equal(t, 0, len(v))
			if td.expectedErrCode == argtoolsError.CodeNone {
				require.Nil(t, err)
				return
			}
			require.Equal(t, td.expectedErrCode, err.Code())

		})
	}
}

func TestGetFlagArgValue(t *testing.T) {
	t.Parallel()

	flag := argParserConfig.Flag(gofakeit.Color())
	value := ArgValue(gofakeit.Color())

	parsedData := &ParsedData{
		FlagData: map[argParserConfig.Flag]*ParsedFlagData{
			flag: {
				Flag: flag,
				ArgData: &ParsedArgData{
					ArgValues: []ArgValue{
						value,
					},
				},
			},
		},
	}
	{
		v, err := parsedData.GetFlagArgValues(flag)
		require.Nil(t, err)
		require.Equal(t, value, v[0])
	}
	{
		v, err := parsedData.GetFlagArgValue(flag)
		require.Nil(t, err)
		require.Equal(t, value, v)
	}
}
