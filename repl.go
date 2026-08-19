package main

import (
	"strings"
	"fmt"
	"os"
	"github.com/CrstKng/pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	register 		map[string]cliCommand
	next 	 		string
	previous 		string
	pokeapiClient   *pokeapi.Client
}

type cliCommand struct {
	name 		string
	description string
	callback 	func(*config, []string) error
}

var commandRegister = map[string]cliCommand{
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"map": {
		name: "map",
		description: "Display the names of next 20 locations",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Display the names of last 20 locations",
		callback: commandMapb,
	},
	"explore": {
		name: "explore",
		description: "Explore the pokemons in a certain location",
		callback: commandExplore,
	},
}

var configuration = config{
	register: commandRegister,
	next: "https://pokeapi.co/api/v2/location-area/",
	previous: "",
	pokeapiClient: pokeapi.NewClient(20 * time.Second, time.Minute),

}

func cleanInput(text string) []string {
	var cleanStrings []string
	separated_strings := strings.Fields(text)
	for _, s := range separated_strings {
		lowerCase := strings.ToLower(s)
		cleanStrings = append(cleanStrings, lowerCase)
	}
	return cleanStrings
}

func commandExit(ptr_config *config, location_names []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(ptr_config *config, location_names []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	register := ptr_config.register
	for command := range register {
		fmt.Printf("%s: %s\n", command, register[command].description)
	}
	return nil
}

func showPage(ptr_config *config, url string) error {
	apiResource, err := ptr_config.pokeapiClient.ListLocationAreas(url)
	if err != nil {
		return fmt.Errorf("error when getting location list: %s", err)
	}
	ptr_config.next = apiResource.Next
	ptr_config.previous = apiResource.Previous
	for _, result := range apiResource.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMap(ptr_config *config, location_names []string) error {
	return showPage(ptr_config, ptr_config.next)
}

func commandMapb(ptr_config *config, location_names []string) error {
	if ptr_config.previous == "" {
		return fmt.Errorf("you're on the first page")
	}
	return showPage(ptr_config, ptr_config.previous)
}

func commandExplore(ptr_config *config, location_names []string) error {
	if len(location_names) == 0 {
		return fmt.Errorf("you must provide a location area name")
	}
	url := "https://pokeapi.co/api/v2/location-area/" + location_names[0] + "/"
	locationArea, err := ptr_config.pokeapiClient.LocationArea(url)
	if err != nil {
		return fmt.Errorf("error when getting location area info: %s", err)
	}
	fmt.Printf("Exploring %s...\n", location_names[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon_encounter := range locationArea.Pokemon_encounters {
		fmt.Println(pokemon_encounter.Pokemon.Name)
	}
	return nil
}