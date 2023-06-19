package gcs

import (
	"bytes"
	"context"
	"fmt"
	"go-simple-template/config"
	"go-simple-template/pkg/storagex"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCS struct {
	client *storage.Client
	bucket string
}

func NewCGS(cfg *config.Config) (*GCS, error) {
	credentialPath := fmt.Sprintf("config/gcs/%s", cfg.GCS.CredentialsFile)

	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(credentialPath))
	if err != nil {
		return nil, err
	}

	return &GCS{
		client: client,
		bucket: cfg.GCS.GCSBucketName,
	}, nil
}

func (g *GCS) Upload(ctx context.Context, object *storagex.FileUploadObject) error {
	defer object.File.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	o := g.client.Bucket(g.bucket).Object(object.FileName)
	o = o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(ctx)
	if _, err := io.Copy(wc, object.File); err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w", err)
	}

	return nil
}

func (g *GCS) Download(ctx context.Context, fileName string) (*bytes.Buffer, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	rc, err := g.client.Bucket(g.bucket).Object(fileName).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %w", fileName, err)
	}
	defer rc.Close()

	buff := new(bytes.Buffer)
	buff.ReadFrom(rc)

	return buff, nil
}

func (g *GCS) Delete(ctx context.Context, fileName string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	if err := g.client.Bucket(g.bucket).Object(fileName).Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %w", fileName, err)
	}

	return nil
}
