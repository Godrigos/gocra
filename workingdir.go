package gocra

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

/*
WorkDir GETs a job working directory files list from CIPRES server
   based on a job handle provided by jobStatus function and returns
   a WorkingDir structure
*/
func WorkDir(job JobStatus) WorkingDir {

	resp := login("GET", job.WorkingDirURI.URL)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var wdir WorkingDir
	xml.Unmarshal(b, &wdir)

	return wdir
}