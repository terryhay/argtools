package main

import (
	"fmt"
	"github.com/terryhay/argtools/examples/example/argTools"
	"github.com/terryhay/argtools/internal/osDecorator"
	"github.com/terryhay/argtools/pkg/parsedData"
	"os"
	"strings"
)

const (
	exitCodeSuccess uint = iota
	exitCodeConvertInt64Error
	exitCodeConvertFloat64Error
)

func main() {
	osd := osDecorator.NewOSDecorator()

	pd, err := argTools.Parse(osd.GetArgs())
	if err != nil {
		fmt.Printf("example.Argparser error: %v\n", err.Error())
		os.Exit(int(err.Code()))
	}

	osd.Exit(logic(pd))
}

func logic(pd *parsedData.ParsedData) (error, uint) {
	switch pd.GetCommandID() {
	case argTools.CommandIDNamelessCommand:
		var (
			builder      strings.Builder
			contain      bool
			err          error
			float64Value float64
			i            int
			int64Value   int64
			values       []parsedData.ArgValue
		)

		if values, contain = pd.GetFlagArgValues(argTools.FlagSl); contain {
			builder.WriteString(fmt.Sprintf("flag %v arguments:\n\t", argTools.FlagSl))

			for i = range values {
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%s ", values[i].ToString())))
			}
			builder.WriteString("\n")
		}

		if values, contain = pd.GetFlagArgValues(argTools.FlagIl); contain {
			builder.WriteString(fmt.Sprintf("flag %v arguments:\n\t", argTools.FlagIl))

			for i = range values {
				int64Value, err = values[i].ToInt64()
				if err != nil {
					return err, exitCodeConvertInt64Error
				}
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%d ", int64Value)))
			}
			builder.WriteString("\n")
		}

		if values, contain = pd.GetFlagArgValues(argTools.FlagFl); contain {
			builder.WriteString(fmt.Sprintf("flag %v arguments:\n\t", argTools.FlagFl))

			for i = range values {
				float64Value, err = values[i].ToFloat64()
				if err != nil {
					return err, exitCodeConvertFloat64Error
				}
				builder.WriteString(fmt.Sprintf(fmt.Sprintf("%f ", float64Value)))
			}
			builder.WriteString("\n")
		}

		fmt.Printf(builder.String())
	}

	return nil, exitCodeSuccess
}
