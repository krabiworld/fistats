package fistats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/uuid"
	"sort"
)

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Usage  uint64 `json:"usage"`
}

func New(key *string, config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	id := "/" + uuid.NewString()
	*key = id

	return func(c *fiber.Ctx) error {
		path := utils.CopyString(c.Path())

		if c.Method() == fiber.MethodGet && path == id {
			appRoutes := c.App().GetRoutes(true)

			mapRoutes := make(map[string]fiber.Route)
			for _, route := range appRoutes {
				if route.Method == "HEAD" {
					continue
				}

				mapRoutes[route.Path] = route
			}

			storedRoutes, err := cfg.Storage.GetAll()
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, err.Error())
			}

			routes := make([]Route, 0)
			for key, route := range mapRoutes {
				if key == id {
					continue
				}

				routes = append(routes, Route{
					Method: route.Method,
					Path:   route.Path,
					Usage:  storedRoutes[key],
				})
			}

			sort.Slice(routes, func(i, j int) bool {
				return routes[i].Usage > routes[j].Usage
			})

			return c.JSON(routes)
		}

		go func() {
			err := cfg.Storage.Increment(path)
			if err != nil {
				panic(err)
			}
		}()

		return c.Next()
	}
}
