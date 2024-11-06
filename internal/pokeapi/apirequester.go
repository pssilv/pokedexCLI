package pokeapi

import (
  "fmt"
	"io"
	"log"
	"net/http"
)


func getLocations() {
  res, err := http.Get(locations_url)
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  
  body, err := io.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(body)
} 
