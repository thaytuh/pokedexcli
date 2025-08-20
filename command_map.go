package main

import (
	"fmt"
	"github.com/thaytuh/pokedexcli/internal/pokeapi"
)


func commandMap(cfg *config) error {
	location := pokeapi.Location{}
	
	var url string
	if cfg.nextLocationsURL == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.nextLocationsURL
	}

	location, err := pokeapi.GetLocation(url)
	if err != nil {
		return err
	}

	for _, result := range location.Results {
		fmt.Println(result.Name)
	}

	cfg.nextLocationsURL = &location.Next
	if location.Previous == nil {
		cfg.prevLocationsURL = nil
	} else {
		if prevURL, ok := location.Previous.(string); ok {
			cfg.prevLocationsURL = &prevURL
		} else {
			cfg.prevLocationsURL = nil
		}
	}

	return nil
}


func commandMapB(cfg *config) error {

	var url string
	if cfg.prevLocationsURL == nil {
		fmt.Printf("You're on the first page\n")
	} else {
		location := pokeapi.Location{}

		url = *cfg.prevLocationsURL
		
		location, err := pokeapi.GetLocation(url)
		if err != nil {
			return err
		}

		for _, result := range location.Results {
			fmt.Println(result.Name)
		}

		cfg.nextLocationsURL = &location.Next
		if location.Previous == nil {
			cfg.prevLocationsURL = nil
		} else {
			if prevURL, ok := location.Previous.(string); ok {
				cfg.prevLocationsURL = &prevURL
			} else {
				cfg.prevLocationsURL = nil
			}
		}
	}

	return nil
}