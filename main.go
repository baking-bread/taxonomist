package main

import (
	"os"

	cmd "github.com/baking-bread/taxonomist/cmd/taxonomist"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
