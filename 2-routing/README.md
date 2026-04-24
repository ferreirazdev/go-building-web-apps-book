# Chapter 2: Routing and URL Parameters

This chapter expands the server with structured routing and route-driven file rendering.

## Objective

Learn how to use a router to handle path patterns, parameters, and fallback behavior.

## Fundamentals Practiced

- integrating third-party packages with Go modules (`github.com/gorilla/mux`)
- creating a router with `mux.NewRouter()`
- defining routes with path params and regex (`/pages/{id:[0-9]+}`)
- reading URL params with `mux.Vars`
- checking file existence with `os.Stat`
- serving files and returning a fallback page when content is missing

## What This Chapter Demonstrates

- dynamic page mapping from URL IDs to files in `files/`
- route registration for numeric pages and named endpoints (`/homepage`, `/contact`)
- a simple 404 fallback using `files/404.html`

## How To Run

```bash
cd 2-routing
go mod tidy
go run maing.go
```

Open:
- `http://localhost:8080/pages/1`
- `http://localhost:8080/pages/2`
- `http://localhost:8080/pages/999`
- `http://localhost:8080/homepage`
- `http://localhost:8080/contact`

## Book Navigation

- Previous: [Chapter 1 - Hello, Handlers, and Responses](../1-hellow/README.md)
- Next: [Chapter 3 - Connecting Data](../3-connecting-data/README.md)
- Back to index: [Main README](../README.md)
