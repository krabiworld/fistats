package stats

import (
	"github.com/gofiber/fiber/v2"
	"sort"
	"strings"
)

func (s *Stats) Route(c *fiber.Ctx) error {
	key := c.Query("key")
	if strings.TrimSpace(key) != s.Key() {
		return c.Status(fiber.StatusNotFound).SendString("404 Not Found")
	}

	appRoutes := c.App().GetRoutes(true)

	mapRoutes := make(map[string]fiber.Route)
	for _, route := range appRoutes {
		if route.Method == "HEAD" {
			continue
		}

		mapRoutes[route.Path] = route
	}

	storedRoutes, err := s.storage.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	routes := make([]Route, 0)
	unused := make([]Route, 0)
	for key, route := range mapRoutes {
		stored, ok := storedRoutes[key]
		if !ok {
			unused = append(unused, Route{Path: route.Path})
			continue
		}

		routes = append(routes, Route{
			Path:  route.Path,
			Usage: stored,
		})
	}

	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Usage > routes[j].Usage
	})

	return c.JSON(Response{
		Routes: routes,
		Unused: unused,
	})
}
