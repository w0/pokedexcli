package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/w0/pokedexcli/cmd"
	"github.com/w0/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := cmd.GetCommands()

	config := cmd.GetConfig()

	cache := pokecache.NewCache(5 * time.Second)

	for {
		fmt.Print("pokedex > ")

		scanner.Scan()

		input := scanner.Text()

		things := strings.Split(input, " ")

		junk := ""

		if len(things) == 1 {
			junk = ""
		} else {
			junk = things[1]
		}

		if cmd, ok := commands[things[0]]; ok {
			cmd.Callback(&config, &cache, &junk)
		}

	}
}
