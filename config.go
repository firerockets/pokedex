package main

import (
	"github.com/firerockets/pokedexcli/internal/pokeapi"
)

type Config struct {
	next     string
	previous string
	client   pokeapi.Client
	pokemons map[string]pokeapi.Pokemon
}
