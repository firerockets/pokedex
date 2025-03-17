package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/firerockets/pokedexcli/internal/pokeapi"
)

func startRepl() {
	initializeCommands()
	scanner := bufio.NewScanner(os.Stdin)

	config := Config{
		client:   pokeapi.NewClient(),
		pokemons: make(map[string]pokeapi.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]

		var param string

		if len(words) > 1 {
			param = words[1]
		}

		cmd, exists := commands[command]
		if exists {
			err := cmd.callback(&config, param)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	split := strings.Fields(lower)
	return split
}
