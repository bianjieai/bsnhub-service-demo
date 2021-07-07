package store

import (
	"encoding/binary"

	"github.com/cockroachdb/pebble"
)

// Store defines a struct for data store
type Store struct {
	db *pebble.DB
}

// NewStore constructs a new Store instance
func NewStore(path string) (*Store, error) {
	db, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		return &Store{}, err
	}

	return &Store{
		db: db,
	}, nil
}

// Set writes the given key-value into the store
func (s *Store) Set(key, value []byte) error {
	return s.db.Set(key, value, pebble.Sync)
}

// SetInt64 is a convenience to store the int64 typed value
func (s *Store) SetInt64(key []byte, value int64) error {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, uint64(value))

	return s.Set(key, bz)
}

// Get retrieves the value of the given key
func (s *Store) Get(key []byte) ([]byte, error) {
	value, closer, err := s.db.Get(key)
	if err != nil {
		return nil, err
	}

	defer closer.Close()

	return value, nil
}

// GetInt64 is a convenience to get the int64 typed value
func (s *Store) GetInt64(key []byte) (int64, error) {
	value, err := s.Get(key)
	if err != nil {
		return 0, err
	}

	return int64(binary.LittleEndian.Uint64(value)), nil
}

// Get retrieves the value of the given key
func (s *Store) Delete(key []byte) error {
	err := s.db.Delete(key, nil)
	if err != nil {
		return err
	}

	return nil
}
