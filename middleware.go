package stats

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func (s *Stats) Middleware(c *fiber.Ctx) error {
	go func() {
		err := s.storage.Increment(utils.CopyString(c.Path()))
		if err != nil {
			panic(err)
		}
	}()

	return c.Next()
}
