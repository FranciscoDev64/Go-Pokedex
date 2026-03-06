// pokemon_structs.go
package main

type StatWrapper struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type TypeWrapper struct {
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

type PokemonAPIResponse struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []StatWrapper `json:"stats"`
	Types          []TypeWrapper `json:"types"`
}

// Internal Pokemon struct stored in Pokedex
type Stat struct {
	BaseStat int
	Name     string
}

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          []Stat
	Types          []TypeWrapper
}
