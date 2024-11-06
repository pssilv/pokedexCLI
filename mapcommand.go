package main

import (
  "fmt"
  

  "github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func showMap() error {
  areas_data := pokeapi.GetLocations()
  fmt.Println(areas_data)

  return nil
}
