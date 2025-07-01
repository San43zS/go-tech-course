package main

import (
	"error/dowload"
	"error/flag"
	"error/model"
	"fmt"
	"os"
)

func main() {
	f, err := flag.New()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	flagParams := f.GetParams()

	err = dowload.DownloadFile(flagParams.FileURL, flagParams.FileName)
	if err != nil {
		model.Resolver(err)
		os.Exit(1)
	}
}
