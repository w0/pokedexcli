package cmd

import (
	"fmt"

	"github.com/w0/pokedexcli/internal/pokecache"
)

func commandPokedex(config *config, cache *pokecache.Cache, name *string) {
	fmt.Println("Your Pokedex:")

	for _, mon := range config.Pokemon {
		fmt.Printf(" - %s\n", mon.Name)
	}
}
