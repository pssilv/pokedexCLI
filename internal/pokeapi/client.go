package pokeapi

import (
  "net/http"
)

func CreateClient() *http.Client {
  client := &http.Client{}

  return client
}
