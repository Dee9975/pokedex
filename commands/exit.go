package commands

import "os"

func commandExit(config *PokedexConfig, args ...string) error {
	os.Exit(0)
	return nil
}
