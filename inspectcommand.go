package main

import "fmt"

func inspect(pokemon string) error {
  pokemon_data, ok := catched_pokemons[pokemon]
  if !ok {
    fmt.Println("you have not caught that pokemon, try - catch - to catch a pokemon")
    return nil
  }

  fmt.Println("------------------------")
  fmt.Printf("Name: %s\n", pokemon_data.Name)
  fmt.Printf("Height: %d\n", pokemon_data.Height)
  fmt.Printf("Weight: %d\n", pokemon_data.Weight)
  fmt.Println("Stats:")
  for _, pokemon_stats := range pokemon_data.Stats {
    fmt.Printf("  -%v: %d\n", pokemon_stats.Stat.Name, pokemon_stats.BaseStat)
  }
  fmt.Println("Types:")
  for _, pokemon_types := range pokemon_data.Types {
    fmt.Printf("  -%v\n", pokemon_types.Type.Name)
  }
  fmt.Println("------------------------")

  return nil
}
