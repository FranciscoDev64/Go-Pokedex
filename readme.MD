# Go Pokedex CLI

A command-line Pokémon database built in **Go** that interacts with the [PokeAPI](https://pokeapi.co/).  
This project demonstrates working with REST APIs, JSON parsing, caching, and building a feature-rich CLI tool with Go.

---

## Features

- Interact with the Pokémon world using a terminal interface.
- Paginated exploration of location areas and their Pokémon.
- Catch Pokémon with chance-based mechanics.
- Inspect caught Pokémon with detailed stats.
- Persistent in-memory Pokedex with fast lookup and caching.
- Fully Go-native: HTTP requests, JSON decoding, concurrency-safe caching.

---

## Commands

| Command | Description |
|---------|-------------|
| `help` | Displays a help message |
| `exit` | Exit the Pokedex |
| `map` | Displays 20 location names, then the next 20 locations when used again |
| `mapb` | Displays the previous 20 location names |
| `explore <location>` | Shows all Pokémon in a given location area |
| `catch <pokemon>` | Attempt to catch a Pokémon by name. Catch rate depends on base experience |
| `inspect <pokemon>` | Inspect a caught Pokémon's details (stats, height, weight, types) |
| `pokedex` | List all Pokémon you have caught |

---

## Implementation Highlights

- **Go CLI Application**: Built with `bufio.Scanner` and modular command callbacks.
- **REST API Consumption**: Fetches data dynamically from PokeAPI endpoints for Pokémon and locations.
- **JSON Parsing**: Decodes complex nested JSON responses into Go structs.
- **Caching Layer**: Concurrent-safe caching to avoid repeated HTTP requests, using `sync.Mutex` and a timed reaper.
- **Randomized Mechanics**: Catch success based on Pokémon's base experience using `math/rand`.
- **Stateful Pokedex**: Keeps track of all caught Pokémon in memory, enabling `inspect` and `pokedex` commands.

---

## Example Usage

```bash
$ go run main.go
Pokedex > help
Pokedex > map
canalave-city-area
eterna-city-area
...
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokémon:
 - tentacool
 - tentacruel
 - magikarp
...
Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  - hp: 20
  - attack: 10
  - defense: 55
  - special-attack: 15
  - special-defense: 20
  - speed: 80
Types:
  - water
Pokedex > pokedex
Your Pokedex:
 - magikarp
 ```

## Tech Stack

Language: Go

APIs: PokeAPI

Concurrency: sync.Mutex for safe cache access

Data Handling: JSON unmarshalling, dynamic structs

CLI Patterns: Modular commands with callbacks, infinite REPL loop

## How to Run

Clone the repository:

git clone https://github.com/franciscodev64/go-pokedex.git

cd go-pokedex

Build and run:

go run main.go

Start exploring Pokémon!