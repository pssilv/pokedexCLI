package main

import (
  "fmt"

  "github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func showmapb() error {
  if pokeapi.GetLocations().Previous == "" {
    fmt.Println("Already at the first page")
    return nil
  }

  pokeapi.PreviousUrl()
  locations_data := pokeapi.GetLocations()

  fmt.Println("------------------------")
  for _, location := range locations_data.Results {
    fmt.Println(location.Name)
  }
  fmt.Println("------------------------")

  return nil
}
