package main

import (
  "fmt"
  

  "github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func showmap() error {
  locations_data := pokeapi.GetLocations()

  fmt.Println("------------------------")
  for _, location := range locations_data.Results {
    fmt.Println(location.Name)
  }
  fmt.Println("------------------------")
  pokeapi.NextUrl()

  return nil
}
