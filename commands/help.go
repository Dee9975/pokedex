package commands

import "fmt"

func commandHelp(config *PokedexConfig, args ...string) error {
	c := GetCommands()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range c {
		fmt.Printf("\"%s\" %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
