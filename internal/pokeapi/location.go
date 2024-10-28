package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/w0/pokedexcli/internal/pokecache"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationPokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocations(URL *string, cache *pokecache.Cache) (Locations, error) {
	reqURL := pokeapi_base + "/location-area"

	if URL != nil {
		reqURL = *URL
	}

	if val, ok := cache.Get(reqURL); ok {
		loc := Locations{}
		err := json.Unmarshal(val, &loc)

		if err != nil {
			return Locations{}, err
		}

		return loc, nil
	}

	res, err := http.Get(reqURL)

	if err != nil {
		return Locations{}, err
	}

	defer res.Body.Close()

	reqBody, err := io.ReadAll(res.Body)

	if err != nil {
		return Locations{}, err
	}

	cache.Add(reqURL, reqBody)

	loc := Locations{}
	err = json.Unmarshal(reqBody, &loc)

	if err != nil {
		return Locations{}, err
	}

	return loc, nil

}

func GetLocationDetail(param string, cache *pokecache.Cache) (LocationPokemon, error) {
	reqURL := pokeapi_base + "/location-area/" + param

	res, err := http.Get(reqURL)

	if err != nil {
		return LocationPokemon{}, err
	}
	defer res.Body.Close()

	reqBody, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationPokemon{}, err
	}

	cache.Add(reqURL, reqBody)

	mons := LocationPokemon{}

	err = json.Unmarshal(reqBody, &mons)

	if err != nil {
		return LocationPokemon{}, err
	}

	return mons, nil
}
