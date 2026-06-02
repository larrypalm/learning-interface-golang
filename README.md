# learn-interfaces-go

Learning Go interfaces by building a small HTTP server backed by PostgreSQL.

## What I'm learning

### Done
- **Implementing a package interface** — `store.Store` satisfies the `goauth.UserStore` interface from my own [`go-auth`](https://github.com/larrypalm/go-auth) package by implementing methods like `CreateUser`. Learned the difference between pointer receivers (`*Store`) and value receivers (`Store`) and why only `*Store` satisfies the interface when methods use pointer receivers.

### In progress
- **HTTP request interfaces** — using Go's `http.Handler` and `http.ResponseWriter` / `*http.Request` to handle incoming requests through a standard interface.

## Stack

- Go 1.26
- PostgreSQL via `pgx/v5`
- [`go-auth`](https://github.com/larrypalm/go-auth) — my own auth package (local replace directive)

## Structure

```
cmd/server/      entry point
internal/
  handler/       HTTP routing and handlers
  store/         PostgreSQL store, implements go-auth interfaces
```

## Run

```bash
go run ./cmd/server
```

This is a learning project. The code reflects my understanding of Go at the time of writing — expect rough edges and iterative improvements as I go.