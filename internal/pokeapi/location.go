package pokeapi

import (
	"encoding/json"
	"io"
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

func (c *Client) GetLocations(URL *string) (Locations, error) {
	reqURL := pokeapi_base + "/location-area"

	if URL != nil {
		reqURL = *URL
	}

	if val, ok := c.cache.Get(reqURL); ok {
		loc := Locations{}
		err := json.Unmarshal(val, &loc)

		if err != nil {
			return Locations{}, err
		}

		return loc, nil
	}

	res, err := c.httpClient.Get(reqURL)

	if err != nil {
		return Locations{}, err
	}

	defer res.Body.Close()

	reqBody, err := io.ReadAll(res.Body)

	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(reqURL, reqBody)

	loc := Locations{}
	err = json.Unmarshal(reqBody, &loc)

	if err != nil {
		return Locations{}, err
	}

	return loc, nil

}

func (c *Client) GetLocationDetail(location string) (LocationPokemon, error) {
	reqURL := pokeapi_base + "/location-area/" + location

	res, err := c.httpClient.Get(reqURL)

	if err != nil {
		return LocationPokemon{}, err
	}
	defer res.Body.Close()

	reqBody, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationPokemon{}, err
	}

	c.cache.Add(reqURL, reqBody)

	mons := LocationPokemon{}

	err = json.Unmarshal(reqBody, &mons)

	if err != nil {
		return LocationPokemon{}, err
	}

	return mons, nil
}
