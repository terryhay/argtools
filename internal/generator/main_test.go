package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/internal/generator/argTools"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/osDecorator"
	"github.com/terryhay/argtools/internal/generator/osDecorator/osDecoratorMock"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
	"io/ioutil"
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
			expectedErrCode: parsingErr.Code(),
		},
		{
			caseName: "get_config_path_arg_error",

			argToolsParseFunc: func(arg []string) (res *parsedData.ParsedData, err *argtoolsError.Error) {
				return nil, nil
			},
			expectedErrCode: argtoolsError.CodeUndefinedError,
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
			expectedErrCode: argtoolsError.CodeUndefinedError,
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
			expectedErrCode: argtoolsError.CodeConfigUndefinedFlag,
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

func dieOn(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

// catchStdOut - returns output to `os.Stdout` from `runnable` as string.
func catchStdOut(t *testing.T, runnable func()) string {
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()

	r, fakeStdout, err := os.Pipe()
	dieOn(err, t)
	os.Stdout = fakeStdout

	runnable()

	// need to close here, otherwise ReadAll never gets "EOF".
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)

	return string(newOutBytes)
}

/*func TestCrasher(t *testing.T) {
	defer func() {
		//require.Nil(t, recover())
	}()
	catchStdOut(t, func() {
		main()
	})
}//*/
