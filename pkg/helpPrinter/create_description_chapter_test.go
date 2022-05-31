package helpPrinter

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"testing"
)

func TestCreateDescriptionChapter(t *testing.T) {
	t.Parallel()

	randomDescriptionHelpInfo := gofakeit.Name()
	randomFlag := argParserConfig.Flag(gofakeit.Name())
	randomFlagDescriptionHelpInfo := gofakeit.Name()

	testData := []struct {
		caseName string

		descriptionHelpInfo string
		flagDescriptions    map[argParserConfig.Flag]*argParserConfig.FlagDescription

		expected string
	}{
		{
			caseName:            "empty",
			descriptionHelpInfo: "",
			flagDescriptions:    nil,

			expected: fmt.Sprintf(descriptionChapterTitle, ""),
		},
		{
			caseName:            "two_flags",
			descriptionHelpInfo: randomDescriptionHelpInfo,
			flagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
				randomFlag: {
					DescriptionHelpInfo: randomFlagDescriptionHelpInfo,
				},
			},

			expected: fmt.Sprintf("%s%s",
				fmt.Sprintf(descriptionChapterTitle, randomDescriptionHelpInfo),
				fmt.Sprintf(descriptionChapterLine, randomFlag, randomFlagDescriptionHelpInfo),
			),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			require.Equal(t, td.expected, CreateDescriptionChapter(td.descriptionHelpInfo, td.flagDescriptions))
		})
	}
}
