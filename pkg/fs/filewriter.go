package fs

import (
	"fmt"
	"log"
	"os"
)

// Writer is used for writing a file as defined format
type Writer struct {
	FileName string
	Perm     os.FileMode
}

func (w *Writer) Write(data string, size int) (int, error) {

	log.Printf("opening %s file\n", w.FileName)

	f, err := os.OpenFile(w.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, w.Perm)

	if err != nil {
		return 0, err
	}

	defer f.Close()

	n, err := f.WriteString(data)

	if n != size {
		return 0, fmt.Errorf("incomplete writing to file > data size : %d , written bytes: %v", size, n)
	}

	if err != nil {
		return 0, err
	}

	log.Printf("write data %v bytes\n", n)

	return n, nil

}
