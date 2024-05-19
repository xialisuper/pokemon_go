package main

import (
	"bufio"
	"fmt"
	"os"
	"pokemoncli/api"
	"strings"
)

type config struct {
	Next          *string
	Prev          *string
	PokemonClient *api.Client
	PokemonDex    map[string]api.PokemonDetail
}


type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args ...string) error
}

func StartRepl(config *config) {

	fmt.Println("Welcome to the Pokedex! ")
	input := bufio.NewScanner(os.Stdin) // 初始化 bufio.Scanner 对象
	for {
		// wait for user input
		fmt.Print("Pokedex > ")

		if !input.Scan() {
			fmt.Println("Error reading input:", input.Err())
			continue
		}

		// clean input
		words := cleanInput(input.Text())

		// check if input is empty
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		// check if command exists
		command, exists := commands()[commandName]
		if exists {
			// execute command
			err := command.callback(config, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command not found. Type 'help' for a list of commands.")
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func commands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"mapf": {
			name:        "mapf",
			description: "Displays next 20 maps of the Pokedex",
			callback:    commandMapf,
		},
		"mapb": {

			name:        "mapb",
			description: "Displays previous 20 maps of the Pokedex",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific area of the Pokedex",
			callback:    commandExplore,
		},

		"catch": {
			name:        "catch",
			description: "Catch a specific pokemon",
			callback:    commandCatch,
		},
		"inspect":{
			name:        "inspect",
			description: "Inspect a specific pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the Pokedex",
			callback:    commandPokedex,
		},
	}
}
