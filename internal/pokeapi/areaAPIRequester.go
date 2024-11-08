package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
  "errors"

	"github.com/pssilv/pokedexCLI/internal/pokecache"
)


type Area struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var areas_cache = pokecache.NewCache(30 * time.Second)

func GetArea(areaName string) (Area, error){
  var area Area
  url := fmt.Sprintf("%v/%v", fixed_locations_url, areaName)
  
  if value, exists := areas_cache.Get(url); exists {
    json.Unmarshal(value, &area)
    return area, nil
  }

  client := &http.Client{
    Timeout: 5 * time.Second,
  }

  req, err := http.NewRequest("GET", url, nil)
    if err != nil {
      log.Fatal(err)
    }

  res, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    return Area{}, errors.New("Area doesn't exist try a valid area")
  }

  body, err := io.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  if err := json.Unmarshal(body, &area); err != nil {
    log.Fatal(err)
  }

  areas_cache.Add(url, body)

  return area, nil
}
