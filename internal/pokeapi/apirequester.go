package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

  "github.com/pssilv/pokedexCLI/internal/pokecache"
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

var cache = pokecache.NewCache(10)

func GetLocations() Locations {
  var locations Locations

  if data, exists := cache.Get(locations_url); exists {
    json.Unmarshal(data, &locations)
    return locations
  }

  client := &http.Client{
    Timeout: 5 * time.Second,
  }
  
  req, err := http.NewRequest("GET", locations_url, nil)
  if err != nil {
    log.Fatal(err)
  }

  res, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()

  decoder := json.NewDecoder(res.Body)
  
  if err := decoder.Decode(&locations); err != nil {
    log.Fatal(err)
  }

  byte_slice, err := json.Marshal(res.Body)
  if err != nil {
    log.Fatal(err)
  }  

  cache.Add(locations_url, byte_slice)  

 
  return locations
}
