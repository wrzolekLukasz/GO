package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sig)
}

// $ cat http.log.gz | gunzip | sha1sum
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close() // defer is used to execute the function after the surrounding function returns
	// defer is used to close the file after the function returns, usually used to close the file or file handling

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file) // := shadows the file variable from the outer scope
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
