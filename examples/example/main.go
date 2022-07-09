package main

import (
	"fmt"
	"github.com/terryhay/argtools/examples/example/argTools"
	"os"
)

func main() {
	parsedData, err := argTools.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example.Argparser error: %v\n", err.Error())
		os.Exit(int(err.Code()))
	}

	switch parsedData.GetCommandID() {
	case argTools.CommandIDNamelessCommand:
		fmt.Println("dir")
	}

	os.Exit(0)
}
