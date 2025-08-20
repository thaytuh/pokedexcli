package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


func commandMap(cfg *config) error {
	type Location struct {
		Count 		int		 `json:"count"`
		Next 		string	 `json:"next"`
		Previous 	any 	 `json:"previous"`
		Results []struct {
			Name 	string	`json:"name"`
			URL 	string	`json:"url"`
		} `json:"results"`
	}

	location := Location{}
	
	var url string
	if cfg.nextLocationsURL == nil {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = *cfg.nextLocationsURL
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	err2 := json.Unmarshal(body, &location)
	if err2 != nil {
		fmt.Println(location)
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
	type Location struct {
		Count 		int		 `json:"count"`
		Next 		string	 `json:"next"`
		Previous 	any 	 `json:"previous"`
		Results []struct {
			Name 	string	`json:"name"`
			URL 	string	`json:"url"`
		} `json:"results"`
	}

	location := Location{}
	
	var url string
	if cfg.prevLocationsURL == nil {
		fmt.Printf("You're on the first page\n")
	} else {
		url = *cfg.prevLocationsURL
		
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		err2 := json.Unmarshal(body, &location)
		if err2 != nil {
			fmt.Println(location)
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