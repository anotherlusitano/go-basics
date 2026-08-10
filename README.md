# Basics of Go

A small repository I use to learn and practice Go.

The projects here follow along with the [Basics of Go](https://master.dev/courses/go-basics/) course from [Master.dev](https://master.dev/). Each directory covers a different topic of the course, with small programs and exercises to experiment with the concepts as I learn them.

## Go version

All the projects uses **Go 1.26**.

If you don't have the right version of Go installed, you can use [mise](https://mise.jdx.dev/) to install the required version:

```bash
mise install
```

The version is defined in `mise.toml`.

## Projects

### `the-basics`

The introductory Go examples, including:

- `first.go` — first Go program
- `calculator/` — a simple calculator
- `files/` — working with files
- `methods/` — methods and related concepts
- `modules/` — Go modules

Run the examples from the repository root:

```bash
cd the-basics
go run first.go
```

Or run one of the subprojects:

```bash
cd the-basics/calculator
go run .
```

```bash
cd the-basics/files
go run .
```

```bash
cd the-basics/methods
go run .
```

```bash
cd the-basics/modules
go run .
```

### `goroutines`

Examples focused on goroutines and concurrency.

```bash
cd goroutines
go run .
```

### `web-server`

A small HTTP server built while learning Go's web/server capabilities.

```bash
cd web-server
go run .
```

The server listens on `:3333`.

### `api-client`

A small API client that fetches cryptocurrency exchange rates from CEX.IO.

```bash
cd api-client
go run .
```

## Why this repository exists

This is primarily a learning repository.

The goal is to experiment with Go, understand how the language works, and keep the examples around for reference.
