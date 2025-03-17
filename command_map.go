package main

import (
	"fmt"
)

func commandMapf(config *Config, param string) error {
	return commandMap(config.next, config)
}

func commandMapb(config *Config, param string) error {
	return commandMap(config.previous, config)
}

func commandMap(url string, config *Config) error {
	resource, err := config.client.ListLocation(url)

	if err != nil {
		return err
	}

	config.next = resource.Next
	config.previous = resource.Previous

	for _, result := range resource.Results {
		fmt.Println(result.Name)
	}

	return nil
}
