package helpPrinter

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"sort"
	"strings"
)

const (
	synopsisChapterTitle = "\u001B[1mSYNOPSIS\u001B[0m\n"
)

// CreateSynopsisChapter - creates synopsis help chapter
func CreateSynopsisChapter(
	appName string,
	nullCommandDescription *argParserConfig.NamelessCommandDescription,
	commandDescriptions []*argParserConfig.CommandDescription,
	flagDescriptions map[argParserConfig.Flag]*argParserConfig.FlagDescription,
) string {

	var (
		builder         strings.Builder
		flagStr         string
		flagDescription *argParserConfig.FlagDescription
		joinedString    string
	)

	builder.WriteString(synopsisChapterTitle)

	if nullCommandDescription != nil {
		// app name part
		builder.WriteString(fmt.Sprintf(`	[1m%s [0m`, appName))

		// required flags part
		for _, flagStr = range getSortedFlags(nullCommandDescription.GetRequiredFlags()) {
			flagDescription = flagDescriptions[argParserConfig.Flag(flagStr)]

			builder.WriteString(fmt.Sprintf(" \u001B[1m%s\u001B[0m", flagStr))
			builder.WriteString(fillUpArgumentsTemplatePart(flagDescription.GetArgDescription()))
		}

		// optional flags part
		for _, flagStr = range getSortedFlags(nullCommandDescription.GetOptionalFlags()) {
			flagDescription = flagDescriptions[argParserConfig.Flag(flagStr)]

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
		joinedString = strings.Join(getSortedCommands(commandDescription.GetCommands()), ", ")
		if len(joinedString) > 0 {
			builder.WriteString(fmt.Sprintf(`[1m%s[0m`, joinedString))
			builder.WriteString(fillUpArgumentsTemplatePart(commandDescription.GetArgDescription()))
		}

		// required flags part
		for _, flagStr = range getSortedFlags(commandDescription.GetRequiredFlags()) {
			flagDescription = flagDescriptions[argParserConfig.Flag(flagStr)]

			builder.WriteString(fmt.Sprintf(" \u001B[1m%s\u001B[0m", flagStr))
			builder.WriteString(fillUpArgumentsTemplatePart(flagDescription.GetArgDescription()))
		}

		// optional flags part
		for _, flagStr = range getSortedFlags(commandDescription.GetOptionalFlags()) {
			flagDescription = flagDescriptions[argParserConfig.Flag(flagStr)]

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
	joinedString := strings.Join(getSortedStrings(argDescription.GetAllowedValues()), ", ")
	if len(joinedString) > 0 {
		allowedValuesTemplatePart = fmt.Sprintf(` [%s]`, joinedString)
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
