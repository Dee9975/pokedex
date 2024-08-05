package repl

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/commands"
	"pokedex/util"
)

func MakeRepl(config commands.PokedexConfig) {
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Starting the pokedex")

	for {
		util.InputMessage()
		reader.Scan()
		text := reader.Text()
		words := util.CleanInput(text)

		commandName := words[0]

		command, exists := commands.GetCommands()[commandName]

		if exists {
			err := command.Callback(&config, words...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		fmt.Println("Unknown command")
	}
}
