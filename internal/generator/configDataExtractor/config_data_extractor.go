package configDataExtractor

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
)

// ExtractFlagDescriptionMap - extracts flag descriptions by flags from config object
func ExtractFlagDescriptionMap(flagDescriptions []*configYaml.FlagDescription) (flagDescriptionMap map[configYaml.Flag]*configYaml.FlagDescription, error *argtoolsError.Error) {
	descriptionCount := len(flagDescriptions)
	if descriptionCount == 0 {
		return nil, nil
	}
	flagDescriptionMap = make(map[configYaml.Flag]*configYaml.FlagDescription, descriptionCount)

	var contain bool
	for _, flagDescription := range flagDescriptions {
		if flagDescription == nil {
			return nil,
				argtoolsError.NewError(
					argtoolsError.CodeUndefinedError,
					fmt.Errorf(`ExtractFlagDescriptionMap: config object contains zero flag description pointer`))
		}

		if _, contain = flagDescriptionMap[flagDescription.GetFlag()]; contain {
			return nil,
				argtoolsError.NewError(
					argtoolsError.CodeConfigContainsDuplicateFlags,
					fmt.Errorf(`ExtractFlagDescriptionMap: yaml config contains duplicate flag "%s"`, flagDescription.GetFlag()))
		}

		flagDescriptionMap[flagDescription.GetFlag()] = flagDescription
	}

	return flagDescriptionMap, nil
}

// ExtractCommandDescriptionMap - extracts command descriptions by commands from config object
func ExtractCommandDescriptionMap(commandDescriptions []*configYaml.CommandDescription) (commandDescriptionMap map[configYaml.Command]*configYaml.CommandDescription, error *argtoolsError.Error) {
	descriptionCount := len(commandDescriptions)
	if descriptionCount == 0 {
		return nil, nil
	}
	commandDescriptionMap = make(map[configYaml.Command]*configYaml.CommandDescription, descriptionCount)
	checkDuplicationsMap := make(map[configYaml.Command]bool, descriptionCount)

	var contain bool
	for _, commandDescription := range commandDescriptions {
		if commandDescription == nil {
			return nil,
				argtoolsError.NewError(
					argtoolsError.CodeUndefinedError,
					fmt.Errorf(`ExtractFlagDescriptionMap: config object contains zero command description pointer`))
		}

		if _, contain = checkDuplicationsMap[commandDescription.GetCommand()]; contain {
			return nil,
				argtoolsError.NewError(
					argtoolsError.CodeConfigContainsDuplicateCommands,
					fmt.Errorf(`ExtractFlagDescriptionMap: yaml config contains duplicate flag "%s"`, commandDescription.GetCommand()))
		}
		checkDuplicationsMap[commandDescription.GetCommand()] = true
		for _, command := range commandDescription.GetAdditionalCommands() {
			if _, contain = checkDuplicationsMap[command]; contain {
				return nil,
					argtoolsError.NewError(
						argtoolsError.CodeConfigContainsDuplicateCommands,
						fmt.Errorf(`ExtractFlagDescriptionMap: yaml config contains duplicate flag "%s"`, command))
			}
			checkDuplicationsMap[command] = true
		}

		commandDescriptionMap[commandDescription.GetCommand()] = commandDescription
	}

	return commandDescriptionMap, nil
}
