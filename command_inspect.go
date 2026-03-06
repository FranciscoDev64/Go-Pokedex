package main

import "fmt"

type locationAreaResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a Pokémon name")
	}

	pokemonName := args[0]

	// Check if user has caught this Pokémon
	pkm, ok := cfg.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// Print details
	fmt.Printf("Name: %s\n", pkm.Name)
	fmt.Printf("Height: %d\n", pkm.Height)
	fmt.Printf("Weight: %d\n", pkm.Weight)

	fmt.Println("Stats:")
	for _, stat := range pkm.Stats {
		fmt.Printf("  -%s: %d\n", stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pkm.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
