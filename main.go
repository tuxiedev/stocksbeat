package main

import (
	"os"

	"github.com/tuxiedev/stocksbeat/cmd"

	_ "github.com/tuxiedev/stocksbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
