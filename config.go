package main

import "github.com/franciscodev64/go-pokedex/internal/pokecache"

type config struct {
	next     *string
	previous *string
	cache    *pokecache.Cache
	Pokedex  map[string]Pokemon
}
