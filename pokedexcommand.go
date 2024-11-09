package main

import "fmt"

func pokedex(none string) error {
  if len(catched_pokemons) == 0 {
    fmt.Println("You doesn't have pokemons try - catch - to catch a pokemon")
    return nil
  }

  fmt.Println("Your pokedex:")
  for pokemon := range catched_pokemons {
    fmt.Printf(" - %v\n", pokemon)
  }

  return nil
}
