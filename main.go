package main

import (
	"os"

	"github.com/dev-ujin/gov-test/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}