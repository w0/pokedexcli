package cmd

import (
	"fmt"

	"github.com/w0/pokedexcli/internal/pokeapi"
	"github.com/w0/pokedexcli/internal/pokecache"
)

func commandExplore(config *config, cache *pokecache.Cache, param *string) {
	pokemons, err := pokeapi.GetLocationDetail(*param, cache)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Exploring %s...\n", *param)
	fmt.Println("Found Pokemon:")
	for _, mon := range pokemons.PokemonEncounters {
		fmt.Println(" -", mon.Pokemon.Name)
	}
}
