package main

type cliCommand struct {
  name string
  description string
  callback func(arg string) error
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

  commands["showmap"] = cliCommand {
    name:       "showmap",
    description: "Show the map",
    callback:    showmap,
  }

  commands["showmapb"] = cliCommand {
    name: "showmapb",
    description: "Show the previous page of the map",
    callback: showmapb,
  }

  commands["explore"] = cliCommand {
    name: "explore",
    description: "Explore a specific area and return found pokemons",
    callback: explore,
  }

  commands["catch"] = cliCommand {
    name: "catch",
    description: "Catch a pokemon in a specific area",
    callback: catch,
  }

  commands["inspect"] = cliCommand {
    name: "inspect",
    description: "Inspect a speficic pokemon for details",
    callback: inspect,
  }

  commands["pokedex"] = cliCommand {
    name: "pokedex",
    description: "Show all your captured pokemons",
    callback: pokedex,
  }

  return commands
}
