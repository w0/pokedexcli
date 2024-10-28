package cmd

import (
	"os"

	"github.com/w0/pokedexcli/internal/pokecache"
)

func commandExit(config *config, cache *pokecache.Cache, param *string) {
	os.Exit(0)
}
