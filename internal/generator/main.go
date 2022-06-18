package main

import (
	"github.com/terryhay/argtools/internal/generator/argTools"
	"github.com/terryhay/argtools/internal/generator/configChecker"
	"github.com/terryhay/argtools/internal/generator/configDataExtractor"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/internal/generator/generate"
	"github.com/terryhay/argtools/internal/generator/osDecorator"
	"github.com/terryhay/argtools/internal/generator/writeFile"
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
) (err *argtoolsError.Error) {

	var argData *parsedData.ParsedData
	argData, err = argToolsParseFunc(osd.GetArgs()) // todo: write getter
	if err != nil {
		return err
	}

	var configYAMLFilePath parsedData.ArgValue
	configYAMLFilePath, err = argData.GetFlagArgValue(argTools.FlagC)
	if err != nil {
		return err
	}

	var generateDirPath parsedData.ArgValue
	generateDirPath, err = argData.GetFlagArgValue(argTools.FlagO)
	if err != nil {
		return err
	}

	var config *configYaml.Config
	config, err = getYAMLConfigFunc(string(configYAMLFilePath))
	if err != nil {
		return err
	}

	var flagDescriptions map[string]*configYaml.FlagDescription
	flagDescriptions, err = configDataExtractor.ExtractFlagDescriptionMap(config.GetFlagDescriptions())
	if err != nil {
		return err
	}

	var commandDescriptions map[string]*configYaml.CommandDescription
	commandDescriptions, err = configDataExtractor.ExtractCommandDescriptionMap(config.GetCommandDescriptions())
	if err != nil {
		return err
	}

	err = configChecker.Check(config.GetNamelessCommandDescription(), commandDescriptions, flagDescriptions)
	if err != nil {
		return err
	}

	err = writeFile.Write(osd, string(generateDirPath), generate.Generate(config, flagDescriptions))
	return err
}
