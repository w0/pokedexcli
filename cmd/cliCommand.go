package cmd

type config struct {
	NextLocation     *string
	PreviousLocation *string
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*config)
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
	}
}

func GetConfig() config {
	return config{}
}
