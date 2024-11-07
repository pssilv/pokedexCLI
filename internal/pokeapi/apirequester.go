package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations() Locations {
  client := &http.Client{}

  req, err := http.NewRequest("GET", locations_url, nil)
  if err != nil {
    log.Fatal(err)
  }

  res, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  
  var locations Locations
  decoder := json.NewDecoder(res.Body)
  

  if err := decoder.Decode(&locations); err != nil {
    log.Fatal(err)
  }
  
  return locations
}
