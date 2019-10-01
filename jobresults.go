package gocra

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

/*
JobResult GETs a finished job files list from CIPRES server
based on a job handle provided by jobStatus function and returns
a JobResults structure
*/
func JobResult(jb JobStatus) JobResults {

	resp := login("GET", jb.ResultsURI.URL, nil)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var jr JobResults
	xml.Unmarshal(b, &jr)
	return jr
}
