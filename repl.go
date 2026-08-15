package main

import (
	"strings"
	"fmt"
	"os"
)

type config struct {
	register map[string]cliCommand
}

type cliCommand struct {
	name 		string
	description string
	callback 	func(*config) error
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
}

var configuration = config{
	register: commandRegister,
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

func commandExit(ptr_config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(ptr_config *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	register := ptr_config.register
	for command := range register {
		fmt.Printf("%s: %s\n", command, register[command].description)
	}
	return nil

}