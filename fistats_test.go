package fistats

import (
	"encoding/json"
	"github.com/alicebob/miniredis/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/krabiworld/fistats/fistorage"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func test(t *testing.T, cfg Config) {
	app := fiber.New()

	var key string

	app.Use(New(&key, cfg))

	log.Infof("Key: %s", key)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	testReq := httptest.NewRequest("GET", "/test", nil)
	for i := 0; i < 10; i++ {
		_, err := app.Test(testReq)
		assert.Nil(t, err)
	}

	time.Sleep(100 * time.Millisecond) // wait for all goroutines to complete

	statsReq := httptest.NewRequest("GET", key, nil)

	resp, err := app.Test(statsReq)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response []Route
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.Nil(t, err)

	assert.Equal(t, len(response), 1)
	assert.Equal(t, response[0].Path, "/test")
	assert.Equal(t, int(response[0].Usage), 10)
}

func TestMemory(t *testing.T) {
	test(t, Config{Storage: fistorage.NewMemory()})
}

func TestRedis(t *testing.T) {
	s, err := miniredis.Run()
	assert.Nil(t, err)

	test(t, Config{Storage: fistorage.NewRedis(s.Addr(), "", 0)})
}
