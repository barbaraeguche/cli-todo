package storage

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	File string
}

func New[T any](filename string) *Storage[T] {
	return &Storage[T]{File: filename}
}

func (s *Storage[T]) Load(data *T) error {
	// check for existence first
	f, err := os.OpenFile(s.File, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := os.ReadFile(s.File)
	if err != nil {
		return err
	}

	// nothing to unmarshal
	if len(b) == 0 {
		return nil
	}

	return json.Unmarshal(b, data)
}

func (s *Storage[T]) Store(data T) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.File, b, 0644)
}
