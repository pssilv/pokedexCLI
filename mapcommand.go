package main

import (
  "fmt"
  

  "github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func showmap() error {
  pokeapi.NextUrl()
  locations_data := pokeapi.GetLocations()

  fmt.Println("------------------------")
  for _, location := range locations_data.Results {
    fmt.Println(location.Name)
  }
  fmt.Println("------------------------")

  return nil
}
