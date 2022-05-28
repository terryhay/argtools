package main

import (
	"argtools/internal/generator/configChecker"
	"argtools/internal/generator/configDataExtractor"
	"argtools/internal/generator/configYaml"
	"argtools/internal/generator/generate"
	"argtools/internal/generator/writeFile"
	"argtools/pkg/argtoolsError"
	"flag"
	"fmt"
	"os"
)

func main() {
	configPath := flag.String("c", "", "yaml config file path")
	generateDirPath := flag.String("o", "", "create package path")
	flag.Parse()

	config, err := configYaml.GetConfig(*configPath)
	if err != nil {
		exitWithError(err)
	}

	var flagDescriptions map[configYaml.Flag]*configYaml.FlagDescription
	flagDescriptions, err = configDataExtractor.ExtractFlagDescriptionMap(config.GetFlagDescriptions())
	if err != nil {
		exitWithError(err)
	}

	var commandDescriptions map[configYaml.Command]*configYaml.CommandDescription
	commandDescriptions, err = configDataExtractor.ExtractCommandDescriptionMap(config.GetCommandDescriptions())
	if err != nil {
		exitWithError(err)
	}

	if err = configChecker.Check(config.GetNullCommandDescription(), commandDescriptions, flagDescriptions); err != nil {
		exitWithError(err)
	}

	err = writeFile.Write(generate.Generate(config, flagDescriptions), *generateDirPath)
	if err != nil {
		exitWithError(err)
	}

	os.Exit(0)
}

func exitWithError(err *argtoolsError.Error) {
	fmt.Println("argParser generator: " + err.Error())
	os.Exit(int(err.Code()))
}
