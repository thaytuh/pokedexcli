package pokeapi

import (
	"net/http"
	"time"

	"github.com/thaytuh/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient			*http.Client
	locationCache		pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		locationCache: pokecache.NewCache(5 * time.Second),
	}
}