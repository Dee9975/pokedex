package commands

import (
	pokedexApi "pokedex/pokedex_api"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(config *PokedexConfig, args ...string) error
}

type PokedexConfig struct {
	ApiClient            pokedexApi.PokedexClient
	LocationsNextUrl     *string
	LocationsPreviousUrl *string
	PokeBag              []string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "print this help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "exits the application",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "gets pokemon map",
			Callback:    commandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "gets previous page of the pokemon map",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "explores the area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "tries to catch a pokemon",
			Callback:    commandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "shows the pokedex",
			Callback:    commandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "inspects the pokemon in your pokedex",
			Callback:    commandInspect,
		},
	}
}
