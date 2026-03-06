package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(c *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if c.next != nil {
		url = *c.next
	}

	var data []byte
	if cached, ok := c.cache.Get(url); ok {
		fmt.Println("Using cache...")
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

		c.cache.Add(url, data)
	}

	var locations locationAreaResponse
	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	c.next = locations.Next
	c.previous = locations.Previous

	return nil
}

func commandMapb(c *config, args []string) error {
	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *c.previous
	var data []byte

	// Use cache if available
	if cached, ok := c.cache.Get(url); ok {
		fmt.Println("Using cache...")
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

		c.cache.Add(url, data)
	}

	// Parse JSON response
	var locations locationAreaResponse
	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}

	// Print all location area names
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	// Update config for next/previous pagination
	c.next = locations.Next
	c.previous = locations.Previous

	return nil
}
