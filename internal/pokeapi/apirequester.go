package pokeapi

import (
	"encoding/json"
	"io"
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

var locations_cache = pokecache.NewCache(30 * time.Second)

func GetLocations() Locations {
  var locations Locations

  if data, exists := locations_cache.Get(locations_url); exists {
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

  body, err := io.ReadAll(res.Body)
 
  if err := json.Unmarshal(body, &locations); err != nil {
    log.Fatal(err)
  }

  locations_cache.Add(locations_url, body)  
 
  return locations
}
