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

		descriptionHelpInfo        []string
		namelessCommandDescription argParserConfig.NamelessCommandDescription
		commandDescriptions        []*argParserConfig.CommandDescription
		flagDescriptions           map[argParserConfig.Flag]*argParserConfig.FlagDescription

		expected string
	}{
		{
			caseName:            "empty",
			descriptionHelpInfo: nil,
			flagDescriptions:    nil,

			expected: fmt.Sprintf(descriptionChapterTitle, ""),
		},
		{
			caseName:            "two_flags",
			descriptionHelpInfo: []string{randomDescriptionHelpInfo},
			flagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
				randomFlag: {
					DescriptionHelpInfo: randomFlagDescriptionHelpInfo,
				},
			},

			expected: fmt.Sprintf("%s\nThe flags are as follows:%s",
				fmt.Sprintf(descriptionChapterTitle, randomDescriptionHelpInfo),
				fmt.Sprintf(descriptionTwoLines, randomFlag, randomFlagDescriptionHelpInfo),
			),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			require.Equal(t, td.expected, CreateDescriptionChapter(
				td.descriptionHelpInfo,
				td.namelessCommandDescription,
				td.commandDescriptions,
				td.flagDescriptions))
		})
	}
}
