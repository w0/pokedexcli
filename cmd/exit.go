package cmd

import (
	"os"
)

func commandExit(config *config, name ...string) {
	os.Exit(0)
}
