package main

import (
  "fmt"
)

func help(none string) error {
  fmt.Println("\nWelcome to the pokedex!")
  fmt.Println("Usage:")
  fmt.Println()

  commands := commands()

  for _, cmd := range commands {
    fmt.Printf("%s: %s\n", cmd.name, cmd.description)
  }
  fmt.Println()
  
  return nil
}
