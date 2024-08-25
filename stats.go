package stats

import (
	"github.com/krabiworld/stats/storage"
	"github.com/lutracorp/foxid-go"
)

type Response struct {
	Routes []Route `json:"routes"`
	Unused []Route `json:"unused"`
}

type Route struct {
	Path  string `json:"path"`
	Usage uint64 `json:"usage"`
}

type Stats struct {
	storage storage.Storage
	key     string
}

func New(storage storage.Storage) *Stats {
	id := foxid.Generate(foxid.Config{})

	return &Stats{
		storage: storage,
		key:     id.String(),
	}
}

func (s *Stats) Key() string {
	return s.key
}
