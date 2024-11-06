package main

import (
  "fmt"
  
  "github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func showMap() error {
  json_data := pokeapi.GetLocations()
  fmt.Println(json_data)

  return nil
}
