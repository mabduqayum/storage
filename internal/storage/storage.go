package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mabduqayum/storage/v2/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(fileName, blob)
	if err != nil {
		return nil, err
	}
	s.files[newFile.Id] = newFile
	return newFile, nil
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	if f, ok := s.files[fileId]; ok {
		return f, nil
	}
	return nil, fmt.Errorf("file %v not found", fileId)
}
