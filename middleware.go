package stats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (s *Stats) Middleware(c *fiber.Ctx) error {
	path := utils.CopyString(c.Path())

	go func() {
		err := s.storage.Increment(path)
		if err != nil {
			panic(err)
		}
	}()

	return c.Next()
}
