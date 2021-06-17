package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Name string  `json:"name,omitempty"`
		Lat  float64 `json:"latitude,omitempty"`
		Long float64 `json:"longitude,omitempty"`
	}
	locations := []location{
		{Name: "name", Lat: -4.58, Long: 137.4},
		{Name: "bbb", Lat: -14.58, Long: 175.4},
		{Name: "ccc", Lat: -122.58, Long: 128.4},
	}
	jsonLocation, err := json.MarshalIndent(locations, "", "\t")
	exitOnError2(err)
	fmt.Println(string(jsonLocation))
}
func exitOnError2(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
