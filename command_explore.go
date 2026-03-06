package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("please provide a location area")
	}

	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)

	url := "https://pokeapi.co/api/v2/location-area/" + areaName

	var data []byte

	// CACHE CHECK
	if cached, ok := cfg.cache.Get(url); ok {
		data = cached
	} else {

		res, err := http.Get(url)
		if err != nil {
			return err
		}

		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, data)
	}

	// STRUCT FOR RESPONSE
	var location locationAreaResponse

	err := json.Unmarshal(data, &location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	return nil
}
