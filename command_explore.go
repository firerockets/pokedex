package main

import "fmt"

func commandExplore(config *Config, param string) error {
	fmt.Printf("Exploring %s...\n", param)

	location, err := config.client.ExploreLocation(param)

	if err != nil {
		return err
	}

	if len(location.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
	} else {
		fmt.Println("No Pokemon found!")
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
