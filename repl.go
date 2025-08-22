package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thaytuh/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient	 pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex 		 map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var commandParams []string
		if len(words) > 1 {
			commandParams = words[1:]
		}
		

		if command, exists := getCommands()[commandName]; exists {
			err := command.callback(cfg, commandParams...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("Unknown command\n")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}


type cliCommand struct {
	name			string
	description 	string
	callback 		func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: 			"help",
			description: 	"Displays a help message",
			callback: 		commandHelp,
		},
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback: 		commandExit,
		},
		"map": {
			name: 			"map",
			description:	"Displays the next 20 areas",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Displays the previous 20 areas",
			callback: 		commandMapB,
		},
		"explore": {
			name:			"explore <area-name>",
			description: 	"Lists the Pokemon available at the area specified",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch <pokemon-name>",
			description: 	"Attempts to catch the specified Pokemon",
			callback: 		commandCatch,
		},
		"inspect": {
			name: 			"inspect <pokemon-name",
			description:	"Displays information about the specified Pokemon, if it is in your Pokedex",
			callback:		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"Lists the entries in your Pokedex",
			callback:		commandPokedex,
		},
	}
}