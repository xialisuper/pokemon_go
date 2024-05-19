package main

import (
	"pokemoncli/api"
	"time"
)

func main() {

	// create a client for pokemon api
	client := api.NewClient(5*time.Second, 5*time.Minute)

	config := &config{
		PokemonClient: &client,
		PokemonDex:    make(map[string]api.PokemonDetail),
	}

	// start the repl
	StartRepl(config)
}
