package main

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/argTools"
	"github.com/terryhay/argtools/internal/generator/configChecker"
	"github.com/terryhay/argtools/internal/generator/configDataExtractor"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/generate"
	"github.com/terryhay/argtools/internal/generator/writeFile"
	"github.com/terryhay/argtools/internal/osDecorator"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
)

func main() {
	osd := osDecorator.NewOSDecorator()
	osd.Exit(logic(argTools.Parse, configYaml.GetConfig, osd))
}

func logic(
	argToolsParseFunc func(args []string) (res *parsedData.ParsedData, err *argtoolsError.Error),
	getYAMLConfigFunc func(configPath string) (*configYaml.Config, *argtoolsError.Error),
	osd osDecorator.OSDecorator,
) (error, uint) {

	argData, err := argToolsParseFunc(osd.GetArgs())
	if err != nil {
		return err, err.Code().ToUint()
	}

	var (
		configYAMLFilePath parsedData.ArgValue
		contain            bool
	)
	configYAMLFilePath, contain = argData.GetFlagArgValue(argTools.FlagC)
	if !contain {
		err = argtoolsError.NewError(
			argtoolsError.CodeGeneratorNoRequiredFlag,
			fmt.Errorf("argTools.generator: can't get required flag \"%v\"", argTools.FlagC))
		return err, err.Code().ToUint()
	}

	var generateDirPath parsedData.ArgValue
	generateDirPath, contain = argData.GetFlagArgValue(argTools.FlagO)
	if !contain {
		err = argtoolsError.NewError(
			argtoolsError.CodeGeneratorNoRequiredFlag,
			fmt.Errorf("argTools.generator: can't get required flag \"%v\"", argTools.FlagO))
		return err, err.Code().ToUint()
	}

	var config *configYaml.Config
	config, err = getYAMLConfigFunc(string(configYAMLFilePath))
	if err != nil {
		return err, err.Code().ToUint()
	}

	var flagDescriptions map[string]*configYaml.FlagDescription
	flagDescriptions, err = configDataExtractor.ExtractFlagDescriptionMap(config.GetFlagDescriptions())
	if err != nil {
		return err, err.Code().ToUint()
	}

	var commandDescriptions map[string]*configYaml.CommandDescription
	commandDescriptions, err = configDataExtractor.ExtractCommandDescriptionMap(config.GetCommandDescriptions())
	if err != nil {
		return err, err.Code().ToUint()
	}

	err = configChecker.Check(config.GetNamelessCommandDescription(), commandDescriptions, flagDescriptions)
	if err != nil {
		return err, err.Code().ToUint()
	}

	err = writeFile.Write(osd, string(generateDirPath), generate.Generate(config, flagDescriptions))
	return err, err.Code().ToUint()
}
