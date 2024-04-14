package minio

import (
	"bytes"
	"context"
	"fmt"
	"go-simple-template/config"
	"go-simple-template/internal/pkg/utils"
	"go-simple-template/pkg/storagex"

	"github.com/minio/minio-go"
)

type Minio struct {
	client *minio.Client
	bucket string
}

func NewMinio() (*Minio, error) {
	client, err := minio.New(
		config.MinioEndpoint(),
		config.MinioAccessKeyID(),
		config.MinioAccessKeySecret(),
		false,
	)
	if err != nil {
		return nil, err
	}

	client.SetAppInfo(config.AppName(), config.AppVersion())

	m := &Minio{
		client: client,
		bucket: config.MinioBucketName(),
	}

	return m, nil
}

type MinioInterface interface {
	Upload(file *storagex.FileUploadObject) error
}

func (r *Minio) Upload(ctx context.Context, file *storagex.FileUploadObject) error {
	fileType, _ := utils.GetContentTypeFromFile(file.File)

	_, err := r.client.PutObjectWithContext(ctx, r.bucket, file.FileName, file.File, -1, minio.PutObjectOptions{
		ContentType: fileType,
	})
	if err != nil {
		return fmt.Errorf("PutObjectWithContext: %w", err)
	}

	return nil
}

func (r *Minio) Download(ctx context.Context, fileName string) (*bytes.Buffer, error) {
	object, err := r.client.GetObjectWithContext(ctx, r.bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("GetObjectWithContext: %w", err)
	}
	defer object.Close()

	buff := new(bytes.Buffer)
	buff.ReadFrom(object)

	return buff, nil
}

func (r *Minio) Delete(ctx context.Context, fileName string) error {
	err := r.client.RemoveObject(r.bucket, fileName)
	if err != nil {
		return fmt.Errorf("RemoveObjectWithContext: %w", err)
	}

	return nil
}
