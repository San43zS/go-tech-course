package flag

import (
	"flag"
	"fmt"
)

type Params struct {
	FileURL  string
	FileName string
}

func New() (IFlag, error) {
	p := &Params{}

	flag.StringVar(&p.FileURL, "url", "", "URL of the file to download")
	flag.StringVar(&p.FileName, "output", "", "FileName file name")
	flag.Parse()

	if p.FileURL == "" || p.FileName == "" {
		return nil, fmt.Errorf("missing required flags")
	}

	return p, nil
}

func (p Params) GetParams() Params {
	return p
}
