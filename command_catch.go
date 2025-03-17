package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, param string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", param)

	pokemon, err := config.client.FetchPokemon(param)

	if err != nil {
		return err
	}

	chance := rand.Intn(pokemon.BaseExperience)

	if chance < 40 {
		fmt.Printf("%s was caught!\n", param)
		config.pokemons[param] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", param)
	}

	return nil
}
