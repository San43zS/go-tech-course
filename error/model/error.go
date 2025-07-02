package model

import (
	"errors"
	"fmt"
)

func Resolver(err error) string {
	switch {
	case errors.Is(err, ErrInvalidURL):

		return "Error: Invalid URL. Please provide a valid URL."
	case errors.Is(err, ErrConnectionFailed):

		return "Error: Connection failed. Please check your internet connection and try again."
	case errors.Is(err, ErrDownloadFailed):

		return "Error: Download failed. Please check if the file is available for download."
	case errors.Is(err, ErrFileNotFound):

		return "Error: File not found. Please verify the URL and file availability on the server."
	default:

		return fmt.Sprintf("Error: An unexpected error occurred: %v\n", err)
	}
}
