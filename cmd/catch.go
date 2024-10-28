package cmd

import (
	"fmt"

	"github.com/w0/pokedexcli/internal/pokeapi"
	"github.com/w0/pokedexcli/internal/pokecache"
)

func commandCatch(config *config, cache *pokecache.Cache, name *string) {
	pokemon, err := pokeapi.CatchPokemon(*name)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if pokeapi.AttemptCatch(pokemon, config.CatchAttempt) {
		config.Pokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.CatchAttempt = 1
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		config.CatchAttempt += 1
	}

}
