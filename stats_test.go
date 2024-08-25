package stats

import (
	"encoding/json"
	"github.com/alicebob/miniredis/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/krabiworld/stats/storage"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func test(t *testing.T, stats *Stats) {
	log.Infof("Key: %s", stats.Key())

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(stats.Middleware)

	app.Get("/stats", stats.Route)
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	testReq := httptest.NewRequest("GET", "/test", nil)
	for i := 0; i < 10; i++ {
		_, err := app.Test(testReq)
		assert.Nil(t, err)
	}

	time.Sleep(2 * time.Second) // wait for all goroutines to complete

	statsReq := httptest.NewRequest("GET", "/stats?key="+stats.Key(), nil)

	resp, err := app.Test(statsReq)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.Nil(t, err)

	assert.Equal(t, len(response.Routes), 2)
	assert.Equal(t, len(response.Unused), 0)
	assert.Equal(t, response.Routes[0].Path, "/test")
	assert.Equal(t, int(response.Routes[0].Usage), 10)
}

func TestMemory(t *testing.T) {
	test(t, New(storage.NewMemory()))
}

func TestRedis(t *testing.T) {
	s, err := miniredis.Run()
	assert.Nil(t, err)

	test(t, New(storage.NewRedis(s.Addr(), "", 0)))
}
