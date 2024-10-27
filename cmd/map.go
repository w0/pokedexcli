package cmd

import (
	"fmt"

	"github.com/w0/pokedexcli/internal/pokeapi"
)

func commandMap(config *config) {
	locations, _ := pokeapi.GetLocations(config.NextLocation)

	config.NextLocation = locations.Next
	config.PreviousLocation = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
}

func commandMapb(config *config) {
	locations, _ := pokeapi.GetLocations(config.PreviousLocation)

	config.NextLocation = locations.Next
	config.PreviousLocation = locations.Previous

	for _, location := range *&locations.Results {
		fmt.Println(location.Name)
	}
}
