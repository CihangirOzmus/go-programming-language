package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors,omitempty"`
}

var movies = []Movie{
	{"Casablanca", 1984, true, []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{"Cool Hand Luke", 1967, true, []string{"Paul Newman"}},
	{"Another One No Color", 1967, false, nil},
}

func main() {
	fmt.Printf("Movies in Array: %v\n", movies)

	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %v", err)
	}
	fmt.Printf("Movies in JSON: %s\n", data)
}
