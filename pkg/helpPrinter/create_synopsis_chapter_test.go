package helpPrinter

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"testing"
)

func TestCreateSynopsisChapter(t *testing.T) {
	t.Parallel()

	var (
		appName = gofakeit.Name()
		chapter string

		expectedChapter string
	)

	t.Run("full_data", func(t *testing.T) {

		nullCommandRequiredFlag := argParserConfig.Flag(gofakeit.Name())
		nullCommandOptionalFlag := argParserConfig.Flag(gofakeit.Name())

		commandFlagWithSingleArgument := argParserConfig.Flag(gofakeit.Name())
		commandFlagDescriptionWithSingleArgument := &argParserConfig.FlagDescription{
			ArgDescription: &argParserConfig.ArgumentsDescription{
				AmountType:              argParserConfig.ArgAmountTypeSingle,
				SynopsisHelpDescription: gofakeit.Name(),
			},
		}

		commandFlagWithListArgument := argParserConfig.Flag(gofakeit.Name())
		commandFlagDescriptionWithListArgumentDefaultValue := gofakeit.Name()
		commandFlagDescriptionWithListArgumentAllowedValue := gofakeit.Name()
		commandFlagDescriptionWithListArgument := &argParserConfig.FlagDescription{
			ArgDescription: &argParserConfig.ArgumentsDescription{
				AmountType:              argParserConfig.ArgAmountTypeList,
				SynopsisHelpDescription: gofakeit.Name(),
				DefaultValues: []string{
					commandFlagDescriptionWithListArgumentDefaultValue,
				},
				AllowedValues: map[string]bool{
					commandFlagDescriptionWithListArgumentAllowedValue: true,
				},
			},
		}

		command := argParserConfig.Command(gofakeit.Name())

		namelessCommandDescription := argParserConfig.NewNamelessCommandDescription(
			0,
			"",
			&argParserConfig.ArgumentsDescription{
				SynopsisHelpDescription: gofakeit.Name(),
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
					SynopsisHelpDescription: gofakeit.Name(),
				},
			},
			commandFlagWithSingleArgument: commandFlagDescriptionWithSingleArgument,
			commandFlagWithListArgument:   commandFlagDescriptionWithListArgument,
		}

		expectedChapter = fmt.Sprintf(`[1mSYNOPSIS[0m
	[1m%s [0m [1m%s[0m [[1m%s[0m]
	[1m%s [0m[1m%s[0m [1m%s[0m [4m%s[0m [[1m%s[0m [4m%s[0m=%s [%s] [4m...[0m]

`,
			// first line
			appName,
			nullCommandRequiredFlag,
			nullCommandOptionalFlag,

			// second line
			appName,
			command,
			commandFlagWithSingleArgument,
			commandFlagDescriptionWithSingleArgument.GetArgDescription().GetSynopsisHelpDescription(),
			commandFlagWithListArgument,
			commandFlagDescriptionWithListArgument.GetArgDescription().GetSynopsisHelpDescription(),
			commandFlagDescriptionWithListArgumentDefaultValue,
			commandFlagDescriptionWithListArgumentAllowedValue,
		)
		chapter = CreateSynopsisChapter(appName, namelessCommandDescription, commandDescriptions, flagDescriptions)
		require.Equal(t,
			expectedChapter,
			chapter)
	})

	t.Run("no_commands", func(t *testing.T) {
		commandDescriptions := []*argParserConfig.CommandDescription{
			{},
		}

		expectedChapter = fmt.Sprintf(`[1mSYNOPSIS[0m
	[1m%s [0m

`, appName)
		chapter = CreateSynopsisChapter(appName, nil, commandDescriptions, nil)
		require.Equal(t,
			expectedChapter,
			chapter)
	})
}
