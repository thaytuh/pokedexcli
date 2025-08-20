package pokeapi

import (
	"encoding/json"
)

type Location struct {
	Count 		int		 `json:"count"`
	Next 		string	 `json:"next"`
	Previous 	any 	 `json:"previous"`
	Results []struct {
		Name 	string	`json:"name"`
		URL 	string	`json:"url"`
	} `json:"results"`
}

func GetLocation(url string) (Location, error) {
	body, err := FetchResource(url)
	if err != nil {
		return Location{}, err
	}
	var location Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}