package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	var pokemonInfo Pokemon

	url := baseURL + "/pokemon/" + pokemonName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(dat, &pokemonInfo)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonInfo, nil
}