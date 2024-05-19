package main

import "fmt"

func commandMapf(config *config, args...string) error {

	res, err := config.PokemonClient.GetLocationAreas(config.Next)
	if err != nil {
		return err
	}

	config.Next = res.Next
	config.Prev = res.Previous

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapb(config *config, args...string) error {

	if config.Prev == nil {
		fmt.Println("No previous maps found.")
		return nil
	}

	res, err := config.PokemonClient.GetLocationAreas(config.Prev)
	if err != nil {
		return err
	}

	config.Next = res.Next
	config.Prev = res.Previous

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
	return nil
}
