package main

import (
	"fmt"
	"log"
	"math/rand/v2"

	//"math/rand"

	"github.com/pssilv/pokedexCLI/internal/pokeapi"
)

func catch(pokemon string) error {
  pokemon_data, err := pokeapi.GetPokemonData(pokemon)
  if err != nil {
    fmt.Println(err)
    fmt.Println("You can try - explore - to get some pokemon names on a specific area")
    return nil
  }

  pokemon_area_data, err := pokeapi.GetArea(area_url)
  if err != nil {
    return err
  }

  pokemon_area_list := make(map[string]pokeapi.Pokemon)

  for _, pokemon := range pokemon_area_data.PokemonEncounters {
    pokemon_area_list[pokemon.Pokemon.Name], err = pokeapi.GetPokemonData(pokemon.Pokemon.Name)
    if err != nil {
      log.Fatalf("Error: %v", err)
    }
  }

  res := rand.IntN(pokemon_data.BaseExperience)

  fmt.Printf("Throwing a pokeball at %v\n", pokemon_data.Name)
  if res > 40 {
    fmt.Printf("%v has escaped!\n", pokemon_data.Name)
    return nil
  }
  fmt.Printf("%v was caugh!\n", pokemon_data.Name)


  catched_pokemons[pokemon_data.Name] = pokemon_data
  return nil
}


