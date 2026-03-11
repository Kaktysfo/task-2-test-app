package storage

import (
	"errors"
	"os"
)

type Storage struct {
	filename string
}

var (
	notExistErr = errors.New("такого файла не существует")
	saveErr     = errors.New("ошибка записи файла")
)

func NewStorage(filename string) *Storage {
	return &Storage{
		filename: filename,
	}
}

func (s *Storage) Save(data []byte) error {
	err := os.WriteFile(s.filename, data, 0644)
	if err != nil {
		return saveErr
	}
	return nil
}

func (s *Storage) Load() ([]byte, error) {
	if _, err := os.ReadFile(s.filename); os.IsNotExist(err) {
		return nil, notExistErr
	}
	data, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, saveErr
	}
	return data, nil
}
