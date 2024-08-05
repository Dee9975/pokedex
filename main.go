package main

import (
	"net/http"
	"pokedex/cache"
	"pokedex/commands"
	pokedexApi "pokedex/pokedex_api"
	"pokedex/repl"
	"time"
)

func main() {
	c := cache.NewCache()

	cfg := commands.PokedexConfig{
		ApiClient: pokedexApi.PokedexClient{
			BseUrl: "https://pokeapi.co/api/v2",
			HttpClient: http.Client{
				Timeout: 5 * time.Second,
			},
			CacheManager: c,
		},
		PokeBag: make([]string, 0),
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				cfg.ApiClient.CacheManager.ClearOld()
			}
		}
	}()

	repl.MakeRepl(cfg)

}
