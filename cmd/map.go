package cmd

import (
	"fmt"

	"github.com/w0/pokedexcli/internal/pokeapi"
	"github.com/w0/pokedexcli/internal/pokecache"
)

func commandMap(config *config, cache *pokecache.Cache, param *string) {
	locations, err := pokeapi.GetLocations(config.NextLocation, cache)

	if err != nil {
		fmt.Println(err)
		return
	}

	config.NextLocation = locations.Next
	config.PreviousLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}

func commandMapb(config *config, cache *pokecache.Cache, param *string) {
	locations, err := pokeapi.GetLocations(config.PreviousLocation, cache)

	if err != nil {
		fmt.Println(err)
		return
	}

	config.NextLocation = locations.Next
	config.PreviousLocation = locations.Previous

	for _, location := range *&locations.Results {
		fmt.Println(location.Name)
	}
}
