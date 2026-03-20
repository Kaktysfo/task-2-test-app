package storage

import (
	"errors"
	"os"
)

type JsonStorage struct {
	*Storage
}

var (
	notExistErr = errors.New("такого файла не существует")
	saveErr     = errors.New("ошибка записи файла")
)

func NewJsonStorage(filename string) *JsonStorage {
	return &JsonStorage{
		&Storage{filename: filename},
	}
}

func (s *JsonStorage) Save(data []byte) error {
	err := os.WriteFile(s.GetFilename(), data, 0644)
	if err != nil {
		return saveErr
	}
	return nil
}

func (s *JsonStorage) Load() ([]byte, error) {
	if _, err := os.ReadFile(s.GetFilename()); os.IsNotExist(err) {
		return nil, notExistErr
	}
	data, err := os.ReadFile(s.GetFilename())
	if err != nil {
		return nil, saveErr
	}
	return data, nil
}
