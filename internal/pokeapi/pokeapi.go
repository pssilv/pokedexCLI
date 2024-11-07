package pokeapi

var locations_url = "https://pokeapi.co/api/v2/location-area"
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
