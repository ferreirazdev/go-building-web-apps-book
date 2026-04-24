# Chapter 1: Hello, Handlers, and Responses

This chapter introduces a first Go web server with basic handlers and response patterns.

## Objective

Understand how to start an HTTP server and map URLs to handler functions.

## Fundamentals Practiced

- `package main` entry point and `main()` execution flow
- constants for server configuration (`Port`)
- function-based handlers with `http.HandleFunc`
- static file serving with `http.ServeFile`
- dynamic response writing with `fmt.Fprintln`
- basic use of the `time` package for runtime data

## What This Chapter Demonstrates

- A running server on port `8080`
- `/static` returns HTML from `static.html`
- `/` returns a dynamic string with current time
- a disabled example function (`serveError`) for future error handling practice

## How To Run

```bash
cd 1-hellow
go mod tidy
go run main.go
```

Open:
- `http://localhost:8080/`
- `http://localhost:8080/static`

## Book Navigation

- Next: [Chapter 2 - Routing and URL Parameters](../2-routing/README.md)
- Back to index: [Main README](../README.md)
