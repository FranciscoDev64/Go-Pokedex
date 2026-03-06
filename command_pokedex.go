package main

import "fmt"

func commandPokedex(c *config, args []string) error {
	if len(c.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range c.Pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
