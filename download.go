package gocra

import (
	"io"
	"os"
)

/*
DownloadFile will download a url to a local file. It's efficient
because it will write as it downloads and not load the whole file
into memory.
The target folder should be given as the first funtion argument, followed
by the file url.
Returns an error if something goes wrong.
*/
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp := login("GET", url, nil)
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return err
}
