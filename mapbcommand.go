package main

import (
  "fmt"

  "github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func showmapb(none string) error {
  if pokeapi.GetLocations().Previous == "" {
    fmt.Println("Already at the first page")
    return nil
  }

  pokeapi.PreviousUrl()
  map_index--
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
