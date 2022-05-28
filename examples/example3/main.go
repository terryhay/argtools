package main

import (
	"argtools/examples/example3/argTools"
	"fmt"
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
