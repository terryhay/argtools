package main

import (
	"bou.ke/monkey"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/internal/generator/argTools"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/osDecorator"
	"github.com/terryhay/argtools/internal/generator/osDecorator/osDecoratorMock"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
	"os"
	"testing"
)

func TestLogic(t *testing.T) {
	parsingErr := argtoolsError.NewError(argtoolsError.CodeUndefinedError, fmt.Errorf(gofakeit.Name()))
	configPath := parsedData.ArgValue(gofakeit.Name())

	getYAMLConfigErr := argtoolsError.NewError(argtoolsError.CodeConfigFlagIsNotUsedInCommands, fmt.Errorf(gofakeit.Name()))

	testData := []struct {
		caseName string

		argToolsParseFunc func(args []string) (res *parsedData.ParsedData, err *argtoolsError.Error)
		getYAMLConfigFunc func(configPath string) (*configYaml.Config, *argtoolsError.Error)
		osd               osDecorator.OSDecorator

		expectedErrCode argtoolsError.Code
	}{
		{
			caseName: "parsing_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return nil, parsingErr
			},
			osd: osDecoratorMock.NewMockOSDecorator(
				osDecoratorMock.MockOSDecoratorInit{
					Args: []string{},
				},
			),
			expectedErrCode: parsingErr.Code(),
		},
		{
			caseName: "get_config_path_arg_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return nil, nil
			},
			osd: osDecoratorMock.NewMockOSDecorator(
				osDecoratorMock.MockOSDecoratorInit{
					Args: []string{},
				},
			),
			expectedErrCode: argtoolsError.CodeParsedDataNilPointer,
		},
		{
			caseName: "get_generate_dir_path_arg_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return &parsedData.ParsedData{
						FlagData: map[argParserConfig.Flag]*parsedData.ParsedFlagData{
							argTools.FlagC: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										configPath,
									},
								},
							},
						},
					},
					nil
			},
			osd: osDecoratorMock.NewMockOSDecorator(
				osDecoratorMock.MockOSDecoratorInit{
					Args: []string{},
				},
			),
			expectedErrCode: argtoolsError.CodeParsedDataFlagDoesNotContainArgs,
		},
		{
			caseName: "get_yaml_config_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return &parsedData.ParsedData{
						FlagData: map[argParserConfig.Flag]*parsedData.ParsedFlagData{
							argTools.FlagC: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										configPath,
									},
								},
							},
							argTools.FlagO: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										parsedData.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*configYaml.Config, *argtoolsError.Error) {
				return nil, getYAMLConfigErr
			},
			osd: osDecoratorMock.NewMockOSDecorator(
				osDecoratorMock.MockOSDecoratorInit{
					Args: []string{},
				},
			),
			expectedErrCode: argtoolsError.CodeConfigFlagIsNotUsedInCommands,
		},
		{
			caseName: "extract_flag_descriptions_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return &parsedData.ParsedData{
						FlagData: map[argParserConfig.Flag]*parsedData.ParsedFlagData{
							argTools.FlagC: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										configPath,
									},
								},
							},
							argTools.FlagO: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										parsedData.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*configYaml.Config, *argtoolsError.Error) {
				return &configYaml.Config{
						FlagDescriptions: []*configYaml.FlagDescription{
							nil,
						},
					},
					nil
			},
			osd: osDecoratorMock.NewMockOSDecorator(
				osDecoratorMock.MockOSDecoratorInit{
					Args: []string{},
				},
			),
			expectedErrCode: argtoolsError.CodeUndefinedError,
		},
		{
			caseName: "extract_command_descriptions_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return &parsedData.ParsedData{
						FlagData: map[argParserConfig.Flag]*parsedData.ParsedFlagData{
							argTools.FlagC: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										configPath,
									},
								},
							},
							argTools.FlagO: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										parsedData.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*configYaml.Config, *argtoolsError.Error) {
				return &configYaml.Config{
						CommandDescriptions: []*configYaml.CommandDescription{
							nil,
						},
					},
					nil
			},
			osd: osDecoratorMock.NewMockOSDecorator(
				osDecoratorMock.MockOSDecoratorInit{
					Args: []string{},
				},
			),
			expectedErrCode: argtoolsError.CodeUndefinedError,
		},
		{
			caseName: "checking_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return &parsedData.ParsedData{
						FlagData: map[argParserConfig.Flag]*parsedData.ParsedFlagData{
							argTools.FlagC: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										configPath,
									},
								},
							},
							argTools.FlagO: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										parsedData.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*configYaml.Config, *argtoolsError.Error) {
				return &configYaml.Config{
						CommandDescriptions: []*configYaml.CommandDescription{
							{
								Command: configYaml.Command(gofakeit.Name()),
								RequiredFlags: []configYaml.Flag{
									configYaml.Flag(gofakeit.Color()),
								},
							},
						},
					},
					nil
			},
			osd: osDecoratorMock.NewMockOSDecorator(
				osDecoratorMock.MockOSDecoratorInit{
					Args: []string{},
				},
			),
			expectedErrCode: argtoolsError.CodeConfigFlagMustHaveDashInFront,
		},
		{
			caseName: "file_write_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return &parsedData.ParsedData{
						FlagData: map[argParserConfig.Flag]*parsedData.ParsedFlagData{
							argTools.FlagC: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										configPath,
									},
								},
							},
							argTools.FlagO: {
								ArgData: &parsedData.ParsedArgData{
									ArgValues: []parsedData.ArgValue{
										parsedData.ArgValue(gofakeit.Name()),
									},
								},
							},
						},
					},
					nil
			},
			getYAMLConfigFunc: func(configPath string) (*configYaml.Config, *argtoolsError.Error) {
				return &configYaml.Config{},
					nil
			},
			osd: osDecoratorMock.NewMockOSDecorator(osDecoratorMock.MockOSDecoratorInit{
				Args: []string{},
				IsNotExistFunc: func(err error) bool {
					return err != nil
				},
				StatFunc: func(path string) (os.FileInfo, error) {
					return nil, argtoolsError.NewError(argtoolsError.CodeUndefinedError, fmt.Errorf(gofakeit.Name()))
				},
			}),
			expectedErrCode: argtoolsError.CodeGeneratorInvalidGeneratePath,
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			err := logic(td.argToolsParseFunc, td.getYAMLConfigFunc, td.osd)
			require.Equal(t, td.expectedErrCode, err.Code())
		})
	}
}

func TestCrasher(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	assert.PanicsWithValue(
		t,
		"os.Exit called",
		func() {
			main()
		},
		"os.Exit was not called")
}
