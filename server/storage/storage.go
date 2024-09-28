package storage

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
)

type StorageClient struct {
	logx.Logger
	Client         *minio.Client
	BaseBucketName string
	ParentDir      string
	ServiceName    string
}

func NewStorageClient(endpoint string, accessKey string, secretAccessKey string, parentDir string, serviceName string) *StorageClient {
	stClient := StorageClient{}
	// Initialize minio client object.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	stClient.Client = client
	stClient.Logger = logx.WithContext(ctx)
	stClient.ParentDir = parentDir
	stClient.BaseBucketName = serviceName
	return &stClient
}

func (sc *StorageClient) PutTypesContent(name string, contentType string, contentDisposition string, data *bytes.Buffer) (string, error) {

	ctx := context.Background()
	opt := minio.PutObjectOptions{}

	if len(contentType) > 0 {
		opt.ContentType = contentType
		opt.ContentDisposition = contentDisposition
	}
	_, err := sc.Client.PutObject(ctx, sc.BaseBucketName, name, data, int64(data.Len()), opt)
	if err != nil {
		return "", err
	}

	return sc.BaseBucketName + "/" + name, nil
}

// remove file with context
func (sc *StorageClient) RemoveFile(ctx context.Context, url string, domain string) error {
	opt := minio.RemoveObjectOptions{}
	objName := url[len(domain)+len(sc.BaseBucketName)+1:]

	err := sc.Client.RemoveObject(ctx, sc.BaseBucketName, objName, opt)
	if err != nil {
		return err
	}
	return nil
}

func (sc *StorageClient) RemoveFileWithObjName(ctx context.Context, objName string) error {
	opt := minio.RemoveObjectOptions{}

	err := sc.Client.RemoveObject(ctx, sc.BaseBucketName, objName, opt)
	if err != nil {
		return err
	}
	return nil
}
