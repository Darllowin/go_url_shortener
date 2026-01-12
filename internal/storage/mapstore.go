package storage

import (
	"sync"
	"time"
)

type MapStore struct {
	mtx  sync.RWMutex
	data map[string]URLRecord
}

func NewMapStore() *MapStore {
	return &MapStore{
		data: make(map[string]URLRecord),
	}
}

func (s *MapStore) Save(key string, originalURL string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	// сhecking for collisions
	if _, ok := s.data[key]; ok {
		return ErrKeyCollision
	}

	s.data[key] = URLRecord{
		Original: originalURL,
		Created:  time.Now(),
		Visits:   0,
	}

	return nil
}

func (s *MapStore) Get(key string) (URLRecord, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	record, ok := s.data[key]
	if !ok {
		return URLRecord{}, ErrURLNotFound
	}

	return record, nil
}

func (s *MapStore) IncrementVisits(key string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	record, ok := s.data[key]
	if !ok {
		return ErrURLNotFound
	}

	record.Visits++
	s.data[key] = record

	return nil
}
