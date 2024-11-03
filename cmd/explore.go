package cmd

import (
	"fmt"
)

func commandExplore(config *config, location ...string) {
	pokemons, err := config.Client.GetLocationDetail(location[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Exploring %s...\n", location[0])
	fmt.Println("Found Pokemon:")
	for _, mon := range pokemons.PokemonEncounters {
		fmt.Println(" -", mon.Pokemon.Name)
	}
}
