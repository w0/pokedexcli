package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/w0/pokedexcli/cmd"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := cmd.GetCommands()

	for {
		fmt.Print("pokedex > ")

		scanner.Scan()

		input := scanner.Text()

		if cmd, ok := commands[input]; ok {
			cmd.Callback()
		}

		fmt.Println()

	}
}
