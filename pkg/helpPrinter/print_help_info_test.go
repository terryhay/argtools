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

// catchStdOut - returns output to `os.Stdout` from `runnable` as string.
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

}
