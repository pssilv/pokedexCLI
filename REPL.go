package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func REPL() {
  scanner := bufio.NewScanner(os.Stdin)
  commands := commands()

  fmt.Print("pokedex > ")
  for scanner.Scan() {
    input := strings.ToLower(scanner.Text())

    if command, ok := commands[input]; ok {
      command.callback()
    } else {
      fmt.Println("Invalid command, try: - help -")
      fmt.Println()
    }
    fmt.Print("pokedex > ")
  }

    fmt.Print("pokedex > ")
}
