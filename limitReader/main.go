package main

import (
	"bytes"
	"fmt"

	"limitReader/limitRead"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func main() {
	data := "ыыпвы вып ы ывп ы ывп"
	buf := bytes.NewBufferString(data)

	limitReader := limitRead.New(buf, 9)

	_, err := limitReader.Read(make([]byte, 8))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	_, err = limitReader.Read(make([]byte, 1))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
