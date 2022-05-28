package main

import (
	"fmt"
	"github.com/terryhay/argtools/examples/example3/argTools"
	"os"
)

func main() {
	_, err := argTools.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("example3 error: %v", err.Error())
		os.Exit(int(err.Code()))
	}

	os.Exit(0)
}
