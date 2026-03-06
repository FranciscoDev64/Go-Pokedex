package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokémon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	var data []byte

	// Use cache
	if cached, ok := c.cache.Get(url); ok {
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

	// Unmarshal API response
	var raw PokemonAPIResponse
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Map stats from API response to internal Stat type
	stats := make([]Stat, len(raw.Stats))
	for i, s := range raw.Stats {
		stats[i] = Stat{
			BaseStat: s.BaseStat,
			Name:     s.Stat.Name, // extract from wrapper
		}
	}

	// Build internal Pokemon
	pkm := Pokemon{
		Name:           raw.Name,
		BaseExperience: raw.BaseExperience,
		Height:         raw.Height,
		Weight:         raw.Weight,
		Stats:          stats,
		Types:          raw.Types,
	}

	// Initialize Pokedex if nil
	if c.Pokedex == nil {
		c.Pokedex = make(map[string]Pokemon)
	}

	// Determine catch chance
	rand.Seed(time.Now().UnixNano())
	catchChance := 1.0 - float64(pkm.BaseExperience)/500.0
	if catchChance < 0.05 {
		catchChance = 0.05
	}

	if rand.Float64() < catchChance {
		fmt.Printf("%s was caught!\n", pkm.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		c.Pokedex[pkm.Name] = pkm
	} else {
		fmt.Printf("%s escaped!\n", pkm.Name)
	}

	return nil
}
