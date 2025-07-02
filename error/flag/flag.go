package flag

import (
	"error/model"
	"flag"
)

type Params struct {
	FileURL  string
	FileName string
}

func New() (IFlag, error) {
	p := Params{}

	flag.StringVar(&p.FileURL, model.URL, "", "URL of the file to download")
	flag.StringVar(&p.FileName, model.OutPut, "", "FileName file name")
	flag.Parse()

	if p.FileURL == "" || p.FileName == "" {
		return nil, model.ErrMissingRequiredFlags
	}

	return p, nil
}

func (p Params) GetParams() Params {
	return p
}
