package main

import (
  "fmt"

  "github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func showmap(none string) error {
  pokeapi.NextUrl()
  map_index++
  locations_data := pokeapi.GetLocations()

  fmt.Println("------------------------")
  fmt.Printf("Map index: %d\n", map_index)
  fmt.Println("------------------------")
  for _, location := range locations_data.Results {
    fmt.Println(location.Name)
  }
  fmt.Println("------------------------")

  return nil
}
