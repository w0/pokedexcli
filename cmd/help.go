package cmd

import "fmt"

func commandHelp(config *config) {
	commandMap := GetCommands()
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, value := range commandMap {
		fmt.Printf("%s: %s\n", value.Name, value.Description)
	}

	fmt.Println()
}
