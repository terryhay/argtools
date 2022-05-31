package helpPrinter

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"strings"
)

// PrintHelpInfo - prints help information by ArgParserConfig object
func PrintHelpInfo(argParserConfig argParserConfig.ArgParserConfig) {
	builder := strings.Builder{}

	builder.WriteString(CreateNameChapter(
		argParserConfig.GetAppDescription().GetAppName(),
		argParserConfig.GetAppDescription().GetNameHelpInfo()))

	builder.WriteString(CreateSynopsisChapter(
		argParserConfig.GetAppDescription().GetAppName(),
		argParserConfig.GetNullCommandDescription(),
		argParserConfig.GetCommandDescriptions(),
		argParserConfig.GetFlagDescriptions()))

	builder.WriteString(CreateDescriptionChapter(
		argParserConfig.GetAppDescription().GetDescriptionHelpInfo(),
		argParserConfig.GetFlagDescriptions()))

	fmt.Println(builder.String())
}
