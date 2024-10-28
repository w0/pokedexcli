package cmd

import (
	"github.com/w0/pokedexcli/internal/pokeapi"
	"github.com/w0/pokedexcli/internal/pokecache"
)

type config struct {
	CatchAttempt     int
	Pokemon          map[string]pokeapi.Pokemon
	NextLocation     *string
	PreviousLocation *string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config, *pokecache.Cache, *string)
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 locations areas in the Pokemon World.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays 20 previous areas in the Pokemon World.",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore the specified area for Pokemon!",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a Pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect captured Pokemon.",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Display all previously caught pokemon!",
			Callback:    commandPokedex,
		},
	}
}

func GetConfig() config {
	return config{
		CatchAttempt: 1,
		Pokemon:      make(map[string]pokeapi.Pokemon),
	}
}
