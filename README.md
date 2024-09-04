# Pokerhands

## Development

```
make dev
```

This starts the application through docker compose, and runs the Go application with Air for hot reloading.

```
make test
```

This runs all the tests within the `api` folder.

## Details

- API written in Go
- Mostly standard libraries, even for routing.
- Workflows present for testing, linting and building.
- Ko used for building, as it has a considerable smaller footprint.
