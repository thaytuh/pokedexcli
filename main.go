package main

import (
	"time"

	"github.com/thaytuh/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient:  pokeClient,
		pokedex: 		make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
