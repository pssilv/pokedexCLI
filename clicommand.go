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

  return commands
}
