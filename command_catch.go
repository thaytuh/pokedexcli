package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, params ...string) error {
	pokemon, err := cfg.pokeapiClient.GetPokemon(params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	attemptCatch := func(exp int64) bool {
		catchChance := math.Exp(-float64(exp) / 200.0)
		return rand.Float64() < catchChance
	}

	if attemptCatch(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		_, exists := cfg.pokedex[pokemon.Name]
		if !exists {
			cfg.pokedex[pokemon.Name] = pokemon
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}