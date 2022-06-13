package helpPrinter

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"strings"
)

const (
	descriptionChapterTitle = "\u001B[1mDESCRIPTION\u001B[0m\n\t%s\n"

	commandDescriptionsSubtitle = "\nThe commands are as follows:"
	descriptionLine             = "\n\t\u001B[1m%s\u001B[0m\t%s\n"
	descriptionTwoLines         = "\n\t\u001B[1m%s\u001B[0m\n\t\t%s\n"

	flagDescriptionsSubtitle = "\nThe flags are as follows:"

	namelessCommandDescriptionName = "<empty>"
)

const tabLen = 7

// CreateDescriptionChapter - create
//s description help chapter
func CreateDescriptionChapter(
	descriptionHelpInfo []string,
	namelessCommandDescription argParserConfig.NamelessCommandDescription,
	commandDescriptions []*argParserConfig.CommandDescription,
	flagDescriptions map[argParserConfig.Flag]*argParserConfig.FlagDescription,
) string {

	var (
		builder         strings.Builder
		callNames       string
		flagDescription *argParserConfig.FlagDescription
		usingPattern    string
	)

	builder.WriteString(fmt.Sprintf(descriptionChapterTitle, strings.Join(descriptionHelpInfo, "\n\n\t")))

	if len(commandDescriptions) > 0 {
		builder.WriteString(commandDescriptionsSubtitle)

		if namelessCommandDescription != nil {
			builder.WriteString(fmt.Sprintf(descriptionLine,
				namelessCommandDescriptionName,
				namelessCommandDescription.GetDescriptionHelpInfo()))
		}

		for i := range commandDescriptions {
			callNames = strings.Join(getSortedCommands(commandDescriptions[i].GetCommands()), ", ")

			usingPattern = descriptionLine
			if len(callNames) > tabLen {
				usingPattern = descriptionTwoLines
			}

			builder.WriteString(fmt.Sprintf(usingPattern,
				callNames,
				commandDescriptions[i].GetDescriptionHelpInfo()))
		}
	}

	if len(flagDescriptions) > 0 {
		builder.WriteString(flagDescriptionsSubtitle)

		for _, callNames = range getSortedFlagsForDescription(flagDescriptions) {

			usingPattern = descriptionLine
			if len(callNames) > tabLen {
				usingPattern = descriptionTwoLines
			}

			flagDescription = flagDescriptions[argParserConfig.Flag(callNames)]
			builder.WriteString(fmt.Sprintf(usingPattern, callNames, flagDescription.GetDescriptionHelpInfo()))
		}
	}

	return builder.String()
}
