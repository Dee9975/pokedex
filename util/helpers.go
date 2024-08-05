package util

import (
	"fmt"
	"strings"
)

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}

func InputMessage() {
	fmt.Print("Pokedex>")
}
