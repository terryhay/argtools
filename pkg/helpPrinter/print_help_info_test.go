package helpPrinter

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/internal/test_tools"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"testing"
)

func TestPrintHelpInfo(t *testing.T) {
	t.Parallel()

	t.Run("empty_config", func(t *testing.T) {
		out := test_tools.CatchStdOut(func() {
			PrintHelpInfo(argParserConfig.ArgParserConfig{})
		})

		require.Equal(t, `[1mNAME[0m
	[1m[0m – 

[1mSYNOPSIS[0m

[1mDESCRIPTION[0m


`, out)
	})

	t.Run("simple_case", func(t *testing.T) {
		out := test_tools.CatchStdOut(func() {
			PrintHelpInfo(argParserConfig.NewArgParserConfig(
				argParserConfig.ApplicationDescription{
					AppName:      "appname",
					NameHelpInfo: "name help info",
				},
				nil,
				[]*argParserConfig.CommandDescription{
					{
						ID:                  1,
						DescriptionHelpInfo: "command id 1 description help info",
						Commands: map[argParserConfig.Command]bool{
							"command": true,
						},
						ArgDescription: &argParserConfig.ArgumentsDescription{
							AmountType:              argParserConfig.ArgAmountTypeSingle,
							SynopsisHelpDescription: "str",
						},
						RequiredFlags: map[argParserConfig.Flag]bool{
							"-rf1": true,
						},
						OptionalFlags: map[argParserConfig.Flag]bool{
							"-of1": true,
						},
					},
					{
						ID:                  2,
						DescriptionHelpInfo: "command id 2 description help info",
						Commands: map[argParserConfig.Command]bool{
							"longcommand": true,
						},
					},
				},
				nil,
				argParserConfig.NewNamelessCommandDescription(
					0,
					"nameless command description",
					nil,
					nil,
					nil,
				),
			))
		})

		require.Equal(t, `[1mNAME[0m
	[1mappname[0m – name help info

[1mSYNOPSIS[0m
	[1mappname[0m
	[1mappname command[0m [4mstr[0m [1m-rf1[0m [[1m-of1[0m]
	[1mappname longcommand[0m

[1mDESCRIPTION[0m

The commands are as follows:
	[1m<empty>[0m	nameless command description

	[1mcommand[0m	command id 1 description help info

	[1mlongcommand[0m
		command id 2 description help info

`, out)
	})
}
