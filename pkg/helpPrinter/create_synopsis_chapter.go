package helpPrinter

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"sort"
	"strings"
)

// CreateSynopsisChapter - creates synopsis help chapter
func CreateSynopsisChapter(
	appName string,
	nullCommandDescription *argParserConfig.NullCommandDescription,
	commandDescriptions []*argParserConfig.CommandDescription,
	flagDescriptions map[argParserConfig.Flag]*argParserConfig.FlagDescription,
) string {

	var (
		builder         strings.Builder
		flagStr         string
		success         bool
		flagDescription *argParserConfig.FlagDescription
	)

	builder.WriteString("\u001B[1mSYNOPSIS\u001B[0m\n")

	if nullCommandDescription != nil {
		// app name part
		builder.WriteString(fmt.Sprintf(`	[1m%s [0m`, appName))

		// required flags part
		for _, flagStr = range getSortedFlags(nullCommandDescription.GetRequiredFlags()) {
			flagDescription, success = flagDescriptions[argParserConfig.Flag(flagStr)]
			if !success {
				continue
			}

			builder.WriteString(fmt.Sprintf(" \u001B[1m%s\u001B[0m", flagStr))
			builder.WriteString(fillUpArgumentsTemplatePart(flagDescription.GetArgDescription()))
		}

		// optional flags part
		for _, flagStr = range getSortedFlags(nullCommandDescription.GetOptionalFlags()) {
			flagDescription, success = flagDescriptions[argParserConfig.Flag(flagStr)]
			if !success {
				continue
			}

			builder.WriteString(fmt.Sprintf(" [\u001B[1m%s\u001B[0m", flagStr))
			builder.WriteString(fillUpArgumentsTemplatePart(flagDescription.GetArgDescription()))

			builder.WriteString("]")
		}

		builder.WriteString("\n")
	}

	for _, commandDescription := range commandDescriptions {
		// app name part
		builder.WriteString(fmt.Sprintf(`	[1m%s [0m`, appName))

		// command part
		if len(commandDescription.GetCommands()) > 0 {
			builder.WriteString(fmt.Sprintf(`[1m%s[0m`, strings.Join(getSortedCommands(commandDescription.GetCommands()), ", ")))
			builder.WriteString(fillUpArgumentsTemplatePart(commandDescription.GetArgDescription()))
		}

		// required flags part
		for _, flagStr = range getSortedFlags(commandDescription.GetRequiredFlags()) {
			flagDescription, success = flagDescriptions[argParserConfig.Flag(flagStr)]
			if !success {
				continue
			}

			builder.WriteString(fmt.Sprintf(" \u001B[1m%s\u001B[0m", flagStr))
			builder.WriteString(fillUpArgumentsTemplatePart(flagDescription.GetArgDescription()))
		}

		// optional flags part
		for _, flagStr = range getSortedFlags(commandDescription.GetOptionalFlags()) {
			flagDescription, success = flagDescriptions[argParserConfig.Flag(flagStr)]
			if !success {
				continue
			}

			builder.WriteString(fmt.Sprintf(" [\u001B[1m%s\u001B[0m", flagStr))
			builder.WriteString(fillUpArgumentsTemplatePart(flagDescription.GetArgDescription()))

			builder.WriteString("]")
		}

		builder.WriteString("\n")
	}
	builder.WriteString("\n")

	return builder.String()
}

func fillUpArgumentsTemplatePart(argDescription *argParserConfig.ArgumentsDescription) string {
	if argDescription == nil {
		return ""
	}

	var builder strings.Builder

	defaultValuesTemplatePart := ""
	if len(argDescription.GetDefaultValues()) > 0 {
		defaultValuesTemplatePart = fmt.Sprintf(`=%s`, strings.Join(argDescription.GetDefaultValues(), ", "))
	}

	allowedValuesTemplatePart := ""
	if len(argDescription.GetAllowedValues()) > 0 {
		allowedValuesTemplatePart = fmt.Sprintf(` [%s]`, strings.Join(getSortedStrings(argDescription.GetAllowedValues()), ", "))
	}

	switch argDescription.GetAmountType() {
	case argParserConfig.ArgAmountTypeSingle:
		builder.WriteString(fmt.Sprintf(` [4m%s[0m%s%s`,
			argDescription.GetSynopsisHelpDescription(),
			defaultValuesTemplatePart,
			allowedValuesTemplatePart))
	case argParserConfig.ArgAmountTypeList:
		builder.WriteString(fmt.Sprintf(` [4m%s[0m%s%s [4m...[0m`,
			argDescription.GetSynopsisHelpDescription(),
			defaultValuesTemplatePart,
			allowedValuesTemplatePart))
	default:
		return ""
	}

	return builder.String()
}

func getSortedCommands(commands map[argParserConfig.Command]bool) (res []string) {
	if len(commands) == 0 {
		return nil
	}
	res = make([]string, 0, len(commands))
	for command := range commands {
		res = append(res, string(command))
	}
	sort.Strings(res)

	return res
}

func getSortedFlags(groupFlagNameMap map[argParserConfig.Flag]bool) (res []string) {
	if len(groupFlagNameMap) == 0 {
		return nil
	}
	res = make([]string, 0, len(groupFlagNameMap))
	for flag := range groupFlagNameMap {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}

func getSortedStrings(strings map[string]bool) (res []string) {
	if len(strings) == 0 {
		return nil
	}
	res = make([]string, 0, len(strings))
	for s := range strings {
		res = append(res, s)
	}
	sort.Strings(res)

	return res
}
