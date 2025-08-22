package main

import (
	"fmt"
)

func commandInspect(cfg *config, params ...string) error {
	pokemon, exists := cfg.pokedex[params[0]]
	if exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, kind := range pokemon.Types {
			fmt.Printf("  - %s\n", kind.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	
	return nil
}