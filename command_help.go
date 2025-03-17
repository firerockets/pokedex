package main

import (
	"fmt"
)

func commandHelp(config *Config, param string) error {
	str := "Welcome to the Pokedex!\nUsage:\n\n"
	for _, command := range commands {
		str += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Print(str)
	return nil
}
