package gocra

import (
	"io"
	"os"
	"sync"
)

/*
DownloadFile will download a url to a local file. It's efficient
because it will write as it downloads and not load the whole file
into memory.
*/
func DownloadFile(filepath string, url string, wg *sync.WaitGroup) error {

	// Get the data
	resp := login("GET", url)
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	wg.Done()

	return err
}
