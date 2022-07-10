package helpPrinter

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/internal/testTools"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"testing"
)

func TestCreateSynopsisChapter(t *testing.T) {
	t.Parallel()

	var (
		appName = "appname"
		chapter string

		expectedChapter string
	)

	t.Run("full_data", func(t *testing.T) {

		nullCommandRequiredFlag := argParserConfig.Flag("-rf")
		nullCommandOptionalFlag := argParserConfig.Flag("-of")

		commandFlagWithSingleArgument := argParserConfig.Flag("-sa")
		commandFlagDescriptionWithSingleArgument := &argParserConfig.FlagDescription{
			ArgDescription: &argParserConfig.ArgumentsDescription{
				AmountType:              argParserConfig.ArgAmountTypeSingle,
				SynopsisHelpDescription: "arg",
			},
		}

		commandFlagWithListArgument := argParserConfig.Flag("-la")
		commandFlagDescriptionWithListArgumentDefaultValue := "val1"
		commandFlagDescriptionWithListArgumentAllowedValue := "val2"
		commandFlagDescriptionWithListArgument := &argParserConfig.FlagDescription{
			ArgDescription: &argParserConfig.ArgumentsDescription{
				AmountType:              argParserConfig.ArgAmountTypeList,
				SynopsisHelpDescription: "str",
				DefaultValues: []string{
					commandFlagDescriptionWithListArgumentDefaultValue,
				},
				AllowedValues: map[string]bool{
					commandFlagDescriptionWithListArgumentAllowedValue: true,
				},
			},
		}

		command := argParserConfig.Command("command")

		namelessCommandDescription := argParserConfig.NewNamelessCommandDescription(
			0,
			"nameless command description",
			&argParserConfig.ArgumentsDescription{
				SynopsisHelpDescription: "args",
			},
			map[argParserConfig.Flag]bool{
				nullCommandRequiredFlag: true,
			},
			map[argParserConfig.Flag]bool{
				nullCommandOptionalFlag: true,
			},
		)
		commandDescriptions := []*argParserConfig.CommandDescription{
			{
				Commands: map[argParserConfig.Command]bool{
					command: true,
				},
				RequiredFlags: map[argParserConfig.Flag]bool{
					commandFlagWithSingleArgument: true,
				},
				OptionalFlags: map[argParserConfig.Flag]bool{
					commandFlagWithListArgument: true,
				},
			},
		}
		flagDescriptions := map[argParserConfig.Flag]*argParserConfig.FlagDescription{
			nullCommandRequiredFlag: {
				ArgDescription: &argParserConfig.ArgumentsDescription{
					SynopsisHelpDescription: "arg1",
				},
			},
			commandFlagWithSingleArgument: commandFlagDescriptionWithSingleArgument,
			commandFlagWithListArgument:   commandFlagDescriptionWithListArgument,
		}

		chapter = CreateSynopsisChapter(appName, namelessCommandDescription, commandDescriptions, flagDescriptions)

		ok, msg := testTools.CheckSpaces(chapter)
		require.True(t, ok, msg)

		require.Equal(t,
			`[1mSYNOPSIS[0m
	[1mappname[0m [1m-rf[0m [[1m-of[0m]
	[1mappname command[0m [1m-sa[0m [4marg[0m [[1m-la[0m [4mstr[0m=val1 [val2] [4m...[0m]

`,
			chapter)
	})

	t.Run("no_commands", func(t *testing.T) {
		commandDescriptions := []*argParserConfig.CommandDescription{
			{},
		}

		expectedChapter = fmt.Sprintf(`[1mSYNOPSIS[0m
	[1m%s[0m

`, appName)

		chapter = CreateSynopsisChapter(appName, nil, commandDescriptions, nil)
		ok, msg := testTools.CheckSpaces(chapter)
		require.True(t, ok, msg)

		require.Equal(t, expectedChapter, chapter)
	})
}
