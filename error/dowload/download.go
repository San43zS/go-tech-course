package dowload

import (
	"error/client"
	"error/model"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(fileURL, fileName string) error {
	clnt := client.New()

	resp, err := clnt.Client.Get(fileURL)

	if err != nil {
		switch {
		case os.IsTimeout(err):
			return model.ErrConnectionFailed
		default:
			return model.ErrFileNotFound
		}
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return model.ErrFileNotFound
	default:
		return model.ErrDownloadFailed
	}

	outFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}
