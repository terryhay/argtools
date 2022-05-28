package configChecker

import (
	"argtools/internal/generator/configYaml"
	"argtools/pkg/argtoolsError"
	"fmt"
)

// Check - checks command and flag descriptions for duplicates
func Check(
	nullCommandDescription *configYaml.NullCommandDescription,
	commandDescriptions map[configYaml.Command]*configYaml.CommandDescription,
	flagDescriptions map[configYaml.Flag]*configYaml.FlagDescription,
) *argtoolsError.Error {

	allUsingFlags, err := getAllFlagsFromCommandDescriptions(nullCommandDescription, commandDescriptions)
	if err != nil {
		return err
	}

	var contain bool

	for flag := range flagDescriptions {
		if _, contain = allUsingFlags[flag]; !contain {
			return argtoolsError.NewError(argtoolsError.CodeConfigFlagIsNotUsedInCommands, fmt.Errorf(`configChecker.Check: flag "%s" is not found in command descriptions`, flag))
		}
	}

	for flag, command := range allUsingFlags {
		if _, contain = flagDescriptions[flag]; !contain {
			return argtoolsError.NewError(argtoolsError.CodeConfigUndefinedFlag, fmt.Errorf(`configChecker.Check: command "%s" conains undefined flag "%s"`, command, flag))
		}
	}

	return nil
}

func getAllFlagsFromCommandDescriptions(
	nullCommandDescription *configYaml.NullCommandDescription,
	commandDescriptionMap map[configYaml.Command]*configYaml.CommandDescription,
) (allUsingFlagMap map[configYaml.Flag]configYaml.Command, err *argtoolsError.Error) {

	allUsingFlagMap = make(map[configYaml.Flag]configYaml.Command, 2*len(commandDescriptionMap))
	checkDuplicateFlagMap := make(map[configYaml.Flag]bool, 2*len(commandDescriptionMap))

	var (
		contain bool
		flag    configYaml.Flag
	)

	// checking for null command
	const nullCommand configYaml.Command = "NullCommand"
	for _, flag = range nullCommandDescription.GetRequiredFlags() {
		if _, contain = checkDuplicateFlagMap[flag]; contain {
			return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, nullCommand, flag))
		}
		checkDuplicateFlagMap[flag] = true

		allUsingFlagMap[flag] = nullCommand
	}

	for _, flag = range nullCommandDescription.GetOptionalFlags() {
		if _, contain = checkDuplicateFlagMap[flag]; contain {
			return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, nullCommand, flag))
		}
		checkDuplicateFlagMap[flag] = true

		allUsingFlagMap[flag] = nullCommand
	}

	// checking for commands
	for _, commandDescription := range commandDescriptionMap {
		checkDuplicateFlagMap = map[configYaml.Flag]bool{}

		for _, flag = range commandDescription.GetRequiredFlags() {
			if _, contain = checkDuplicateFlagMap[flag]; contain {
				return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, commandDescription.GetCommand(), flag))
			}
			checkDuplicateFlagMap[flag] = true

			allUsingFlagMap[flag] = commandDescription.GetCommand()
		}

		for _, flag = range commandDescription.GetOptionalFlags() {
			if _, contain = checkDuplicateFlagMap[flag]; contain {
				return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, commandDescription.GetCommand(), flag))
			}
			checkDuplicateFlagMap[flag] = true

			allUsingFlagMap[flag] = commandDescription.GetCommand()
		}
	}

	return allUsingFlagMap, nil
}
