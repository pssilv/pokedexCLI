package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)


func GetLocations() string {
  res, err := http.Get(locations_url)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  
  body, err := io.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  var locations string

  json.Unmarshal(body, &locations)
  
  return locations
} 
