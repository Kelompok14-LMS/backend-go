package helper

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/Kelompok14-LMS/backend-go/constants"
)

type StorageConfig struct {
	StorageClient *storage.Client
	BucketName    string
}

func NewCloudStorage(storageClient *storage.Client, bucketName string) *StorageConfig {
	return &StorageConfig{
		StorageClient: storageClient,
		BucketName:    bucketName,
	}
}

// UploadImage helper to upload image into cloud storage
func (s *StorageConfig) UploadImage(ctx context.Context, objName string, file multipart.File) (string, error) {
	// bucket name to store the images
	imgBucket := s.StorageClient.Bucket(s.BucketName)

	imageDir := fmt.Sprintf("%s/%s", constants.IMAGES_DIR, objName)

	// object to be stored in cloud storage
	object := imgBucket.Object(imageDir)
	wc := object.NewWriter(ctx)

	// skip the object cache, always retrieve fresh object
	wc.ObjectAttrs.CacheControl = "Chace-Control:no-cache, max-age=0"

	// upload the object with storage.Writer
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	// should be closed when process is done
	if err := wc.Close(); err != nil {
		return "", err
	}

	imgUrl := fmt.Sprintf("%s/%s/%s", constants.STORAGE_URL, s.BucketName, imageDir)

	return imgUrl, nil
}

// UploadVideo helper to upload video into cloud storage
func (s *StorageConfig) UploadVideo(ctx context.Context, objName string) (string, error) {
	panic("implement me")
}

// UploadAsset helper to upload asset (i.e pdf, etc) into cloud storage
func (s *StorageConfig) UploadAsset(ctx context.Context, objName string) (string, error) {
	panic("implement me")
}

// DeleteObject helper to delete object from cloud storage
func (s *StorageConfig) DeleteObject(ctx context.Context, objName string) error {
	bucket := s.StorageClient.Bucket(s.BucketName)

	// remove the base url and the bucket name
	path := fmt.Sprintf("%s/%s/", constants.STORAGE_URL, s.BucketName)
	objDir := strings.Replace(objName, path, "", 1)

	object := bucket.Object(objDir)

	if err := object.Delete(ctx); err != nil {
		return fmt.Errorf("delete err: %v", err)
	}

	return nil
}
