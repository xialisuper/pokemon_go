package main

import "fmt"

func commandHelp(config *config, args...string) error {
	// fmt.Println("Welcome to the Pokedex! \nUsage:")
	for _, command := range commands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
