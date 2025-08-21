package main

import (
	"fmt"
)

func commandExplore(cfg *config, params ...string) error {
	location, err := cfg.pokeapiClient.GetLocation(params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", location.Name)
	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}