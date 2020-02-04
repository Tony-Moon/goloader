package main

import (
	"io"
	"net/http"
	"os"
)

var fileURL string

func main() {
	fileURL = "https://golangcode.com/images/avatar.jpg"
	if err := DownloadFile("avatar.jpg", fileURL); err != nil {
		panic(err)
	}
}

// DownloadFile will download a url to a locoal file.
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create a file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write to the file
	_, err = io.Copy(out, resp.Body)
	return err
}
