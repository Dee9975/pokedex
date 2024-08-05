package pokedexApi

import (
	"encoding/json"
	"net/http"
	"pokedex/cache"
)

type PokedexClient struct {
	CacheManager *cache.Cache
	BseUrl       string
	HttpClient   http.Client
}

func parsePokemon(val []byte) (*PokemonData, error) {
	var data PokemonData

	err := json.Unmarshal(val, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func parsePokemonData(val []byte) (*PokemonAreaData, error) {
	var data PokemonAreaData

	err := json.Unmarshal(val, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func parseJson(val []byte) (*PokedexLocation, error) {
	var locations PokedexLocation

	err := json.Unmarshal(val, &locations)

	if err != nil {
		return nil, err
	}

	return &locations, nil
}
