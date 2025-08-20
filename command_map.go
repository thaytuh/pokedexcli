package main

import (
	"errors"
	"fmt"
)


func commandMap(cfg *config) error {
	location, err := cfg.pokeapiClient.GetLocation(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = location.Next
	cfg.prevLocationsURL = location.Previous

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}


func commandMapB(cfg *config) error {

	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	location, err := cfg.pokeapiClient.GetLocation(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = location.Next
	cfg.prevLocationsURL = location.Previous

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}