package main

import (
	"errors"
	"fmt"
)

func commandPokedex(config *config, args ...string) error {

	dex := config.PokemonDex

	if len(dex) == 0 {
		return errors.New("no pokemon found in the dex")
	}

	fmt.Printf("Your Pokemon Dex:\n")

	for _, pokemon := range dex {
		fmt.Printf(" -%s\n", pokemon.Name)
	}

	return nil
}
