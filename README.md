# Fiber Stats

Small middleware to find unused routes in app.

## Example

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/krabiworld/stats"
    "github.com/krabiworld/stats/storage"
)

func main() {
    s := stats.New(storage.NewMemory())

    app := fiber.New()
    app.Use(s.Middleware)

    app.Get("/stats", s.Route)
	
    s.Key() // Key to access to /stats route
}
```
