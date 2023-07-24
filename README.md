# todo-go

[![GoDoc](https://godoc.org/github.com/lib/pq?status.svg)](https://pkg.go.dev/github.com/lib/pq?tab=doc)

It's a simple TODO API on Golang.

## Routes

### Basic TODO

Сlassic to-do list.

- GET: /basic
- POST: /basic
- UPDATE: /basic/{id}
- DELETE: /basic/{id}

### Stack TODO

Implementation of the LIFO principle for a list of tasks.

- GET: /stack
- POST: /stack
- DELETE: /stack