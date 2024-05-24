// storage/memory.go
package storage

import (
	"sync"
)

type MemoryStorage struct {
	URLs  map[string]string
	Count map[string]int
	mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		URLs:  make(map[string]string),
		Count: make(map[string]int),
	}
}

func (m *MemoryStorage) SaveURL(original, short string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.URLs[original] = short
	m.Count[short]++
}

func (m *MemoryStorage) GetOriginalURL(short string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for original, s := range m.URLs {
		if s == short {
			return original, true
		}
	}
	return "", false
}

func (m *MemoryStorage) GetShortURL(original string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	short, exists := m.URLs[original]
	return short, exists
}

func (m *MemoryStorage) GetTopDomains() map[string]int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.Count
}
