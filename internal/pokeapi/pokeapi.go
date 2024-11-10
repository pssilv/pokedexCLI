package pokeapi

// For locations requests
var locations_url = "https://pokeapi.co/api/v2/location-area"
// For area requests
const fixed_locations_url = "https://pokeapi.co/api/v2/location-area"
// for pokemon requests
const fixed_pokemon_url = "https://pokeapi.co/api/v2/pokemon/"

var at_first = true

func NextUrl() {
  if !at_first {
    locations_url = GetLocations().Next
  } else {
    at_first = false
  }
}

func PreviousUrl() {
  if GetLocations().Previous != "" {
    locations_url = GetLocations().Previous
  }
}
