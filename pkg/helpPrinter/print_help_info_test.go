package helpPrinter

import (
	"github.com/stretchr/testify/require"
	"github.com/terryhay/argtools/pkg/argParserConfig"
	"io/ioutil"
	"os"
	"testing"
)

func dieOn(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

// catchStdOut returns output to `os.Stdout` from `runnable` as string.
func catchStdOut(t *testing.T, runnable func()) string {
	realStdout := os.Stdout
	defer func() { os.Stdout = realStdout }()

	r, fakeStdout, err := os.Pipe()
	dieOn(err, t)
	os.Stdout = fakeStdout

	runnable()

	// need to close here, otherwise ReadAll never gets "EOF".
	dieOn(fakeStdout.Close(), t)
	newOutBytes, err := ioutil.ReadAll(r)
	dieOn(err, t)
	dieOn(r.Close(), t)

	return string(newOutBytes)
}

func TestPrintHelpInfo(t *testing.T) {
	t.Parallel()

	t.Run("empty_config", func(t *testing.T) {
		out := catchStdOut(t, func() {
			PrintHelpInfo(argParserConfig.ArgParserConfig{})
		})

		require.Equal(t, `[1mNAME[0m
	[1m[0m – 

[1mSYNOPSIS[0m

[1mDESCRIPTION[0m
	

`, out)
	})

	t.Run("simple_case", func(t *testing.T) {
		out := catchStdOut(t, func() {
			PrintHelpInfo(argParserConfig.ArgParserConfig{
				NamelessCommandDescription: argParserConfig.NewNamelessCommandDescription(
					0,
					"nameless command description",
					nil,
					nil,
					nil,
				),
				CommandDescriptions: []*argParserConfig.CommandDescription{
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
			})
		})

		require.Equal(t, `[1mNAME[0m
	[1m[0m – 

[1mSYNOPSIS[0m
	[1m [0m
	[1m [0m[1mcommand[0m [4mstr[0m [1m-rf1[0m [[1m-of1[0m]
	[1m [0m[1mlongcommand[0m

[1mDESCRIPTION[0m
	

The commands are as follows:
	[1m<empty>[0m	nameless command description

	[1mcommand[0m	command id 1 description help info

	[1mlongcommand[0m
		command id 2 description help info

`, out)
	})
}
