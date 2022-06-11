package argParserImpl

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"github.com/terryhay/argtools/pkg/parsedData"
	"strings"
)

func checkNoDashInFront(arg string) bool {
	if len(arg) == 0 {
		return true
	}
	return arg[:1] != "-"
}

func checkParsedData(
	usingCommandDescription *argParserConfig.CommandDescription,
	data *parsedData.ParsedData,
) *argtoolsError.Error {
	var (
		argDescription *argParserConfig.ArgumentsDescription
		contain        bool
		flag           argParserConfig.Flag
		parsedFlagData = data.GetFlagData()
	)

	// check if all required flags is set
	for flag = range usingCommandDescription.GetRequiredFlags() {
		if _, contain = parsedFlagData[flag]; !contain {
			return argtoolsError.NewError(
				argtoolsError.CodeArgParserRequiredFlagIsNotSet,
				fmt.Errorf("CmdArgParser.checkParsedData: required flag is not set: %s", flag))
		}
	}

	// check command arguments
	argDescription = usingCommandDescription.GetArgDescription()
	if argDescription.GetAmountType() != argParserConfig.ArgAmountTypeNoArgs {
		if data.GetAgrData() == nil {
			return argtoolsError.NewError(
				argtoolsError.CodeArgParserCommandDoesNotContainArgs,
				fmt.Errorf("CmdArgParser.checkParsedData: command arg is not set: %s", flag))
		}
	}

	return nil
}

func isValueAllowed(argDescription *argParserConfig.ArgumentsDescription, value string) *argtoolsError.Error {
	if argDescription == nil {
		return argtoolsError.NewError(
			argtoolsError.CodeArgParserCheckValueAllowabilityError,
			fmt.Errorf("isValueAllowed: try to check a value \"%s\" allowability by nil pointer", value))
	}

	if len(argDescription.GetAllowedValues()) == 0 {
		return nil
	}

	if _, allow := argDescription.GetAllowedValues()[value]; !allow {
		allowedValuesSlice := make([]string, 0, len(argDescription.GetAllowedValues()))
		for allowedValue := range argDescription.GetAllowedValues() {
			allowedValuesSlice = append(allowedValuesSlice, allowedValue)
		}

		return argtoolsError.NewError(
			argtoolsError.CodeArgParserArgValueIsNotAllowed,
			fmt.Errorf(`isValueAllowed: value "%s" is not found in list allowed values: [%s]`,
				value, strings.Join(allowedValuesSlice, ", ")))
	}

	return nil
}
