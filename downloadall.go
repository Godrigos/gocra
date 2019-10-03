package gocra

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

/*
DownloadAll will use download function in parallel to get all
the result files from Cipres server.
It will create a folder named as the job metadata clientJobName or
the jobhandle code if the previuos field is empty.
*/
func DownloadAll(jr JobResults, job JobStatus) {

	var wg sync.WaitGroup
	usr, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Create folder to store doenloaded files
	folder := job.JobHandle

	if job.Metadata.Entry.Value != "" {
		folder = job.Metadata.Entry.Value
	}
	os.Mkdir(usr+"/Downloads/"+folder, 0766)

	// Parallel downloading
	for i := range jr.Jobfiles.Jobfile {

		wg.Add(1)

		go download(usr+"/Downloads/"+folder+"/"+
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
