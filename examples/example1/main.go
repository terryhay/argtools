package main

import (
	"argtools/examples/example1/argParser"
	"fmt"
	"os"
)

func main() {
	parsedData, err := argParser.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example1.Argparser error: %v", err.Error())
		os.Exit(int(err.Code()))
	}

	switch parsedData.GetCommandID() {
	case argParser.CommandIDNullCommand:
		fmt.Println("dir")
	}

	os.Exit(0)
}
