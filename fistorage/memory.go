package fistorage

import (
	"sync"
)

type Memory struct {
	sync.RWMutex
	m map[string]uint64
}

func NewMemory() *Memory {
	return &Memory{m: make(map[string]uint64)}
}

func (m *Memory) Increment(key string) error {
	m.Lock()
	m.m[key]++
	m.Unlock()

	return nil
}

func (m *Memory) GetAll() (map[string]uint64, error) {
	m.RLock()

	cp := make(map[string]uint64, len(m.m))
	for k, v := range m.m {
		cp[k] = v
	}

	m.RUnlock()

	return cp, nil
}

func (m *Memory) Clear() error {
	m.Lock()
	m.m = make(map[string]uint64)
	m.Unlock()

	return nil
}

func (m *Memory) Close() error {
	m.Lock()
	m.m = nil
	m.Unlock()

	return nil
}
