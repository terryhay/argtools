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

	randDescriptionHelpInfo := gofakeit.Name()
	randCommand := argParserConfig.Command(gofakeit.Color())
	randFlag := argParserConfig.Flag(gofakeit.Color())
	randFlagDescriptionHelpInfo := gofakeit.Name()

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
			descriptionHelpInfo: []string{randDescriptionHelpInfo},
			flagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
				randFlag: {
					DescriptionHelpInfo: randFlagDescriptionHelpInfo,
				},
			},

			expected: fmt.Sprintf(`[1mDESCRIPTION[0m
	%s

The flags are as follows:
	[1m%s[0m
		%s
`,
				randDescriptionHelpInfo,
				randFlag,
				randFlagDescriptionHelpInfo,
			),
		},
		{
			caseName:            "command_and_flag_descriptions",
			descriptionHelpInfo: []string{randDescriptionHelpInfo},
			commandDescriptions: []*argParserConfig.CommandDescription{
				{
					Commands: map[argParserConfig.Command]bool{randCommand: true},
				},
			},
			flagDescriptions: map[argParserConfig.Flag]*argParserConfig.FlagDescription{
				randFlag: {
					DescriptionHelpInfo: randFlagDescriptionHelpInfo,
				},
			},

			expected: fmt.Sprintf(`[1mDESCRIPTION[0m
	%s

The commands are as follows:
	[1m%s[0m
		

The flags are as follows:
	[1m%s[0m
		%s
`, randDescriptionHelpInfo, randCommand, randFlag, randFlagDescriptionHelpInfo),
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
