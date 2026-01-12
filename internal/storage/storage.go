package storage

import (
	"errors"
	"time"
)

var (
	ErrURLNotFound  = errors.New("url not found")
	ErrKeyCollision = errors.New("key collision occurred")
)

type URLRecord struct {
	Original string    `json:"original"`
	Created  time.Time `json:"created"`
	Visits   int       `json:"visits"`
}

type URLStore interface {
	Save(key string, originalURL string) error
	Get(key string) (URLRecord, error)
	IncrementVisits(key string) error
}
