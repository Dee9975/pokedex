package pokedexApi

import (
	"fmt"
	"io"
)

type PokemonData struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c *PokedexClient) FetchPokemonData(name string) (*PokemonData, error) {
	cachedPokemon := c.CacheManager.Get(name)

	if cachedPokemon != nil {
		return parsePokemon(cachedPokemon.Value)
	}

	url := fmt.Sprintf("%s/pokemon/%s", c.BseUrl, name)

	resp, err := c.HttpClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	c.CacheManager.Set(name, body)

	return parsePokemon(body)
}
