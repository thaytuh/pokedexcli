package main

import (
	"fmt"
)

func commandPokedex(cfg *config, params ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range cfg.pokedex {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}
	return nil
}