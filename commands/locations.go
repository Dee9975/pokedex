package commands

import (
	"errors"
	"fmt"
)

func commandMapf(config *PokedexConfig, args ...string) error {
	var url *string

	if config.LocationsNextUrl != nil {
		url = config.LocationsNextUrl
	}

	locations, err := config.ApiClient.FetchLocations(url)

	if err != nil {
		return err
	}

	config.LocationsPreviousUrl = locations.Previous
	config.LocationsNextUrl = locations.Next

	fmt.Println()
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(config *PokedexConfig, args ...string) error {
	var url *string

	if config.LocationsPreviousUrl == nil {
		return errors.New("no previous url specified")
	}

	url = config.LocationsPreviousUrl

	locations, err := config.ApiClient.FetchLocations(url)

	if err != nil {
		return err
	}

	config.LocationsPreviousUrl = locations.Previous
	config.LocationsNextUrl = locations.Next

	fmt.Println()
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	return nil
}

func commandExplore(config *PokedexConfig, args ...string) error {
	name := args[1]

	data, err := config.ApiClient.FetchPokemonInArea(name)

	if err != nil {
		return err
	}

	for _, encounter := range data.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
