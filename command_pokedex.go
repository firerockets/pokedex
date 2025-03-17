package main

import "fmt"

func commmandPokedex(config *Config, param string) error {
	fmt.Println("Your Pokedex:")

	for pokemon, _ := range config.pokemons {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}
