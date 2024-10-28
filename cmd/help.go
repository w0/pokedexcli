package cmd

import (
	"fmt"

	"github.com/w0/pokedexcli/internal/pokecache"
)

func commandHelp(config *config, cache *pokecache.Cache, param *string) {
	commandMap := GetCommands()
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, value := range commandMap {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}

	fmt.Println()
}
