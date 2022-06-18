package configChecker

import (
	"fmt"
	"github.com/terryhay/argtools/internal/generator/configYaml"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"regexp"
	"strings"
)

const (
	maxFlagLen = 12
)

// Check checks command and flag descriptions for duplicates
func Check(
	namelessCommandDescription *configYaml.NamelessCommandDescription,
	commandDescriptions map[string]*configYaml.CommandDescription,
	flagDescriptions map[string]*configYaml.FlagDescription,
) *argtoolsError.Error {

	var (
		contain bool
		err     *argtoolsError.Error
	)
	if err = checkArgumentDescription(namelessCommandDescription.GetArgumentsDescription()); err != nil {
		return err
	}

	var allUsingFlags map[string]string
	allUsingFlags, err = getAllFlagsFromCommandDescriptions(namelessCommandDescription, commandDescriptions)
	if err != nil {
		return err
	}

	for flag, flagDescription := range flagDescriptions {
		if _, contain = allUsingFlags[flag]; !contain {
			return argtoolsError.NewError(argtoolsError.CodeConfigFlagIsNotUsedInCommands, fmt.Errorf(`configChecker.Check: flag "%s" is not found in command descriptions`, flag))
		}
		if err = checkArgumentDescription(flagDescription.GetArgumentsDescription()); err != nil {
			return err
		}
	}

	for flag, command := range allUsingFlags {
		if _, contain = flagDescriptions[flag]; !contain {
			return argtoolsError.NewError(argtoolsError.CodeConfigUndefinedFlag, fmt.Errorf(`configChecker.Check: command "%s" conains undefined flag "%s"`, command, flag))
		}
	}

	return nil
}

// CheckFlag checks if flag has dash in front and is not too long
func CheckFlag(checkFlagCharsFunc func(s string) bool, flag string) *argtoolsError.Error {
	if !checkFlagCharsFunc(flag) {
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

func checkArgumentDescription(argDescription *configYaml.ArgumentsDescription) *argtoolsError.Error {
	defaultValuesCount := len(argDescription.GetDefaultValues())
	if defaultValuesCount == 0 {
		return nil
	}

	if defaultValuesCount == 1 {
		if argDescription.GetAmountType() != argParserConfig.ArgAmountTypeSingle {
			return argtoolsError.NewError(
				argtoolsError.CodeConfigUnexpectedDefaultValue,
				fmt.Errorf(`configChecker.checkArgumentDescription: you need to set amount_type "single" if you want to use default_values logic`))
		}
	} else {
		if argDescription.GetAmountType() != argParserConfig.ArgAmountTypeList {
			return argtoolsError.NewError(
				argtoolsError.CodeConfigUnexpectedDefaultValue,
				fmt.Errorf(`configChecker.checkArgumentDescription: you need to set amount_type "list" if you want to use default_values logic`))
		}
	}

	allowedValuesCount := len(argDescription.GetAllowedValues())
	if allowedValuesCount == 0 {
		return nil
	}

	var allowed bool
	for i := 0; i < defaultValuesCount; i++ {
		allowed = false
		for j := 0; j < allowedValuesCount; j++ {
			if argDescription.GetDefaultValues()[i] == argDescription.GetAllowedValues()[j] {
				allowed = true
			}
		}

		if !allowed {
			return argtoolsError.NewError(
				argtoolsError.CodeConfigDefaultValueIsNotAllowed,
				fmt.Errorf(`configChecker.checkArgumentDescription: default value "%s" is not found in allowed values list: [%s]`,
					argDescription.GetDefaultValues()[i], strings.Join(argDescription.GetAllowedValues(), ", ")),
			)
		}
	}

	return nil
}

func getAllFlagsFromCommandDescriptions(
	namelessCommandDescription *configYaml.NamelessCommandDescription,
	commandDescriptionMap map[string]*configYaml.CommandDescription,
) (allUsingFlagMap map[string]string, err *argtoolsError.Error) {

	checkFlagCharsFunc := regexp.MustCompile(`^[a-zA-Z-]+$`).MatchString
	allUsingFlagMap = make(map[string]string, 2*len(commandDescriptionMap))
	checkDuplicateFlagMap := make(map[string]bool, 2*len(commandDescriptionMap))

	var (
		contain bool
		flag    string
	)

	// checking for nameless command
	const namelessCommand string = "NamelessCommand"
	for _, flag = range namelessCommandDescription.GetRequiredFlags() {
		if err = CheckFlag(checkFlagCharsFunc, flag); err != nil {
			return nil, err
		}

		if _, contain = checkDuplicateFlagMap[flag]; contain {
			return nil, argtoolsError.NewError(argtoolsError.CodeConfigContainsDuplicateFlags, fmt.Errorf(`getAllFlagsFromCommandDescriptions: command "%s" contains duplicate flag "%s"`, namelessCommand, flag))
		}
		checkDuplicateFlagMap[flag] = true

		allUsingFlagMap[flag] = namelessCommand
	}

	for _, flag = range namelessCommandDescription.GetOptionalFlags() {
		if err = CheckFlag(checkFlagCharsFunc, flag); err != nil {
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
		checkDuplicateFlagMap = map[string]bool{}

		if err = checkArgumentDescription(commandDescription.GetArgumentsDescription()); err != nil {
			return nil, err
		}

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
