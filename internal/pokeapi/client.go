package pokeapi

import (
	"io"
	"net/http"
	"fmt"
)

func FetchResource(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	return body, nil
}