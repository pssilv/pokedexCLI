package pokeapi

var locations_url = "https://pokeapi.co/api/v2/location-area"
var do_previous = false

func NextUrl() {
    locations_url = GetLocations().Next
    do_previous = true
}

func PreviousUrl() {
  if do_previous == true && GetLocations().Previous != "" {
    locations_url = GetLocations().Previous
    do_previous = false
  }


  if GetLocations().Previous != "" {
    locations_url = GetLocations().Previous
  }
}
