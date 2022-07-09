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

			expected: descriptionChapterTitle + "\n",
		},
		{
			caseName:            "two_flags",
			descriptionHelpInfo: []string{randomDescriptionHelpInfo},
			flagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
				randomFlag: {
					DescriptionHelpInfo: randomFlagDescriptionHelpInfo,
				},
			},

			expected: fmt.Sprintf(`[1mDESCRIPTION[0m
	%s

The flags are as follows:
	[1m%s[0m
		Nickolas Emard
`,
				randomDescriptionHelpInfo,
				randomFlag,
			),
		},
	}

	for _, td := range testData {
		t.Run(td.caseName, func(t *testing.T) {
			actual := CreateDescriptionChapter(
				td.descriptionHelpInfo,
				td.namelessCommandDescription,
				td.commandDescriptions,
				td.flagDescriptions)
			require.Equal(t, td.expected, actual)
		})
	}
}
