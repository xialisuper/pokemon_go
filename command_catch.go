package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// commandCatch is the function to catch a pokemon.
func commandCatch(config *config, args ...string) error {

	if len(args) == 0 {

		return errors.New("please input the pokemon name to catch")

	}

	fmt.Println("Throwing a Pokeball at", args[0], "...")

	pokemonName := args[0]

	pokemon, err := config.PokemonClient.CatchPokemon(pokemonName)

	if err != nil {
		return err
	}

	if isCaught(pokemon.BaseExperience) {
		fmt.Println(pokemon.Name, "is caught!")

		// add the pokemon to the user's inventory
		config.PokemonDex[pokemon.Name] = pokemon

	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}

	return nil
}

// a function to determine if a pokemon is caught or not based on pokemon details base_experience value.
func isCaught(base_experience int) bool {

	// create a random number between 0 and 100
	randNum := rand.Intn(200)

	fmt.Println("Random number:", randNum, "base_experience:", base_experience)
	// if the pokemon's base_experience is greater than or equal to the random number, the pokemon is caught.
	return base_experience <= randNum

}
