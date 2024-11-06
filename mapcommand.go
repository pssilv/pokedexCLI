package main

import (
  "fmt"
  
  "github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func print_locations() error {
  json_data := pokeapi.getLocations()
  fmt.Println(json_data)

  return nil
}
