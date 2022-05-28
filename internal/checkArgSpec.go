package internal

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"github.com/terryhay/argtools/pkg/argtoolsError"
)

func CheckArgSpec(commandDescriptions []*argParserConfig.CommandDescription, flagDescriptions map[argParserConfig.Flag]*argParserConfig.FlagDescription) *argtoolsError.Error {
	var (
		flag               argParserConfig.Flag
		flagSpec           *argParserConfig.FlagDescription
		commandDescription *argParserConfig.CommandDescription
		contain            bool
	)

	// check: all flag names in flagDescriptions must be used in command descriptions
	contain = false
	for flag, flagSpec = range flagDescriptions {
		if flagSpec == nil {
			return argtoolsError.NewError(argtoolsError.CodeFlagHasNilSpec, fmt.Errorf(`argtools: flag "%s" has nil flagSpec`, flag))
		}

		for _, commandDescription = range commandDescriptions {
			if _, contain = commandDescription.GetRequiredFlags()[flag]; contain {
				if string(flag[0]) != "-" {
					return argtoolsError.NewError(argtoolsError.CodeFlagMustHaveDashPrefix, fmt.Errorf(`argtools: flag "%s" is not a group flag, it must have a dash prefix`, flag))
				}
				break
			}
			if _, contain = commandDescription.GetOptionalFlags()[flag]; contain {
				if string(flag[0]) != "-" {
					return argtoolsError.NewError(argtoolsError.CodeFlagMustHaveDashPrefix, fmt.Errorf(`argtools: flag "%s" is not a group flag, it must have a dash prefix`, flag))
				}
				break
			}
		}

		if !contain {
			return argtoolsError.NewError(argtoolsError.CodeFlagMustBeInGroup, fmt.Errorf(`argtools: groupSpecSlice contains flag name "%s", this flag is not found in flagDescriptions`, flag))
		}
	}

	for _, commandDescription = range commandDescriptions {
		for flag = range commandDescription.RequiredFlags {
			if _, contain = flagDescriptions[flag]; !contain {
				return argtoolsError.NewError(argtoolsError.CodeFlagMustHaveSpec, fmt.Errorf(`argtools: flagDescriptions contains flag name "%s", this flag is not found in groupSpecSlice`, flag))
			}
		}

		for flag = range commandDescription.OptionalFlags {
			if _, contain = flagDescriptions[flag]; !contain {
				return argtoolsError.NewError(argtoolsError.CodeFlagMustHaveSpec, fmt.Errorf(`argtools: flagDescriptions contains flag name "%s", this flag is not found in groupSpecSlice`, flag))
			}
		}
	}

	return nil
}
