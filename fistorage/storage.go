package fistorage

type Storage interface {
	Increment(key string) error
	GetAll() (map[string]uint64, error)
	DeleteAll() error
	Close() error
}
