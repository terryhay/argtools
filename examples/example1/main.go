package main

import (
	"fmt"
	"github.com/terryhay/argtools/examples/example1/argTools"
	"os"
)

func main() {
	parsedData, err := argTools.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example1.Argparser error: %v", err.Error())
		os.Exit(int(err.Code()))
	}

	switch parsedData.GetCommandID() {
	case argTools.CommandIDNamelessCommand:
		fmt.Println("dir")
	}

	os.Exit(0)
}
