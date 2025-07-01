package model

import (
	"errors"
	"fmt"
)

func Resolver(err error) {
	switch {
	case errors.Is(err, ErrInvalidURL):
		fmt.Println("Error: Invalid URL. Please provide a valid URL.")
	case errors.Is(err, ErrConnectionFailed):
		fmt.Println("Error: Connection failed. Please check your internet connection and try again.")
	case errors.Is(err, ErrDownloadFailed):
		fmt.Println("Error: Download failed. Please check if the file is available for download.")
	case errors.Is(err, ErrFileNotFound):
		fmt.Println("Error: File not found. Please verify the URL and file availability on the server.")
	default:
		fmt.Printf("Error: An unexpected error occurred: %v\n", err)
	}
}
