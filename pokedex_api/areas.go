package pokedexApi

import (
	"fmt"
	"io"
)

type PokedexLocation struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonAreaData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *PokedexClient) FetchLocations(pageUrl *string) (*PokedexLocation, error) {
	url := fmt.Sprintf("%s/location-area", c.BseUrl)

	if pageUrl != nil {
		url = *pageUrl
	}

	cachedLocations := c.CacheManager.Get(url)

	if cachedLocations != nil {
		return parseJson(cachedLocations.Value)
	}

	resp, err := c.HttpClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	locations, err := parseJson(body)

	if err != nil {
		return nil, err
	}

	c.CacheManager.Set(url, body)

	return locations, nil
}

func (c *PokedexClient) FetchPokemonInArea(name string) (*PokemonAreaData, error) {
	cachedPokemon := c.CacheManager.Get(name)

	if cachedPokemon != nil {
		return parsePokemonData(cachedPokemon.Value)
	}

	url := fmt.Sprintf("%s/location-area/%s", c.BseUrl, name)

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

	return parsePokemonData(body)
}
