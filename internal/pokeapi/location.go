package pokeapi

import (
	"encoding/json"
	"net/http"
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

func GetLocations(URL *string) (Locations, error) {
	reqURL := "https://pokeapi.co/api/v2/location-area/"

	if URL != nil {
		reqURL = *URL
	}

	res, err := http.Get(reqURL)

	if err != nil {
		return Locations{}, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	loc := Locations{}

	err = decoder.Decode(&loc)

	if err != nil {
		return Locations{}, err
	}

	return loc, nil
}
