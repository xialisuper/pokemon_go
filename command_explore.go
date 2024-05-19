package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) == 0 {

		return errors.New("please input the area name to explore pokemon")

	}
	areaName := args[0]
	fmt.Printf("Exploring area %s...\n", areaName)

	res, err := config.PokemonClient.ExploreLocationAreaByName(areaName)
	if err != nil {
		return err
	}

	pokemons := res.PokemonEncounters
	fmt.Printf("Found %d pokemon in area %s:\n", len(pokemons), areaName)

	for _, pokemon := range pokemons {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
