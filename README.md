# Pokedex CLI

A command-line Pokedex application built with Go that allows you to explore Pokemon locations, catch Pokemon, and inspect your collection. This project was created as part of the [Boot.dev](https://www.boot.dev) curriculum to practice Go programming concepts including HTTP clients, caching, and CLI development.

## Features

- **Interactive CLI** - Command-line interface with intuitive commands
- **Location Exploration** - Browse Pokemon locations and discover what Pokemon live there
- **Pokemon Catching** - Attempt to catch Pokemon with randomized success rates
- **Personal Pokedex** - View and inspect your caught Pokemon collection
- **Intelligent Caching** - HTTP response caching to improve performance and reduce API calls
- **Pagination Support** - Navigate through location lists with forward/backward commands

## Commands

| Command              | Description                       | Usage                      |
| -------------------- | --------------------------------- | -------------------------- |
| `help`               | Display available commands        | `help`                     |
| `map`                | Show next page of locations       | `map`                      |
| `mapb`               | Show previous page of locations   | `mapb`                     |
| `explore <location>` | Explore a location to see Pokemon | `explore pallet-town-area` |
| `catch <pokemon>`    | Attempt to catch a Pokemon        | `catch pikachu`            |
| `inspect <pokemon>`  | View details of a caught Pokemon  | `inspect pikachu`          |
| `pokedex`            | List all caught Pokemon           | `pokedex`                  |
| `exit`               | Exit the application              | `exit`                     |

## Sample Gameplay

```
Pokedex > help

Welcome to the Pokedex!
Usage:

help: Displays a help message
catch <pokemon_name>: catch a pokemon
inspect <pokemon_name: prints out pokemons details
pokedex: Lists all pokemon in Pokedex
explore <location_name>: Explore a location
map: Get the next page of locations
mapb: Get the previous page of locations
exit: Exit the Pokedex

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
...

Pokedex > explore pallet-town-area
Exploring pallet-town-area...
Found Pokemon:
 - rattata
 - pidgey

Pokedex > catch pikachu
pikachu was caught!
You may now inspect it with the inspect command.

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
- hp: 35
- attack: 55
- defense: 40
- special-attack: 50
- special-defense: 50
- speed: 90
Types:
- electric
```

## Prerequisites & Installation

### Required Software

- **Go 1.23.2+** - Download from [golang.org](https://golang.org)
- **Internet connection** - Required for Pokemon API access

### Installation Steps

1. **Clone the repository**

   ```bash
   git clone https://github.com/Ghostvox/pokedexcli.git
   cd pokedexcli
   ```

2. **Initialize Go modules** (if needed)

   ```bash
   go mod tidy
   ```

3. **Build the application**

   ```bash
   go build -o pokedexcli
   ```

4. **Run the application**

   ```bash
   ./pokedexcli
   ```

   Or run directly with Go:

   ```bash
   go run .
   ```

## Project Structure

```
pokedexcli/
├── main.go                          # Application entry point
├── repel.go                         # REPL (interactive CLI) implementation
├── repel_test.go                    # REPL tests
├── command_*.go                     # Individual command implementations
├── go.mod                           # Go module definition
├── .gitignore                       # Git ignore rules
├── internal/
│   ├── pokeAPI/                     # Pokemon API client package
│   │   ├── client.go                # HTTP client with caching
│   │   ├── location_get.go          # Location fetching
│   │   ├── location_list.go         # Location listing
│   │   ├── pokemon_get.go           # Pokemon fetching
│   │   ├── pokeApi.go               # API constants
│   │   ├── type_pokemon.go          # Pokemon data structures
│   │   └── types_locations.go       # Location data structures
│   └── pokecache/                   # Caching package
│       ├── pokecache.go             # Cache implementation
│       └── pokecache_test.go        # Cache tests
└── README.md                        # This file
```

## How It Works

### API Integration

The application integrates with the [PokeAPI](https://pokeapi.co/) to fetch real Pokemon data including:

- Location information and Pokemon encounters
- Detailed Pokemon stats, types, and abilities
- Paginated location lists

### Caching System

- **Intelligent HTTP caching** reduces API calls and improves performance
- **Automatic cache expiration** (5-minute default) ensures data freshness
- **Thread-safe implementation** using mutexes for concurrent access

### Catch Mechanics

- Pokemon catching uses a **randomized success system**
- **Base experience** affects catch difficulty (higher experience = harder to catch)
- Successfully caught Pokemon are stored in your personal Pokedex

### Command Architecture

- **Modular command system** with individual files for each command
- **Flexible argument handling** for commands that require parameters
- **Error handling** with user-friendly error messages

## Configuration

### Cache Settings

Modify cache behavior in `main.go`:

```go
pokeClient := pokeapi.NewClient(
    5*time.Second,    // HTTP timeout
    time.Minute*5,    // Cache expiration
)
```

### API Base URL

API endpoint configured in `internal/pokeAPI/pokeApi.go`:

```go
const baseURL = "https://pokeapi.co/api/v2"
```

## Testing

Run the included tests:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run specific package tests
go test ./internal/pokecache
```

## Development

### Adding New Commands

1. Create a new `command_<name>.go` file
2. Implement the command function with signature: `func(cfg *config, args ...string) error`
3. Add the command to the `getCommands()` map in `repel.go`

### Extending API Functionality

- Add new API endpoints in the `internal/pokeAPI` package
- Create corresponding data structures in the `types_*.go` files
- Implement caching for new endpoints following existing patterns

## Learning Objectives

This project demonstrates key Go programming concepts:

- **HTTP client development** with custom timeouts and error handling
- **JSON marshaling/unmarshaling** for API responses
- **Package organization** with internal packages and clean interfaces
- **Concurrent programming** with goroutines and mutexes for caching
- **CLI development** with interactive command processing
- **Testing** with table-driven tests and test coverage
- **Error handling** with Go's explicit error return pattern
- **Interface design** for modular and testable code

## Dependencies

- **Standard Library Only** - No external dependencies required
- Uses built-in packages: `net/http`, `encoding/json`, `time`, `sync`, etc.

## Troubleshooting

### Common Issues

**"Connection refused" or network errors**

- Check your internet connection
- Verify that `pokeapi.co` is accessible
- Firewall may be blocking HTTP requests

**"Pokemon not found" errors**

- Use exact Pokemon names (lowercase)
- Try common Pokemon like: `pikachu`, `charmander`, `bulbasaur`
- Use the `explore` command to find Pokemon in specific locations

**Commands not working**

- Type `help` to see all available commands
- Ensure proper command syntax (check spacing and spelling)
- Some commands require parameters (e.g., `catch pikachu`)

**Build errors**

- Ensure you have Go 1.23.2 or later installed
- Run `go mod tidy` to resolve any module issues
- Check that you're in the correct directory

## API Rate Limiting

The application implements intelligent caching to be respectful of the PokeAPI:

- **5-minute cache expiration** reduces redundant requests
- **HTTP timeouts** prevent hanging connections
- **Error handling** for rate limits and network issues

Perfect for Go developers looking to learn CLI development, HTTP clients, and working with external APIs in a fun, interactive way!
