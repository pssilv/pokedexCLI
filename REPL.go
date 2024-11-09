package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func REPL() {
  scanner := bufio.NewScanner(os.Stdin)
  commands := commands()

  fmt.Print("pokedex > ")
  for scanner.Scan() {
    separated_input := strings.Fields(scanner.Text())
    input := strings.ToLower(separated_input[0])

    command := commands[input]
    switch (command.name) {
      case "showmap", "showmapb", "help", "exit":
        command.callback("no arguments required")
      case "explore":
        if len(separated_input) > 1 && len(separated_input) < 3 {
          area := separated_input[1]
          command.callback(area)
          // area_url is from "catchcommand.go"
          area_url = area
        } else if len(separated_input) > 2 {
            fmt.Println("explore got 2 arguments, need 1")
          break
        } else {
          fmt.Println("Needs a area name first, try - showmap - for some areas")
          break
        }
      case "catch":
        if len(separated_input) > 1 && len(separated_input) < 3 {
          pokemon := separated_input[1]
          command.callback(pokemon)
        } else if len(separated_input) > 2 {
            fmt.Println("catch got 2 arguments, need 1")
        } else {
          fmt.Println("Needs a pokemon first, try - explore - for some pokemons")
        }
      default:
        fmt.Println("Invalid command, try - help -")

    }
    fmt.Print("pokedex > ")
  }

    fmt.Print("pokedex > ")
}

