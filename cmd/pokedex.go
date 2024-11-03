package cmd

import (
	"fmt"
)

func commandPokedex(config *config, name ...string) {
	fmt.Println("Your Pokedex:")

	for _, mon := range config.Pokemon {
		fmt.Printf(" - %s\n", mon.Name)
	}
}
