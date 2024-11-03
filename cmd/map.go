package cmd

import (
	"fmt"
)

func commandMap(config *config, location ...string) {
	locations, err := config.Client.GetLocations(config.NextLocation)
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

func commandMapb(config *config, location ...string) {
	locations, err := config.Client.GetLocations(config.PreviousLocation)

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
