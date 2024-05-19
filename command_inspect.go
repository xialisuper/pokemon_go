package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, args ...string) error {

	if len(args) < 1 {
		return errors.New("you need to specify a pokemon name to inspect")

	}

	pokemonName := args[0]

	fmt.Printf("Inspecting pokemon %s\n", pokemonName)

	pokemonDex, ok := config.PokemonDex[pokemonName]

	if !ok {
		return errors.New("you have not caught this pokemon yet")
	}
	var hp int
	var attack int
	var defense int
	var spAtk int
	var spDef int
	var speed int

	for _, stat := range pokemonDex.Stats {

		if stat.Stat.Name == "hp" {
			hp = stat.BaseStat
		}

		if stat.Stat.Name == "attack" {
			attack = stat.BaseStat
		}

		if stat.Stat.Name == "defense" {
			defense = stat.BaseStat
		}

		if stat.Stat.Name == "special-attack" {
			spAtk = stat.BaseStat
		}

		if stat.Stat.Name == "special-defense" {
			spDef = stat.BaseStat
		}

		if stat.Stat.Name == "speed" {
			speed = stat.BaseStat
		}

	}

	types := []string{}
	for _, t := range pokemonDex.Types {
		types = append(types, t.Type.Name)
	}

	fmt.Printf("Name: %s\nWeight: %d\nHeight: %d\nStats:\n -hp: %d\n -attack: %d\n -defense: %d\n -sp.attack: %d\n -sp.defense: %d\n -speed: %d\n", pokemonName, pokemonDex.Weight, pokemonDex.Height, hp, attack, defense, spAtk, spDef, speed)

	fmt.Println("Types:")
	for _, t := range types {
		fmt.Printf(" -%s\n", t)
	}

	return nil
}
