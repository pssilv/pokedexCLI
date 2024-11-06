package main

type cliCommand struct {
  name string
  description string
  callback func () error
}

func commands() map[string]cliCommand {
  commands := make(map[string]cliCommand)

  commands["help"] = cliCommand {
    name:        "help", 
    description: "Displays a help message", 
    callback:     help,
  }

  commands["exit"] = cliCommand {
    name:        "exit",
    description: "Exit the Pokedex",
    callback:    exit,
  }

  commands["showMap"] = cliCommand {
    name:       "showMap",
    description: "Show the map",
    callback:    showMap,
  }

  return commands
}
