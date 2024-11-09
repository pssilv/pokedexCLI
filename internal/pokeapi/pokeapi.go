package pokeapi

// For locations request
var locations_url = "https://pokeapi.co/api/v2/location-area"
// For area request
const fixed_locations_url = "https://pokeapi.co/api/v2/location-area"
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
