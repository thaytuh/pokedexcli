package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(pageURL *string) (Location, error) {
	var location Location

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	
	// check cache
	if cachedBytes, ok := c.locationCache.Get(url); ok {
		if err := json.Unmarshal(cachedBytes, &location); err == nil {
			return location, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	c.locationCache.Add(url, dat)

	err = json.Unmarshal(dat, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}