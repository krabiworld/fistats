# Fiber Stats

Small middleware to find unused routes in app.

## Example

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/krabiworld/fistats"
    "github.com/krabiworld/fistats/fistorage"
)

func main() {
    var key string
	
    app := fiber.New()
    app.Use(fistats.New(&key, fistats.Config{Storage: fistorage.NewMemory()}))
}
```
