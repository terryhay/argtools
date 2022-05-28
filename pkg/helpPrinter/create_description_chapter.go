package helpPrinter

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"sort"
	"strings"
)

// CreateDescriptionChapter - creates description help chapter
func CreateDescriptionChapter(descriptionHelpInfo string, flagDescriptions map[argParserConfig.Flag]*argParserConfig.FlagDescription) string {
	var (
		builder         strings.Builder
		flagDescription *argParserConfig.FlagDescription
		success         bool
	)

	builder.WriteString(fmt.Sprintf("\u001B[1mDESCRIPTION\u001B[0m\n\t%s\n", descriptionHelpInfo))

	for _, flagStr := range getSortedFlagsForDescription(flagDescriptions) {
		flagDescription, success = flagDescriptions[argParserConfig.Flag(flagStr)]
		if !success {
			continue
		}

		builder.WriteString(fmt.Sprintf("\n\t\u001B[1m%s\u001B[0m\t%s\n", flagStr, flagDescription.GetDescriptionHelpInfo()))
	}

	return builder.String()
}

func getSortedFlagsForDescription(flags map[argParserConfig.Flag]*argParserConfig.FlagDescription) (res []string) {
	if len(flags) == 0 {
		return nil
	}
	res = make([]string, 0, len(flags))
	for flag := range flags {
		res = append(res, string(flag))
	}
	sort.Strings(res)

	return res
}
