package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/w0/pokedexcli/cmd"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := cmd.GetCommands()

	config := cmd.GetConfig()

	for {
		fmt.Print("pokedex > ")

		scanner.Scan()

		input := scanner.Text()

		things := strings.Split(input, " ")

		if cmd, ok := commands[things[0]]; ok {
			cmd.Callback(&config, things...)
		}

	}
}
