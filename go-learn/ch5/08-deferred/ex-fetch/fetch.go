package main

import (
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if e := f.Close(); e != nil {
			err = e
		}
	}()

	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	if err != nil {
		return "", 0, err
	}
	return local, n, err
}

func main() {
	for _, url := range os.Args[1:] {
		fetch(url)
	}
}
