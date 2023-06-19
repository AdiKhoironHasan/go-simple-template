package storagex

import (
	"bytes"
	"context"
)

type Storage struct {
	Client StorageInterface
}

func NewStorage(s StorageInterface) *Storage {
	return &Storage{Client: s}
}

type StorageInterface interface {
	Upload(ctx context.Context, file *FileUploadObject) error
	Download(ctx context.Context, fileName string) (*bytes.Buffer, error)
	Delete(ctx context.Context, fileName string) error
}
