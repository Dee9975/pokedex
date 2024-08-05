package commands

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(config *PokedexConfig, args ...string) error {
	name := args[1]

	fmt.Printf("Catching %s\n", name)

	if rand.Intn(100) < 50 {
		fmt.Printf("You cought %s!\n", name)
		config.PokeBag = append(config.PokeBag, name)
		return nil
	}

	fmt.Printf("%s escaped!\n", name)

	return nil
}

func commandPokedex(config *PokedexConfig, args ...string) error {
	if len(args) == 1 {
		for _, pokemon := range config.PokeBag {
			fmt.Println(pokemon)
		}

		return nil
	}

	name := args[1]

	if contains(config.PokeBag, name) {
		fmt.Println(name + " is your selected pokemon")
		return nil
	}

	fmt.Println(name + " does not exist in your pokedex")

	return nil
}

func commandInspect(config *PokedexConfig, args ...string) error {
	if len(args) < 2 {
		return errors.New("not enough arguments")
	}

	name := args[1]

	if !contains(config.PokeBag, name) {
		return errors.New("pokemon not in your pokedex")
	}

	data, err := config.ApiClient.FetchPokemonData(name)

	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight: %d\n", data.Weight)
	fmt.Printf("Species: %s\n", data.Species.Name)
	fmt.Println("Stats:")
	for _, stat := range data.Stats {
		fmt.Printf("\t%s:%d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range data.Types {
		fmt.Printf("\t%s\n", t.Type.Name)
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
