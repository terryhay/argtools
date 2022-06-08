package helpPrinter

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"sort"
	"strings"
)

const (
	descriptionChapterTitle = "\u001B[1mDESCRIPTION\u001B[0m\n\t%s\n"
	descriptionChapterLine  = "\n\t\u001B[1m%s\u001B[0m\t%s\n"
)

// CreateDescriptionChapter - creates description help chapter
func CreateDescriptionChapter(descriptionHelpInfo []string, flagDescriptions map[argParserConfig.Flag]*argParserConfig.FlagDescription) string {
	var (
		builder         strings.Builder
		flagDescription *argParserConfig.FlagDescription
	)

	builder.WriteString(fmt.Sprintf(descriptionChapterTitle, strings.Join(descriptionHelpInfo, "\n")))

	for _, flagStr := range getSortedFlagsForDescription(flagDescriptions) {
		flagDescription = flagDescriptions[argParserConfig.Flag(flagStr)]
		builder.WriteString(fmt.Sprintf(descriptionChapterLine, flagStr, flagDescription.GetDescriptionHelpInfo()))
	}

	return builder.String()
}

func getSortedFlagsForDescription(flagDescriptions map[argParserConfig.Flag]*argParserConfig.FlagDescription) (res []string) {
	if len(flagDescriptions) == 0 {
		return nil
	}
	res = make([]string, 0, len(flagDescriptions))
	for flag := range flagDescriptions {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}
