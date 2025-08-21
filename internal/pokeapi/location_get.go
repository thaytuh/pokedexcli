package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (LocationInfo, error) {
	var locationInfo LocationInfo

	url := baseURL + "/location-area/" + locationName
	
	// check cache
	if cachedBytes, ok := c.locationCache.Get(url); ok {
		if err := json.Unmarshal(cachedBytes, &locationInfo); err == nil {
			return locationInfo, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationInfo{}, err
	}

	c.locationCache.Add(url, dat)

	err = json.Unmarshal(dat, &locationInfo)
	if err != nil {
		return LocationInfo{}, err
	}

	return locationInfo, nil
}