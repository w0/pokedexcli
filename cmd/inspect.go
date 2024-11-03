package cmd

import (
	"fmt"
)

func commandInspect(config *config, name ...string) {
	if val, ok := config.Pokemon[name[1]]; ok {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)
		fmt.Println("Stats:")

		for _, stat := range val.Stats {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")

		for _, pType := range val.Types {
			fmt.Printf(" - %s\n", pType.Type.Name)
		}

		return
	}

	fmt.Println("you have not caught that pokemon.")
}
