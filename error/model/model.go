package model

import (
	"errors"
	"time"
)

var (
	ErrInvalidURL           = errors.New("invalid or incorrect URL")
	ErrConnectionFailed     = errors.New("connection to the server failed")
	ErrDownloadFailed       = errors.New("file download failed")
	ErrFileNotFound         = errors.New("file not found(404)")
	ErrMissingRequiredFlags = errors.New("missing required flags")
)

const TimeOut = 30 * time.Second

const (
	URL    = "url"
	OutPut = "output"
)
