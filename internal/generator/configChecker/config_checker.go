package configChecker

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"regexp"
)

const (
	maxFlagLen = 12
)

// Check - checks command and flag descriptions for duplicates
func Check(
	namelessCommandDescription *configYaml.NamelessCommandDescription,
	commandDescriptions map[configYaml.Command]*configYaml.CommandDescription,
	flagDescriptions map[configYaml.Flag]*configYaml.FlagDescription,
) *argtoolsError.Error {

	allUsingFlags, err := getAllFlagsFromCommandDescriptions(namelessCommandDescription, commandDescriptions)
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

// CheckFlag - checks if flag has dash in front and is not too long
func CheckFlag(checkFlagCharsFunc func(s string) bool, flag configYaml.Flag) *argtoolsError.Error {
	if !checkFlagCharsFunc(string(flag)) {
		return argtoolsError.NewError(
			argtoolsError.CodeConfigIncorrectCharacterInFlagName,
			fmt.Errorf("configChecker.CheckFlag: flag \"%s\" must contain a dash in front and latin chars", flag))
	}

	flagLen := len(flag)
	if flagLen > maxFlagLen {
		return argtoolsError.NewError(
			argtoolsError.CodeConfigIncorrectFlagLen,
			fmt.Errorf("configChecker.CheckFlag: flag \"%s\" has len=%d, max len=%d", flag, flagLen, maxFlagLen))
	}

	if flag[:1] != "-" {
		return argtoolsError.NewError(
			argtoolsError.CodeConfigFlagMustHaveDashInFront,
			fmt.Errorf("configChecker.CheckFlag: flag \"%s\" must have a dash in front", flag))
	}

	return nil
}

func getAllFlagsFromCommandDescriptions(
	namelessCommandDescription *configYaml.NamelessCommandDescription,
	commandDescriptionMap map[configYaml.Command]*configYaml.CommandDescription,
) (allUsingFlagMap map[configYaml.Flag]configYaml.Command, err *argtoolsError.Error) {

	checkFlagCharsFunc := regexp.MustCompile(`^[a-zA-Z-]+$`).MatchString
	allUsingFlagMap = make(map[configYaml.Flag]configYaml.Command, 2*len(commandDescriptionMap))
	checkDuplicateFlagMap := make(map[configYaml.Flag]bool, 2*len(commandDescriptionMap))

	var (
		contain bool
		flag    configYaml.Flag
	)

	// checking for nameless command
	const namelessCommand configYaml.Command = "NamelessCommand"
	for _, flag = range namelessCommandDescription.GetRequiredFlags() {
		err = CheckFlag(checkFlagCharsFunc, flag)
		if err != nil {
			return nil, err
		}
		if _, contain = checkDuplicateFlagMap[flag]; contain {
			return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, namelessCommand, flag))
		}
		checkDuplicateFlagMap[flag] = true

		allUsingFlagMap[flag] = namelessCommand
	}

	for _, flag = range namelessCommandDescription.GetOptionalFlags() {
		err = CheckFlag(checkFlagCharsFunc, flag)
		if err != nil {
			return nil, err
		}
		if _, contain = checkDuplicateFlagMap[flag]; contain {
			return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, namelessCommand, flag))
		}
		checkDuplicateFlagMap[flag] = true

		allUsingFlagMap[flag] = namelessCommand
	}

	// checking for commands
	for _, commandDescription := range commandDescriptionMap {
		checkDuplicateFlagMap = map[configYaml.Flag]bool{}

		for _, flag = range commandDescription.GetRequiredFlags() {
			err = CheckFlag(checkFlagCharsFunc, flag)
			if err != nil {
				return nil, err
			}
			if _, contain = checkDuplicateFlagMap[flag]; contain {
				return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, commandDescription.GetCommand(), flag))
			}
			checkDuplicateFlagMap[flag] = true

			allUsingFlagMap[flag] = commandDescription.GetCommand()
		}

		for _, flag = range commandDescription.GetOptionalFlags() {
			err = CheckFlag(checkFlagCharsFunc, flag)
			if err != nil {
				return nil, err
			}
			if _, contain = checkDuplicateFlagMap[flag]; contain {
				return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, commandDescription.GetCommand(), flag))
			}
			checkDuplicateFlagMap[flag] = true

			allUsingFlagMap[flag] = commandDescription.GetCommand()
		}
	}

	return allUsingFlagMap, nil
}
