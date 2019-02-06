package storage

import (
	"io"
	"time"
)

const (
	// e.g. "2006-01-02T15:04:05Z"
	TimeFormat = time.RFC3339

	DeprecationWarning = "The `storage` parameter is deprecated. Please migrate to using built-in Terraform backends as described here: https://github.com/ljfranklin/terraform-resource/tree/WIP-tf-backends#backend-migration."
)

type Storage interface {
	Download(string, io.Writer) (Version, error)
	Upload(string, io.Reader) (Version, error)
	Delete(string) error
	Version(string) (Version, error)
	LatestVersion(string) (Version, error)
}

func BuildDriver(m Model) Storage {
	driverType := m.Driver
	if driverType == "" {
		driverType = S3Driver
	}

	var storageDriver Storage
	switch driverType {
	case S3Driver:
		storageDriver = NewS3(m)
	default:
		// calling model.Validate will throw error for this case
		return null{}
	}

	return storageDriver
}
