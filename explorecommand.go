package main

import (
	"fmt"

	"github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func explore(area string) error {

  area_data, err := pokeapi.GetArea(area)

  if err != nil {
    fmt.Println(err)
    fmt.Println("You can try - showmap - for some valid areas")
    return nil
  }
 
  fmt.Println("Found Pokemon:")
  fmt.Println("------------------------")
  for _, pokemon := range area_data.PokemonEncounters {
    fmt.Println(pokemon.Pokemon.Name)
  }
  fmt.Println("------------------------")

  return nil
}


