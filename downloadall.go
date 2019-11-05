package gocra

import (
	"fmt"
	"io"
	"os"
	"sync"
)

/*
DownloadAll will use parallelization to get all the result files
from Cipres server.
It will create a new folder at the given destination directory,
named as the job metadata clientJobName or the jobhandle code if the
previuos field is empty.
*/
func DownloadAll(jr JobResults, job JobStatus, destDir string) {

	var wg sync.WaitGroup

	// Create folder to store downloaded files
	folder := job.JobHandle

	if job.Metadata.Entry.Value != "" {
		folder = job.Metadata.Entry.Value
	}
	os.Mkdir(destDir+folder, 0766)

	// Parallel downloading
	for i := range jr.Jobfiles.Jobfile {

		wg.Add(1)

		go download(destDir+folder+"/"+
			jr.Jobfiles.Jobfile[i].Filename,
			jr.Jobfiles.Jobfile[i].DownloadURI.URL, &wg)
	}
	wg.Wait()
	fmt.Println("Download completed!")
}

func download(filepath string, url string, wg *sync.WaitGroup) error {

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
	wg.Done()

	return err
}
